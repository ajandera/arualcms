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

type MenuNode struct {
	model.Menu
	Children []*MenuNode
}

func GetMenu(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		response := simplejson.New()
		response.Set("success", true)
		var item model.Menu
		err := c.Db.Model(&model.Menu{}).Where("root = TRUE AND site_id = ?", siteId).Find(&item).Error
		if err != nil {
			log.Fatalf(err.Error())
		}
		items, err := GetAllMenuItems(siteId.String(), c)
		if err != nil {
			response.Set("error", err.Error())
		}
		response.Set("menu", BuildMenuTree(items, item))

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func GetMenuRoot(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		var item model.Menu
		err := c.Db.Model(&model.Menu{}).Where("root = TRUE AND site_id = ?", siteId).Find(&item).Error
		if err != nil {
			log.Fatalf(err.Error())
		}
		response := simplejson.New()
		response.Set("success", true)
		response.Set("menu", item)

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func GetMenuPublic(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	apiToken := vars["apiToken"]
	if auth, siteId := utils.IsAuthorizedByApiKey(w, r, apiToken, c); auth == true {
		response := simplejson.New()
		var item model.Menu
		err := c.Db.Model(&model.Menu{}).Where("root = TRUE AND site_id = ?", siteId).Find(&item).Error
		if err != nil {
			log.Fatalf(err.Error())
		}
		items, _ := GetAllMenuItems(siteId, c)
		response.Set("success", true)
		response.Set("menu", BuildMenuTree(items, item))

		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func CreateMenu(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["siteId"])
	if auth, _, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		// Declare a new Menu struct.
		var menu decode.Menu

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&menu)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		response := simplejson.New()
		uuidParent, err := uuid.Parse(menu.ParentId)
		if err != nil {
			uuidParent = uuid.New()
		}
		c.Db.Create(&model.Menu{
			Name:     menu.Name,
			Url:      menu.Url,
			ParentId: uuidParent,
			SiteId:   siteId.String(),
			Root:     false,
		})
		response.Set("success", true)
		response.Set("message", "Menu created successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func UpdateMenu(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	menuId, _ := uuid.Parse(vars["menuId"])
	if auth, _, _ := utils.IsAuthorized(w, r, menuId, c); auth == true {

		// Declare a new Page struct.
		var menu decode.Menu

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&menu)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return
		}

		log.Println(menu.Id)
		response := simplejson.New()
		uuidParent, _ := uuid.Parse(menu.ParentId)
		c.Db.Model(&model.Menu{}).Where("id = ?", menu.Id).Updates(model.Menu{
			Name:     menu.Name,
			Url:      menu.Url,
			ParentId: uuidParent,
			SiteId:   menu.SiteId})

		response.Set("success", true)
		response.Set("message", "Menu updated successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Fatalf(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func DeleteMenu(w http.ResponseWriter, r *http.Request, c utils.ClientData) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	vars := mux.Vars(r)
	siteId, _ := uuid.Parse(vars["menuId"])
	if auth, _, _ := utils.IsAuthorized(w, r, siteId, c); auth == true {
		id, _ := uuid.Parse(vars["menuId"])
		e := c.Db.Delete(&model.Menu{}, id).Error
		if e != nil {
			log.Println(e.Error())
		}
		response := simplejson.New()
		response.Set("success", true)
		response.Set("message", "Menu deleted successfully.")
		w.WriteHeader(http.StatusOK)

		payload, err := response.MarshalJSON()
		if err != nil {
			log.Println(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

func GetAllMenuItems(siteId string, c utils.ClientData) ([]model.Menu, error) {
	var items []model.Menu
	err := c.Db.Model(&model.Menu{}).Where("root = FALSE AND site_id = ?", siteId).Order("parent_id ASC").Scan(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func BuildMenuTree(items []model.Menu, rootMenu model.Menu) []*MenuNode {
	// Create a map to look up nodes by ID
	nodes := make(map[string]*MenuNode)

	// Create the root node
	root := &MenuNode{}

	// Iterate over items and create a node for each
	for _, item := range items {
		node := &MenuNode{
			Menu: item,
		}
		nodes[item.Id.String()] = node

		// Check if this node is the root
		if item.ParentId == rootMenu.Id {
			root.Children = append(root.Children, node)
		} else {
			// Find the parent node and add this node as a child
			parent, ok := nodes[item.ParentId.String()]
			if !ok {
				// If the parent node hasn't been created yet, create it
				uuidNode, _ := uuid.Parse(item.ParentId.String())
				parent = &MenuNode{
					Menu: model.Menu{
						Id: uuidNode,
					},
				}
				nodes[item.ParentId.String()] = parent
			}
			parent.Children = append(parent.Children, node)
		}
	}

	// Return the root node's children
	return root.Children
}
