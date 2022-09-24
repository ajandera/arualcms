package endpoints

import (
	"github.com/bitly/go-simplejson"
	"log"
	utils "main/utils"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	response := simplejson.New()
	response.Set("success", true)

	payload, err := response.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
