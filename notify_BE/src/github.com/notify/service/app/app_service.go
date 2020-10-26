package app

import (
	"encoding/json"
	model "github.com/notify/genaral/models"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Login used when User login
func AddApplication(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	var application Application
	reqBody, err := ioutil.ReadAll(r.Body)

	//get response model
	resp := model.Response{}

	// set Default value
	resp.Default()

	if err != nil {
		log.Println("ERROR: Payload Error", err)
		resp.BadRequest()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		err := json.Unmarshal(reqBody, &application)
		if err != nil {
			log.Println("ERROR: Payload Error", err)
			resp.BadRequest()
			w.WriteHeader(http.StatusBadRequest)
		} else {
			applicationIDs, err := InsertApp(application.ApplicationName)
			if err != nil {
				resp.Code = 500
				resp.Message = err.Error()
				w.WriteHeader(500)
			} else {
				resp.Code = 200
				resp.Message = "Successfully added application data"
				resp.Data = map[string]interface{}{
					"application_details": applicationIDs,
				}
				w.WriteHeader(200)
			}

		}
	}
	json.NewEncoder(w).Encode(resp)

}

// GetApplication Level this used for get application levels
func GetApplicationLevels(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	//get response model
	resp := model.Response{}
	var header = r.Header.Get("x-app-access-token") //Grab the token from the header

	accessToken := strings.TrimSpace(header)

	// set Default value
	resp.Default()
	levels, err := GetApplicationLevelsDetails(accessToken)

	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		w.WriteHeader(500)
	} else {
		resp.Code = 200
		resp.Message = "Successfully got application levels"
		resp.Data = map[string]interface{}{
			"application_levels":levels,
		}
		w.WriteHeader(200)
	}

	json.NewEncoder(w).Encode(resp)

}

