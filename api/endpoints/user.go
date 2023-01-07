package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"main/decode"
	"main/model"
	utils "main/utils"
	"net/http"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		var users []model.User
		c.Db.Model(&model.User{}).Scan(&users)
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

func CreateRegistration(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
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

	pw, hashErr := utils.HashPassword(user.Password)
	if hashErr != nil {
		log.Fatalf(hashErr.Error())
	}
	e := c.Db.Create(&model.User{
		Name:     user.Name,
		Username: user.Username,
		Password: pw,
		ParentId: "",
	}).Error
	if e != nil {
		response.Set("success", true)
		response.Set("message", e.Error())
	} else {
		response.Set("success", true)
		response.Set("message", "Your registration was successful.")
	}
	w.WriteHeader(http.StatusOK)

	payload, err := response.MarshalJSON()
	if err != nil {
		log.Fatalf(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func CreateUser(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		// Declare a new User struct.
		var user decode.User
		response := simplejson.New()

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			response.Set("success", false)
			response.Set("message", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			payload, err := response.MarshalJSON()
			if err != nil {
				log.Fatalf(err.Error())
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(payload)
			return
		}

		pw, hashErr := utils.HashPassword(user.Password)
		if hashErr != nil {
			log.Fatalf(hashErr.Error())
		}

		var u = model.User{}

		c.Db.Create(&model.User{
			Name:     user.Name,
			Username: user.Username,
			Password: pw,
			ParentId: user.ParentId,
		}).Scan(&u)

		// generate permission
		var sites []model.Site
		c.Db.Model(&model.Site{}).Where("user_id = ?", user.ParentId).Scan(&sites)
		for _, v := range sites {
			c.Db.Create(&model.Permission{
				UserId: u.Id.String(),
				SiteId: v.Id.String(),
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

func UpdateUser(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
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

		c.Db.Model(model.User{}).Where("id = ?", user.Id).Updates(model.User{
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

func GetUser(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		vars := mux.Vars(r)
		userId := vars["userId"]

		response := simplejson.New()

		var user model.User
		c.Db.First(&model.User{}, "id = ?", userId).Scan(&user)
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

func DeleteUser(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		var u model.User
		vars := mux.Vars(r)
		id := vars["userId"]
		c.Db.Model(&model.User{}).Where("id = ?", id).Delete(&u)
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

func GetUserByEmail(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		vars := mux.Vars(r)
		username := vars["username"]

		response := simplejson.New()

		var user model.User
		c.Db.First(&model.User{}, "username = ?", username).Scan(&user)
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

func Me(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	if auth, _ := utils.IsAuthorized(w, r, uuid.New(), c); auth == true {
		var sendToken = strings.Replace(r.Header["Authorization"][0], "Bearer ", "", 1)
		claims, err := utils.GetClaim(sendToken)
		response := simplejson.New()
		var user model.User
		var permission []model.Permission
		c.Db.First(&model.User{}, "id = ?", fmt.Sprintf("%v", claims["id"])).Scan(&user)
		c.Db.Model(&model.Permission{}).Where("user_id = ?", fmt.Sprintf("%v", claims["id"])).Scan(&permission)
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
