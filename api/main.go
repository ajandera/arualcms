package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"io"
	"log"
	model "main/model"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"strings"
	"time"
)

type ClientData struct {
	db *gorm.DB
}

type tokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}

type Account struct {
	Username string
	Password string
}

type Reset struct {
	Token    string
	Password string
}

type Email struct {
	Email      string
	Subject    string
	HtmlString string
}

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
	mySigningKey   = []byte("DFGDFGhcsadkjhfwe+ě+23123asldxjhsdljfh1234234")
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewConnect(dsn string) ClientData {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	errExt := db.Raw("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if errExt != nil {
		panic(errExt)
	}
	// Migrate the schema
	db.AutoMigrate(
		&model.Post{},
		&model.User{},
		&model.Text{},
		&model.Language{},
		&model.Setting{},
		&model.File{},
		&model.Site{})
	return ClientData{db}
}

func IsValidUUID(uuid uuid.UUID) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid.String())
}

type userPermission struct {
	ParentId uuid.UUID
	Id       uuid.UUID
	Name     string
}

func isAuthorized(w http.ResponseWriter, r *http.Request, site uuid.UUID, c ClientData) (bool, string) {
	if r.Header["Authorization"] != nil {
		var sendToken = strings.Replace(r.Header["Authorization"][0], "Bearer ", "", 1)
		token, err := jwt.Parse(sendToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("JWT token not pass")
			}
			return mySigningKey, nil
		})

		claims := token.Claims.(jwt.MapClaims)
		// check if store belongs to account
		if IsValidUUID(site) {
			var user userPermission
			idFromClaim, _ := uuid.Parse(fmt.Sprintf("%v", claims["id"]))
			c.db.Model(&model.User{Id: idFromClaim}).First(&user)
			if IsValidUUID(user.ParentId) {
				if isPermitted(user.ParentId, site, c) == false {
					http.Error(w, "Not Permitted", http.StatusForbidden)
					return false, ""
				}
			} else {
				if isPermitted(user.Id, site, c) == false {
					http.Error(w, "Not Permitted", http.StatusForbidden)
					return false, ""
				}
			}
		}

		if err != nil && token != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return false, ""
		}
		return true, fmt.Sprintf("%v", claims["id"])
	} else {
		http.Error(w, "Not Authorized", http.StatusForbidden)
		return false, ""
	}
}

func isPermitted(userId uuid.UUID, siteId uuid.UUID, c ClientData) bool {
	var site model.Site
	// special check for exit uuid
	c.db.First(&model.Site{}, "id = ?", userId, siteId).Scan(&site)
	if IsValidUUID(site.Id) == false {
		return true
	} else {
		c.db.First(&model.Site{}, "user_id = ? AND id = ?", userId, siteId).Scan(&site)
		return IsValidUUID(site.Id)
	}
}

func setupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GenerateJWT(name string, id string) (map[string]string, error) {

	// access token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = name
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["id"] = id
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refreshTokenString, err := refreshToken.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return map[string]string{
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
	}, nil
}

func refresh(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	var tokenReq = tokenReqBody{}
	err := json.NewDecoder(r.Body).Decode(&tokenReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf(err.Error())
		return
	}

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		var user model.User
		c.db.First(&model.User{}, "id = ?", fmt.Sprint(claims["id"])).Scan(&user)
		if IsValidUUID(user.Id) {

			newTokenPair, err := GenerateJWT(user.Name, user.Id.String())
			if err != nil {
				return
			}

			response := simplejson.New()
			if IsValidUUID(user.Id) == true && err == nil {
				response.Set("success", true)
				response.Set("jwt", newTokenPair)
			} else {
				response.Set("success", false)
				response.Set("error", "Refresh token is not valid.")
			}

			payload, err := response.MarshalJSON()
			if err != nil {
				log.Fatalf(err.Error())
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(payload)
		}
	}
}

