package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type tokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}

type Pw struct {
	Password string
}

type Account struct {
	Username string
	Password string
}

var mySigningKey = []byte("DFGDFGhcsadkjhfwe+ě+23123asldxjhsdljfh1234234")

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func isAuthorized(w http.ResponseWriter, r *http.Request) bool {
	if r.Header["Authorization"] != nil {
		var sendToken = strings.Replace(r.Header["Authorization"][0], "Bearer ", "", 1)
		token, err := jwt.Parse(sendToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("JWT token not pass")
			}
			return mySigningKey, nil
		})

		if err != nil && token != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return false
		}
		return true
	} else {
		http.Error(w, "Not Authorized", http.StatusForbidden)
		return false
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

func refresh(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
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
		var account Account = m.Database("test").Collection(fmt.Sprint(claims["id"]))
		if account.Id != "" {

			newTokenPair, err := GenerateJWT(account.Name, account.Id)
			if err != nil {
				return
			}

			response := simplejson.New()
			if IsValidUUID(account.Id) == true && err == nil {
				response.Set("success", true)
				response.Set("account", account)
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

func auth(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
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

	account := m.Auth(a.Email, a.Password)
	token, err := GenerateJWT(account.Name, account.Id)
	if IsValidUUID(account.Id) == true && err == nil {
		response.Set("success", true)
		response.Set("account", account)
		response.Set("jwt", token)
		w.WriteHeader(http.StatusOK)
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

func getPosts(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func createPost(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func updatePost(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getPostDetail(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func deletePost(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getSetting(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func createSetting(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func updateSettings(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func updateSetting(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getTexts(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func createText(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func updateText(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getText(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func deleteText(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func createUser(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func updateUser(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getUserByEmail(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getFiles(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func updateFile(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func uploadFile(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getFile(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func deleteFile(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getLanguages(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func createLanguage(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func updateLanguage(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func getLanguageByCode(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func deleteLanguage(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func sendEmail(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func recovery(w http.ResponseWriter, r *http.Request, m *mongo.Client) {
	setupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
}

func main() {

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r := mux.NewRouter()
	api := r.PathPrefix("/v1").Subrouter()

	// posts
	api.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		getPosts(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		createPost(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		updatePost(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		getPostDetail(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		deletePost(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// setting
	api.HandleFunc("/setting", func(w http.ResponseWriter, r *http.Request) {
		getSetting(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/setting", func(w http.ResponseWriter, r *http.Request) {
		createSetting(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/setting", func(w http.ResponseWriter, r *http.Request) {
		updateSettings(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/setting/{settingId}", func(w http.ResponseWriter, r *http.Request) {
		updateSetting(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	// texts
	api.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		getTexts(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		createText(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		updateText(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/text/{textId}", func(w http.ResponseWriter, r *http.Request) {
		getText(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/text/{textId}", func(w http.ResponseWriter, r *http.Request) {
		deleteText(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// users
	api.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		getUsers(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		createUser(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		updateUser(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		getUser(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		deleteUser(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	api.HandleFunc("/user/{username}", func(w http.ResponseWriter, r *http.Request) {
		getUserByEmail(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	// files
	api.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		getFiles(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		updateFile(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/files/upload", func(w http.ResponseWriter, r *http.Request) {
		uploadFile(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		getFile(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		deleteFile(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// languages
	api.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		getLanguages(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		createLanguage(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		updateLanguage(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/languages/{code}", func(w http.ResponseWriter, r *http.Request) {
		getLanguageByCode(w, r, client)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/languages/{languageId}", func(w http.ResponseWriter, r *http.Request) {
		deleteLanguage(w, r, client)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// auth
	api.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		auth(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/recovery", func(w http.ResponseWriter, r *http.Request) {
		refresh(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	// password recovery
	api.HandleFunc("/recovery", func(w http.ResponseWriter, r *http.Request) {
		recovery(w, r, client)
	}).Methods(http.MethodPut, http.MethodOptions)

	// mail
	api.HandleFunc("/mail", func(w http.ResponseWriter, r *http.Request) {
		sendEmail(w, r, client)
	}).Methods(http.MethodPost, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":8888", r))
}
