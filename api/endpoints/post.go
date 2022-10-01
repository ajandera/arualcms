package endpoints

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"main/decode"
	"main/model"
	utils "main/utils"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		var posts []model.Post
		var files []model.File
		c.Db.Model(&model.Post{}).Where("site_id = ?", siteId).Scan(&posts)
		for _, v := range posts {
			var f model.File
			c.Db.Model(&model.File{}).Where("id = ?", v.File).Scan(&f)
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

func CreatePost(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Post struct.
		var post decode.Post

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()

		c.Db.Create(&model.Post{
			Body:        post.Body,
			Title:       post.Title,
			Excerpt:     post.Excerpt,
			Published:   post.Published,
			MetaTitle:   post.MetaTitle,
			Keywords:    post.Keywords,
			Description: post.Description,
			SiteId:      siteId.String(),
		})
		response.Set("success", true)
		response.Set("message", "Post created successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {

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
		c.Db.Model(model.Post{}).Where("id = ?", post.Id).Updates(model.Post{
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
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func GetPostDetail(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()

		postId := vars["postId"]

		var post model.Post
		c.Db.First(&model.Post{}, "id = ?", postId).Scan(&post)
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

func DeletePost(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		id := vars["postId"]
		e := c.Db.Delete(&model.Post{}, id).Error
		if e != nil {
			log.Println(e.Error())
		}
		response := simplejson.New()
		response.Set("success", true)
		response.Set("message", "Post deleted successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}
