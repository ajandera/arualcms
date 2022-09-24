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

func GetLanguages(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		var languages []model.Language
		c.Db.Model(&model.Language{}).Where("site_id = ?", siteId).Scan(&languages)
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

func CreateLanguage(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
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
		c.Db.Model(&model.Site{}).Where("id = ?", siteId).Scan(&site)

		response := simplejson.New()

		c.Db.Create(&model.Language{
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

func UpdateLanguage(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
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
		c.Db.Model(&model.Site{}).Where("id = ?", siteId).Scan(&site)

		response := simplejson.New()

		c.Db.Model(model.Language{}).Where("id = ?", language.Id).Updates(model.Language{
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

func GetLanguageByCode(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])

	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		code := vars["code"]

		response := simplejson.New()

		var language model.Language
		c.Db.First(&model.Language{}, "code = ?", code).Scan(&language)
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

func DeleteLanguage(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		var l model.Language
		id := vars["languageId"]
		c.Db.Model(&model.Language{}).Where("id = ?", id).Delete(&l)
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
