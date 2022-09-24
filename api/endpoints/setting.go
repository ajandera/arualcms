package endpoints

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"main/model"
	utils "main/utils"
	"net/http"
)

func GetSetting(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		var settings []model.Setting
		c.Db.Model(&model.Setting{}).Where("site_id = ?", siteId).Scan(&settings)
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

func CreateSetting(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
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
		c.Db.Model(&model.Site{}).Where("id = ?", siteId).Scan(&site)

		response := simplejson.New()

		c.Db.Create(&model.Setting{
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

func UpdateSetting(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
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
		c.Db.Model(model.Setting{}).Where("id = ?", setting.Id).Updates(model.Setting{
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
