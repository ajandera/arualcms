package endpoints

import (
	"encoding/json"
	"log"
	"main/decode"
	"main/model"
	utils "main/utils"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetPages(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		var pages []model.Page
		c.Db.Model(&model.Page{}).Where("site_id = ?", siteId).Scan(&pages)
		response.Set("success", true)
		response.Set("pages", pages)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func GetPagesPublic(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	apiToken := vars["apiToken"]
	if auth, siteId := utils.IsAuthorizedByApiKey(w, r, apiToken, c); auth == true {
		response := simplejson.New()

		var pages []model.Page
		c.Db.Model(&model.Page{}).Where("site_id = ?", siteId).Scan(&pages)
		response.Set("success", true)
		response.Set("pages", pages)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func CreatePage(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Page struct.
		var page decode.Page

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()

		c.Db.Create(&model.Page{
			Key:         page.Key,
			Body:        page.Body,
			Title:       page.Title,
			MetaTitle:   page.MetaTitle,
			Keywords:    page.Keywords,
			Description: page.Description,
			SiteId:      siteId.String(),
		})
		response.Set("success", true)
		response.Set("message", "Page created successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func UpdatePage(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {

		// Declare a new Page struct.
		var page decode.Page

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		log.Println(page.Id)
		response := simplejson.New()
		c.Db.Model(&model.Page{}).Where("id = ?", page.Id).Updates(model.Page{
			Body:        page.Body,
			Title:       page.Title,
			MetaTitle:   page.MetaTitle,
			Keywords:    page.Keywords,
			Description: page.Description})

		response.Set("success", true)
		response.Set("message", "Page updated successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func GetPageByKeyPublic(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	apiToken := vars["apiToken"]

	if auth, siteId := utils.IsAuthorizedByApiKey(w, r, apiToken, c); auth == true {
		response := simplejson.New()
		key := vars["key"]

		var page model.Page
		c.Db.First(&model.Page{}, "site_id = ? AND key = ?", siteId, key).Scan(&page)
		response.Set("success", true)
		response.Set("page", page)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func DeletePage(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		id, _ := uuid.Parse(vars["pageId"])
		e := c.Db.Delete(&model.Page{}, id).Error
		if e != nil {
			log.Println(e.Error())
		}
		response := simplejson.New()
		response.Set("success", true)
		response.Set("message", "Page deleted successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
