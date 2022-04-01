package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type tokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}

type Pw struct {
    Password    string
}

var mySigningKey = []byte("DFGDFGhcsadkjhfwe+ě+23123asldxjhsdljfh1234234")

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func isAuthorized(w http.ResponseWriter, r *http.Request, m model.Repository, storeId string) bool {
	if r.Header["Authorization"] != nil {
		var sendToken = strings.Replace(r.Header["Authorization"][0], "Bearer ", "", 1)
		token, err := jwt.Parse(sendToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("JWT token not pass")
			}
			return mySigningKey, nil
		})

		// check if store belongs to account
		if IsValidUUID(storeId) {
			claims := token.Claims.(jwt.MapClaims)
			if m.IsPermitted(fmt.Sprintf("%v", claims["id"]), storeId) == false {
				http.Error(w, "Not Permitted", http.StatusForbidden)
				return false
			}
		}

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

func setupCORS(w *http.ResponseWriter, req *http.Request) {
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

func refresh(w http.ResponseWriter, r *http.Request, m model.Repository) {
	setupCORS(&w, r)
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
		var account rdbsClientInfo.Accounts = m.GetAccountById(fmt.Sprint(claims["id"]))
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

func auth(w http.ResponseWriter, r *http.Request, m model.Repository) {
	setupCORS(&w, r)
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

func main() {

    client, err := mongo.NewClient(options.Client().ApplyURI("<ATLAS_URI_HERE>"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

    err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	r := mux.NewRouter()
	api := r.PathPrefix("/v1").Subrouter()

	// posts
	api.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		posts(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		getAccounts(w, r, repository)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		updateAccounts(w, r, repository)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		deleteAccounts(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/posts/{postId}", func(w http.ResponseWriter, r *http.Request) {
		getAccount(w, r, repository)
	}).Methods(http.MethodDelete, http.MethodOptions)

    // setting
	api.HandleFunc("/setting", func(w http.ResponseWriter, r *http.Request) {
		getAccountByEmail(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/setting", func(w http.ResponseWriter, r *http.Request) {
		addOpenData(w, r, repository)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/setting", func(w http.ResponseWriter, r *http.Request) {
		getOpenData(w, r, repository)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/setting/{settingId}", func(w http.ResponseWriter, r *http.Request) {
		stores(w, r, repository)
	}).Methods(http.MethodPost, http.MethodOptions)

    // texts
	api.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		getStores(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		updateStores(w, r, repository)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		deleteStores(w, r, repository)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/text/{textId}", func(w http.ResponseWriter, r *http.Request) {
		getPredictedVisitors(w, r, repository, influx)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/text/{textId}", func(w http.ResponseWriter, r *http.Request) {
		getPredictedOrders(w, r, repository, influx)
	}).Methods(http.MethodDelete, http.MethodOptions)

	// users
	api.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		createStoreWeights(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		updateStoreWeights(w, r, repository)
	}).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		getStoreWeights(w, r, repository)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		plans(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/user/{userId}", func(w http.ResponseWriter, r *http.Request) {
		updatePlans(w, r, repository)
	}).Methods(http.MethodDelete, http.MethodOptions)

	api.HandleFunc("/user/{username}", func(w http.ResponseWriter, r *http.Request) {
		getPlans(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

    // files
	api.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		deletePlans(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

    api.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		deletePlans(w, r, repository)
	}).Methods(http.MethodPut, http.MethodOptions)

    api.HandleFunc("/files/upload", func(w http.ResponseWriter, r *http.Request) {
		deletePlans(w, r, repository)
	}).Methods(http.MethodPost, http.MethodOptions)

    api.HandleFunc("/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		deletePlans(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/files/{fileId}", func(w http.ResponseWriter, r *http.Request) {
		deletePlans(w, r, repository)
	}).Methods(http.MethodDelete, http.MethodOptions)


    // languages
	api.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		deletePlans(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

    api.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		deletePlans(w, r, repository)
	}).Methods(http.MethodPost, http.MethodOptions)

    api.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		deletePlans(w, r, repository)
	}).Methods(http.MethodPut, http.MethodOptions)

    api.HandleFunc("/languages/{code}", func(w http.ResponseWriter, r *http.Request) {
		deletePlans(w, r, repository)
	}).Methods(http.MethodGet, http.MethodOptions)

    api.HandleFunc("/languages/{languageId}", func(w http.ResponseWriter, r *http.Request) {
		deletePlans(w, r, repository)
	}).Methods(http.MethodDelete, http.MethodOptions)


	// auth
	api.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		auth(w, r, repository)
	}).Methods(http.MethodPost, http.MethodOptions)

    // password recovery
	api.HandleFunc("/recovery", func(w http.ResponseWriter, r *http.Request) {
		refresh(w, r, repository)
	}).Methods(http.MethodPut, http.MethodOptions)

	api.HandleFunc("/recovery", func(w http.ResponseWriter, r *http.Request) {
		healthCheck(w, r, repository)
	}).Methods(http.MethodPost, http.MethodOptions)

	// mail
	api.HandleFunc("/mail", func(w http.ResponseWriter, r *http.Request) {
		getOrders(w, r, repository)
	}).Methods(http.MethodPost, http.MethodOptions)

	log.Fatal(http.ListenAndServe(":8888", r))
}

