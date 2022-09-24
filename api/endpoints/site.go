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

func GetSites(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	if auth, userId := utils.IsAuthorized(w, r, uuid.New(), c); auth == true {
		response := simplejson.New()

		var sites []model.Site
		c.Db.Model(&model.Site{}).Where("user_id = ?", userId).Scan(&sites)
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

func CreateSite(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	if auth, userId := utils.IsAuthorized(w, r, uuid.New(), c); auth == true {
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
		c.Db.Create(&res).Scan(&s)

		// generate permission
		var users []model.User
		c.Db.Model(&model.User{}).Where("parent_id = ?", userId).Scan(&users)
		for _, u := range users {
			c.Db.Create(&model.Permission{
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

func createDefaultSetting(siteId string, c utils.ClientData) {
	settings := [...]string{"subtitle", "author", "contactEmail", "contactPhone", "metaTitle", "metaDescription", "metaKeywords", "smtp", "smtp_user", "smtp_password", "smtp_port", "title"}
	for _, v := range settings {
		c.Db.Create(&model.Setting{
			Key:    v,
			Value:  "",
			SiteId: siteId,
		})
	}
}

func UpdateSite(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
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

		c.Db.Model(model.Site{}).Where("id = ?", site.Id).Updates(model.Site{
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

func DeleteSite(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		var s model.Site
		c.Db.Model(&model.Language{}).Where("id = ?", siteId).Delete(&s)
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
