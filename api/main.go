package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"log"
	endpoints "main/endpoints"
	model "main/model"
	utils "main/utils"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"time"
)

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

func NewConnect(dsn string) utils.ClientData {
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
		&model.Site{},
		&model.Permission{})
	return utils.ClientData{db}
}

func refresh(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
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

	if claims, err := utils.GetClaim(tokenReq.RefreshToken); err == nil {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		var user model.User
		c.Db.First(&model.User{}, "id = ?", fmt.Sprint(claims["id"])).Scan(&user)
		if utils.IsValidUUID(user.Id) {
			newTokenPair, err := utils.GenerateJWT(user.Name, user.Id.String())
			if err != nil {
				return
			}

			response := simplejson.New()
			if utils.IsValidUUID(user.Id) == true && err == nil {
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

func auth(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
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
	c.Db.First(&model.User{}, "username = ?", a.Username).Scan(&account)
	pw := utils.CheckPasswordHash(a.Password, account.Password)
	if pw == true {
		token, err := utils.GenerateJWT(account.Name, account.Id.String())
		if utils.IsValidUUID(account.Id) == true && err == nil {
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

func sendEmail(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
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

func forgot(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	rand.Seed(time.Now().Unix())
	minSpecialChar := 1
	minNum := 1
	minUpperCase := 1
	passwordLength := 8
	token := utils.GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase)

	vars := mux.Vars(r)
	username := vars["username"]

	var user model.User
	c.Db.First(&model.User{}, "username = ?", username).Scan(&user)

	t := time.Now().Add(time.Hour * 24)
	c.Db.Model(model.User{}).Where("username = ?", username).Updates(model.User{
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

func recovery(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
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
	c.Db.Model(&model.User{}).Where("forgot_token = ?", reset.Token).First(&a)

	t1 := time.Now()
	t2 := a.ValidToken
	hourDiff := t2.Sub(t1).Hours()

	if hourDiff < 24 {
		a.ForgotToken = ""
		if len(reset.Password) > 6 {
			hash, _ := utils.HashPassword(reset.Password)
			a.Password = hash
		}
	}
	c.Db.Save(&a)

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

func sendEmailWithTemplate(email string, subject string, templateName string, token string, c utils.ClientData, r *http.Request) {
	vars := mux.Vars(r)
	siteId := vars["siteId"]
	var settings []model.Setting
	c.Db.Model(&model.Setting{}).Where("site_id = ?", siteId).Scan(&settings)
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

func sendEmailWithoutTemplate(email string, subject string, htmlString string, c utils.ClientData, r *http.Request) {

	vars := mux.Vars(r)
	siteId := vars["siteId"]
	var settings []model.Setting
	c.Db.Model(&model.Setting{}).Where("site_id = ?", siteId).Scan(&settings)
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

func logout(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
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
		endpoints.GetPosts(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/posts", func(w http.ResponseWriter, r *http.Request) {
		endpoints.CreatePost(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/posts", func(w http.ResponseWriter, r *http.Request) {
		endpoints.UpdatePost(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/{siteId}/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetPostDetail(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.DeletePost(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// setting
	api.HandleFunc("/{siteId}/setting", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetSetting(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/setting", func(w http.ResponseWriter, r *http.Request) {
		endpoints.CreateSetting(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/setting/{settingId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.UpdateSetting(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	// texts
	api.HandleFunc("/{siteId}/text", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetTexts(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/text", func(w http.ResponseWriter, r *http.Request) {
		endpoints.CreateText(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/text", func(w http.ResponseWriter, r *http.Request) {
		endpoints.UpdateText(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/{siteId}/text/{key}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetText(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/text/{textId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.DeleteText(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// users
	api.HandleFunc("/{siteId}/users", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetUsers(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/users", func(w http.ResponseWriter, r *http.Request) {
		endpoints.CreateUser(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		endpoints.CreateRegistration(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/users", func(w http.ResponseWriter, r *http.Request) {
		endpoints.UpdateUser(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/{siteId}/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetUser(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.DeleteUser(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	api.HandleFunc("/{siteId}/user/{username}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetUserByEmail(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	// files
	api.HandleFunc("/{siteId}/files", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetFiles(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.UpdateFile(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/{siteId}/files/upload", func(w http.ResponseWriter, r *http.Request) {
		endpoints.UploadFile(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetFile(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.DeleteFile(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// languages
	api.HandleFunc("/{siteId}/languages", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetLanguages(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/languages", func(w http.ResponseWriter, r *http.Request) {
		endpoints.CreateLanguage(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/{siteId}/languages", func(w http.ResponseWriter, r *http.Request) {
		endpoints.UpdateLanguage(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/{siteId}/languages/{code}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetLanguageByCode(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/{siteId}/languages/{languageId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.DeleteLanguage(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// sites
	api.HandleFunc("/sites", func(w http.ResponseWriter, r *http.Request) {
		endpoints.GetSites(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/sites", func(w http.ResponseWriter, r *http.Request) {
		endpoints.CreateSite(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/sites", func(w http.ResponseWriter, r *http.Request) {
		endpoints.UpdateSite(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/sites/{siteId}", func(w http.ResponseWriter, r *http.Request) {
		endpoints.DeleteSite(w, r, client)
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
		endpoints.Me(w, r, client)
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

	// health check
	api.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		endpoints.HealthCheck(w, r)
	}).Methods(http.MethodGet, http.MethodOptions)

	// add swagger
	if os.Getenv("API_DOC") == "true" {
		r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	}

	log.Fatal(http.ListenAndServe(":8888", r))
}