func auth(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	// Declare a new Visitor struct.
	var a Account

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf(err.Error())
		return
	}

	response := simplejson.New()

	var account model.User
	c.db.First(&model.User{}, "username = ?", a.Username).Scan(&account)
	pw := CheckPasswordHash(a.Password, account.Password)
	if pw == true {
		token, err := GenerateJWT(account.Name, account.Id.String())
		if IsValidUUID(account.Id) == true && err == nil {
			response.Set("success", true)
			response.Set("jwt", token)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusForbidden)
			response.Set("success", false)
			response.Set("error", "Email and password not match.")
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		response.Set("success", false)
		response.Set("error", "Email and password not match.")
	}

	payload, err := response.MarshalJSON()
	if err != nil {
		log.Fatalf(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func getPosts(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		var posts []model.Post
		var files []model.File
		c.db.Model(&model.Post{}).Where("site_id = ?", siteId).Scan(&posts)
		for _, v := range posts {
			var f model.File
			c.db.Model(&model.File{}).Where("id = ?", v.File).Scan(&f)
			files = append(files, f)
		}
		response.Set("success", true)
		response.Set("posts", posts)
		response.Set("files", files)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func createPost(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Post struct.
		var post model.Post

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		var site model.Site
		c.db.Model(&model.Site{}).Where("id = ?", siteId).Scan(&site)

		response := simplejson.New()

		c.db.Create(&model.Post{
			Body:        post.Body,
			Title:       post.Title,
			Excerpt:     post.Excerpt,
			Published:   post.Published,
			MetaTitle:   post.MetaTitle,
			Keywords:    post.Keywords,
			Description: post.Description,
			SiteId:      siteId.String(),
			Site:        site,
		})
		response.Set("success", true)
		response.Set("message", "Post created successfully.")
		response.Set("post", post.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func updatePost(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {

		// Declare a new Post struct.
		var post model.Post

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()
		c.db.Model(model.Post{}).Where("id = ?", post.Id).Updates(model.Post{
			Body:        post.Body,
			Title:       post.Title,
			Excerpt:     post.Excerpt,
			Published:   post.Published,
			MetaTitle:   post.MetaTitle,
			Keywords:    post.Keywords,
			Description: post.Description,
			File:        post.File})
		response.Set("success", true)
		response.Set("message", "Post updated successfully.")
		response.Set("post", post.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getPostDetail(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		postId := vars["postId"]

		var post model.Post
		c.db.First(&model.Post{}, "id = ?", postId).Scan(&post)
		response.Set("success", true)
		response.Set("post", post)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func deletePost(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		var p model.Post
		id := vars["postId"]
		c.db.Model(&model.Post{}).Where("id = ?", id).Delete(&p)
		response := simplejson.New()
		response.Set("success", true)
		response.Set("message", "Post deleted successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getSetting(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		var settings []model.Setting
		c.db.Model(&model.Setting{}).Where("site_id = ?", siteId).Scan(&settings)
		response.Set("success", true)
		response.Set("settings", settings)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func createSetting(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Setting struct.
		var setting model.Setting

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&setting)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		var site model.Site
		c.db.Model(&model.Site{}).Where("id = ?", siteId).Scan(&site)

		response := simplejson.New()

		c.db.Create(&model.Setting{
			Key:    setting.Key,
			Value:  setting.Value,
			SiteId: siteId.String(),
			Site:   site,
		})
		response.Set("success", true)
		response.Set("message", "Setting created successfully.")
		response.Set("post", setting.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func updateSetting(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Setting struct.
		var setting model.Setting

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&setting)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()
		c.db.Model(model.Setting{}).Where("id = ?", setting.Id).Updates(model.Setting{
			Key:   setting.Key,
			Value: setting.Value})
		response.Set("success", true)
		response.Set("message", "Setting updated successfully.")
		response.Set("setting", setting.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getTexts(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		var texts []model.Text
		c.db.Model(&model.Text{}).Where("site_id = ?", siteId).Scan(&texts)
		response.Set("success", true)
		response.Set("texts", texts)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func createText(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Text struct.
		var text model.Text

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		var site model.Site
		c.db.Model(&model.Site{}).Where("id = ?", siteId).Scan(&site)

		response := simplejson.New()

		c.db.Create(&model.Text{
			Key:    text.Key,
			Value:  text.Value,
			SiteId: siteId.String(),
			Site:   site,
		})
		response.Set("success", true)
		response.Set("message", "Text created successfully.")
		response.Set("post", text.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func updateText(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Text struct.
		var text model.Text

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()

		c.db.Model(model.Text{}).Where("id = ?", text.Id).Updates(model.Text{
			Key:   text.Key,
			Value: text.Value})
		response.Set("success", true)
		response.Set("message", "Text updated successfully.")
		response.Set("post", text.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getText(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		vars := mux.Vars(r)
		key := vars["key"]

		var text model.Text
		c.db.First(&model.Text{}, "key = ? AND site_id = ?", key, siteId).Scan(&text)
		response.Set("success", true)
		response.Set("text", text)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func deleteText(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		var t model.Text
		id := vars["textId"]
		c.db.Model(&model.Text{}).Where("id = ?", id).Delete(&t)
		response := simplejson.New()
		response.Set("success", true)
		response.Set("message", "text deleted successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		var users []model.User
		c.db.Model(&model.User{}).Scan(&users)
		response.Set("success", true)
		response.Set("users", users)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func createRegistration(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	// Declare a new User struct.
	var user model.User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf(err.Error())
		return
	}

	response := simplejson.New()

	pw, hashErr := HashPassword(user.Password)
	if hashErr != nil {
		log.Fatalf(hashErr.Error())
	}
	c.db.Create(&model.User{
		Name:     user.Name,
		Username: user.Username,
		Password: pw,
		ParentId: "",
	})
	response.Set("success", true)
	response.Set("message", "Your registration was successful.")
	response.Set("post", user.Id)
	w.WriteHeader(http.StatusOK)

	payload, err := response.MarshalJSON()
	if err != nil {
		log.Fatalf(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func createUser(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new User struct.
		var user model.User

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()

		pw, hashErr := HashPassword(user.Password)
		if hashErr != nil {
			log.Fatalf(hashErr.Error())
		}

		var u = model.User{}

		c.db.Create(&model.User{
			Name:     user.Name,
			Username: user.Username,
			Password: pw,
			ParentId: user.ParentId,
		}).Scan(&u)

		// generate permission
		var sites []model.Site
		c.db.Model(&model.Site{}).Where("user_id = ?", user.ParentId).Scan(&sites)
		for _, v := range sites {
			c.db.Create(&model.Permission{
				UserId: u.Id.String(),
				SiteId: v.Id.String(),
				Site:   v,
				User:   u,
				Role:   "admin",
			})
		}

		response.Set("success", true)
		response.Set("message", "User created successfully.")
		response.Set("user", u.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func updateUser(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new User struct.
		var user model.User

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()

		c.db.Model(model.User{}).Where("id = ?", user.Id).Updates(model.User{
			Name:     user.Name,
			Username: user.Username})
		response.Set("success", true)
		response.Set("message", "User updated successfully.")
		response.Set("post", user.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getUser(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		vars := mux.Vars(r)
		userId := vars["userId"]

		response := simplejson.New()

		var user model.User
		c.db.First(&model.User{}, "id = ?", userId).Scan(&user)
		response.Set("success", true)
		response.Set("user", user)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		var u model.User
		vars := mux.Vars(r)
		id := vars["userId"]
		c.db.Model(&model.User{}).Where("id = ?", id).Delete(&u)
		response := simplejson.New()
		response.Set("success", true)
		response.Set("message", "User deleted successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getUserByEmail(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		vars := mux.Vars(r)
		username := vars["username"]

		response := simplejson.New()

		var user model.User
		c.db.First(&model.User{}, "username = ?", username).Scan(&user)
		response.Set("success", true)
		response.Set("user", user)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getFiles(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {

		response := simplejson.New()

		var files []model.File
		c.db.Model(&model.File{}).Where("site_id = ?", siteId).Scan(&files)
		response.Set("success", true)
		response.Set("files", files)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func updateFile(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	fileId := vars["fileId"]
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new User struct.
		var file model.File

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()

		c.db.Model(model.File{}).Where("id = ?", fileId).Updates(model.File{
			Gallery: file.Gallery})
		response.Set("success", true)
		response.Set("message", "File updated successfully.")
		response.Set("post", file.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		e := r.ParseMultipartForm(50 << 20)
		if e != nil {
			log.Printf("F: " + e.Error())
		}
		file, handle, err := r.FormFile("file")
		if err != nil {
			log.Printf("From file: " + err.Error())
		}

		defer file.Close()

		// Create file
		path := "files/" + siteId.String()
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}

		dst, err := os.Create(path + "/" + handle.Filename)
		defer dst.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var site model.Site
		c.db.Model(&model.Site{}).Where("id = ?", siteId).Scan(&site)

		var f model.File

		c.db.Create(&model.File{
			Name:    handle.Filename,
			Src:     "/" + siteId.String() + "/" + handle.Filename,
			Gallery: "",
			SiteId:  siteId.String(),
			Site:    site,
		}).Scan(&f)

		// prepare data to response
		response := simplejson.New()
		response.Set("success", true)
		response.Set("message", "File uploaded successfully.")
		response.Set("file", f.Id)
		response.Set("src", "/"+siteId.String()+"/"+handle.Filename)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	}
}

func getFile(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		vars := mux.Vars(r)
		fileId := vars["fileId"]

		response := simplejson.New()

		var file model.File
		c.db.First(&model.File{}, "id = ?", fileId).Scan(&file)
		response.Set("success", true)
		response.Set("file", file)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func deleteFile(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		var f model.File
		vars := mux.Vars(r)
		id := vars["fileId"]
		c.db.Model(&model.File{}).Where("id = ?", id).Delete(&f)
		response := simplejson.New()
		response.Set("success", true)
		response.Set("message", "File deleted successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getLanguages(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		var languages []model.Language
		c.db.Model(&model.Language{}).Where("site_id = ?", siteId).Scan(&languages)
		response.Set("success", true)
		response.Set("languages", languages)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func createLanguage(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Language struct.
		var language model.Language

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&language)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		var site model.Site
		c.db.Model(&model.Site{}).Where("id = ?", siteId).Scan(&site)

		response := simplejson.New()

		c.db.Create(&model.Language{
			Name:    language.Name,
			Key:     language.Key,
			Default: language.Default,
			SiteId:  siteId.String(),
			Site:    site,
		})

		response.Set("success", true)
		response.Set("message", "Language created successfully.")
		response.Set("language", language.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func updateLanguage(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Language struct.
		var language model.Language

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&language)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		var site model.Site
		c.db.Model(&model.Site{}).Where("id = ?", siteId).Scan(&site)

		response := simplejson.New()

		c.db.Model(model.Language{}).Where("id = ?", language.Id).Updates(model.Language{
			Name:    language.Name,
			Key:     language.Key,
			Default: language.Default,
			SiteId:  siteId.String(),
			Site:    site})

		response.Set("success", true)
		response.Set("message", "Language updated successfully.")
		response.Set("language", language.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getLanguageByCode(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])

	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		code := vars["code"]

		response := simplejson.New()

		var language model.Language
		c.db.First(&model.Language{}, "code = ?", code).Scan(&language)
		response.Set("success", true)
		response.Set("language", language)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func deleteLanguage(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		var l model.Language
		id := vars["languageId"]
		c.db.Model(&model.Language{}).Where("id = ?", id).Delete(&l)
		response := simplejson.New()
		response.Set("success", true)
		response.Set("message", "Language deleted successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func getSites(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	if auth, userId := isAuthorized(w, r, uuid.New(), c); auth == true {
		response := simplejson.New()

		var sites []model.Site
		c.db.Model(&model.Site{}).Where("user_id = ?", userId).Scan(&sites)
		response.Set("success", true)
		response.Set("sites", sites)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func createSite(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	if auth, userId := isAuthorized(w, r, uuid.New(), c); auth == true {
		// Declare a new Language struct.
		var site model.Site
		var res model.Site
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&site)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}
		response := simplejson.New()
		res.UserId = userId
		res.Name = site.Name

		var s = model.Site{}
		c.db.Create(&res).Scan(&s)

		// generate permission
		var users []model.User
		c.db.Model(&model.User{}).Where("parent_id = ?", userId).Scan(&users)
		for _, u := range users {
			c.db.Create(&model.Permission{
				UserId: u.Id.String(),
				SiteId: s.Id.String(),
				User:   u,
				Site:   s,
				Role:   "admin",
			})
		}

		createDefaultSetting(res.Id.String(), c)
		response.Set("success", true)
		response.Set("message", "Site created successfully.")
		response.Set("site", res.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func createDefaultSetting(siteId string, c ClientData) {
	settings := [...]string{"subtitle", "author", "contactEmail", "contactPhone", "metaTitle", "metaDescription", "metaKeywords", "smtp", "smtp_user", "smtp_password", "smtp_port", "title"}
	for _, v := range settings {
		c.db.Create(&model.Setting{
			Key:    v,
			Value:  "",
			SiteId: siteId,
		})
	}
}

func updateSite(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Language struct.
		var site model.Site

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&site)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()

		c.db.Model(model.Site{}).Where("id = ?", site.Id).Updates(model.Site{
			Name: site.Name})

		response.Set("success", true)
		response.Set("message", "Setting updated successfully.")
		response.Set("site", site.Id)
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func deleteSite(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		var s model.Site
		c.db.Model(&model.Language{}).Where("id = ?", siteId).Delete(&s)
		response := simplejson.New()
		response.Set("success", true)
		response.Set("message", "Setting deleted successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func sendEmail(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := isAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Email struct.
		var mail Email

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&mail)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()

		sendEmailWithoutTemplate(mail.Email, mail.Subject, mail.HtmlString, c, r)

		response.Set("success", true)
		response.Set("message", "Email send successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func forgot(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 8
	token := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)

	vars := mux.Vars(r)
	username := vars["username"]

	var user model.User
	c.db.First(&model.User{}, "username = ?", username).Scan(&user)

	t := time.Now().Add(time.Hour * 24)
	c.db.Model(model.User{}).Where("username = ?", username).Updates(model.User{
		ForgotToken: token,
		ValidToken:  t})

	sendEmailWithTemplate(username, "Recover your password", "forgot", token, c, r)

	response := simplejson.New()
	response.Set("success", true)

	payload, err := response.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)

}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func recovery(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	// Declare a new Reset struct.
	var reset Reset

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&reset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf(err.Error())
		return
	}

	var a model.User
	c.db.Model(&model.User{}).Where("forgot_token = ?", reset.Token).First(&a)

	t1 := time.Now()
	t2 := a.ValidToken
	hourDiff := t2.Sub(t1).Hours()

	if hourDiff < 24 {
		a.ForgotToken = ""
		if len(reset.Password) > 6 {
			hash, _ := HashPassword(reset.Password)
			a.Password = hash
		}
	}
	c.db.Save(&a)

	response := simplejson.New()
	response.Set("success", true)
	response.Set("message", "You password updated successfully.")

	payload, err := response.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func sendEmailWithTemplate(email string, subject string, templateName string, token string, c ClientData, r *http.Request) {
	vars := mux.Vars(r)
	siteId := vars["siteId"]
	var settings []model.Setting
	c.db.Model(&model.Setting{}).Where("site_id = ?", siteId).Scan(&settings)
	var smtpHost string
	var smtpPort string
	var from string
	var password string

	for _, v := range settings {
		if v.Key == "smtpHost" {
			smtpHost = v.Value
		}

		if v.Key == "smtpPort" {
			smtpPort = v.Value
		}

		if v.Key == "smtpUser" {
			from = v.Value
		}

		if v.Key == "smtpPassword" {
			password = v.Value
		}
	}

	// prepare template
	t, _ := template.ParseFiles("templates/" + templateName + ".html")
	var body bytes.Buffer

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = email
	headers["Subject"] = subject

	for k, v := range headers {
		body.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v)))
	}

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: "+subject+" \n%s\n\n", mimeHeaders)))

	err := t.Execute(&body, struct {
		Email string
		Link  string
	}{
		Email: email,
		Link:  "/restore?token=" + token,
	})

	if err != nil {
		log.Printf(err.Error())
		return
	} else {
		// TLS config
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         smtpHost,
		}

		// Here is the key, you need to call tls.Dial instead of smtp.Dial
		// for smtp servers running on 465 that require an ssl connection
		// from the very beginning (no starttls)
		servername := smtpHost + ":" + smtpPort
		conn, err := tls.Dial("tcp", servername, tlsconfig)
		if err != nil {
			log.Println(err)
		}

		c, err := smtp.NewClient(conn, smtpHost)
		if err != nil {
			log.Println(err)
		}

		// Auth
		auth := smtp.PlainAuth("", from, password, smtpHost)
		if err = c.Auth(auth); err != nil {
			log.Println(err)
		}

		// To && From
		if err = c.Mail(from); err != nil {
			log.Println(err)
		}

		if err = c.Rcpt(email); err != nil {
			log.Println(err)
		}

		// Data
		wr, err := c.Data()
		if err != nil {
			log.Println(err)
		}

		_, err = wr.Write(body.Bytes())
		if err != nil {
			log.Println(err)
		}

		err = wr.Close()
		if err != nil {
			log.Println(err)
		}
		errC := c.Quit()
		if errC != nil {
			log.Println(err.Error())
		}
	}
}

func sendEmailWithoutTemplate(email string, subject string, htmlString string, c ClientData, r *http.Request) {

	vars := mux.Vars(r)
	siteId := vars["siteId"]
	var settings []model.Setting
	c.db.Model(&model.Setting{}).Where("site_id = ?", siteId).Scan(&settings)
	var smtpHost string
	var smtpPort string
	var from string
	var password string

	for _, v := range settings {
		if v.Key == "smtpHost" {
			smtpHost = v.Value
		}

		if v.Key == "smtpPort" {
			smtpPort = v.Value
		}

		if v.Key == "smtpUser" {
			from = v.Value
		}

		if v.Key == "smtpPassword" {
			password = v.Value
		}
	}

	var body bytes.Buffer

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = email
	headers["Subject"] = subject

	for k, v := range headers {
		body.Write([]byte(fmt.Sprintf("%s: %s\r\n", k, v)))
	}

	sub := "Subject: " + subject + "\n"
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	content := "<html><body>" + htmlString + "</body></html>"
	body.Write([]byte(sub + mimeHeaders + content))

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpHost,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	servername := smtpHost + ":" + smtpPort
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Println(err)
	}

	cl, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		log.Println(err)
	}

	// Auth
	auth := smtp.PlainAuth("", from, password, smtpHost)
	if err = cl.Auth(auth); err != nil {
		log.Println(err)
	}

	// To && From
	if err = cl.Mail(from); err != nil {
		log.Println(err)
	}

	if err = cl.Rcpt(email); err != nil {
		log.Println(err)
	}

	// Data
	wr, err := cl.Data()
	if err != nil {
		log.Println(err)
	}

	_, err = wr.Write(body.Bytes())
	if err != nil {
		log.Println(err)
	}

	err = wr.Close()
	if err != nil {
		log.Println(err)
	}
	errC := cl.Quit()
	if errC != nil {
		log.Println(err.Error())
	}
}

func me(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	if auth, _ := isAuthorized(w, r, uuid.New(), c); auth == true {
		var sendToken = strings.Replace(r.Header["Authorization"][0], "Bearer ", "", 1)
		token, err := jwt.Parse(sendToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("JWT token not pass")
			}
			return mySigningKey, nil
		})
		claims := token.Claims.(jwt.MapClaims)
		response := simplejson.New()
		var user model.User
		var permission []model.Permission
		c.db.First(&model.User{}, "id = ?", fmt.Sprintf("%v", claims["id"])).Scan(&user)
		c.db.Model(&model.Permission{}).Where("user_id = ?", fmt.Sprintf("%v", claims["id"])).Scan(&permission)
		response.Set("success", true)
		response.Set("user", user)
		response.Set("permission", permission)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Printf(err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	}
}

func logout(w http.ResponseWriter, r *http.Request, c ClientData) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	response := simplejson.New()
	response.Set("success", true)
	response.Set("message", "You have sign out successfully.")

	payload, err := response.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func main() {

	dsn := "host=" + os.Getenv("HOST") + " user=" + os.Getenv("USER") + " password=" + os.Getenv("PASSWORD") + " dbname=" + os.Getenv("DATABASE") + " port=" + os.Getenv("PORT") + " sslmode=disable"
	client := NewConnect(dsn)

	r := mux.NewRouter()
	// For serving api endpoints
	api := r.PathPrefix("/v1").Subrouter()
	// For serving static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./files")))

	// posts
	api.HandleFunc("/{siteId}/posts", func(w http.ResponseWriter, r *http.Request) {
		getPosts(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/posts", func(w http.ResponseWriter, r *http.Request) {
		createPost(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/posts", func(w http.ResponseWriter, r *http.Request) {
		updatePost(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/{siteId}/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		getPostDetail(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		deletePost(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// setting
	api.HandleFunc("/{siteId}/setting", func(w http.ResponseWriter, r *http.Request) {
		getSetting(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/setting", func(w http.ResponseWriter, r *http.Request) {
		createSetting(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/setting/{settingId}", func(w http.ResponseWriter, r *http.Request) {
		updateSetting(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	// texts
	api.HandleFunc("/{siteId}/text", func(w http.ResponseWriter, r *http.Request) {
		getTexts(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/text", func(w http.ResponseWriter, r *http.Request) {
		createText(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/text", func(w http.ResponseWriter, r *http.Request) {
		updateText(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/{siteId}/text/{key}", func(w http.ResponseWriter, r *http.Request) {
		getText(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/text/{textId}", func(w http.ResponseWriter, r *http.Request) {
		deleteText(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// users
	api.HandleFunc("/{siteId}/users", func(w http.ResponseWriter, r *http.Request) {
		getUsers(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/users", func(w http.ResponseWriter, r *http.Request) {
		createUser(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		createRegistration(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/users", func(w http.ResponseWriter, r *http.Request) {
		updateUser(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/{siteId}/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		getUser(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		deleteUser(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	api.HandleFunc("/{siteId}/user/{username}", func(w http.ResponseWriter, r *http.Request) {
		getUserByEmail(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	// files
	api.HandleFunc("/{siteId}/files", func(w http.ResponseWriter, r *http.Request) {
		getFiles(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		updateFile(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/{siteId}/files/upload", func(w http.ResponseWriter, r *http.Request) {
		uploadFile(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		getFile(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		deleteFile(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// languages
	api.HandleFunc("/{siteId}/languages", func(w http.ResponseWriter, r *http.Request) {
		getLanguages(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/languages", func(w http.ResponseWriter, r *http.Request) {
		createLanguage(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/languages", func(w http.ResponseWriter, r *http.Request) {
		updateLanguage(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/{siteId}/languages/{code}", func(w http.ResponseWriter, r *http.Request) {
		getLanguageByCode(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/languages/{languageId}", func(w http.ResponseWriter, r *http.Request) {
		deleteLanguage(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// sites
	api.HandleFunc("/sites", func(w http.ResponseWriter, r *http.Request) {
		getSites(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/sites", func(w http.ResponseWriter, r *http.Request) {
		createSite(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/sites", func(w http.ResponseWriter, r *http.Request) {
		updateSite(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/sites/{siteId}", func(w http.ResponseWriter, r *http.Request) {
		deleteSite(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// auth
	api.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		auth(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/forgot/{username}", func(w http.ResponseWriter, r *http.Request) {
		forgot(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	// password recovery
	api.HandleFunc("/recovery", func(w http.ResponseWriter, r *http.Request) {
		recovery(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		me(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/refresh", func(w http.ResponseWriter, r *http.Request) {
		refresh(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		logout(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	// mail
	api.HandleFunc("/{siteId}/mail", func(w http.ResponseWriter, r *http.Request) {
		sendEmail(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":8888", r))
}
