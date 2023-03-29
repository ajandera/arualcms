package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"main/decode"
	utils "main/utils"
	"net/http"
	"os"
	"strings"

	"github.com/bitly/go-simplejson"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	var toTranslate decode.Deepl

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&toTranslate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf(err.Error())
		return
	}

	url := "https://api-free.deepl.com/v2/translate"
	params := strings.NewReader("text=" + toTranslate.Text + "&target_lang=" + toTranslate.Lang)
	req, _ := http.NewRequest("POST", url, params)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "DeepL-Auth-Key "+os.Getenv("DEEPL"))
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)

	response := simplejson.New()
	response.Set("success", true)
	response.Set("text", string(body))

	payload, err := response.MarshalJSON()
	if err != nil {
		log.Printf(err.Error())
		return
	}

	defer res.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
