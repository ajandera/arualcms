package endpoints

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"main/model"
	utils "main/utils"
	"net/http"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetFiles(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {

		response := simplejson.New()

		var files []model.File
		c.Db.Model(&model.File{}).Where("site_id = ?", siteId).Scan(&files)
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

func GetFilesPublic(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	apiToken := vars["siteId"]
	if auth, siteId := utils.IsAuthorizedByApiKey(w, r, apiToken, c); auth == true {

		response := simplejson.New()

		var files []model.File
		c.Db.Model(&model.File{}).Where("site_id = ?", siteId).Scan(&files)
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

func UpdateFile(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	fileId := vars["fileId"]
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
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

		c.Db.Model(model.File{}).Where("id = ?", fileId).Updates(model.File{
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

func UploadFile(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
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
		c.Db.Model(&model.Site{}).Where("id = ?", siteId).Scan(&site)

		var f model.File

		c.Db.Create(&model.File{
			Name:    handle.Filename,
			Src:     "/" + siteId.String() + "/" + handle.Filename,
			Gallery: "",
			SiteId:  siteId.String(),
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

func GetFile(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		vars := mux.Vars(r)
		fileId := vars["fileId"]

		response := simplejson.New()

		var file model.File
		c.Db.First(&model.File{}, "id = ?", fileId).Scan(&file)
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

func GetFilePublic(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	apiToken := vars["apiToken"]
	if auth, _ := utils.IsAuthorizedByApiKey(w, r, apiToken, c); auth == true {
		vars := mux.Vars(r)
		fileId := vars["fileId"]

		response := simplejson.New()

		var file model.File
		c.Db.First(&model.File{}, "id = ?", fileId).Scan(&file)
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

func DeleteFile(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		var f model.File
		vars := mux.Vars(r)
		id := vars["fileId"]
		c.Db.Model(&model.File{}).Where("id = ?", id).Delete(&f)
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
