package services

import (
	"encoding/json"
	model "github.com/notify/genaral/models"
	"net/http"
)

func Events(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	//get response model
	resp := model.Response{}
	//var header = r.Header.Get("x-app-access-token") //Grab the token from the header

	//accessToken := strings.TrimSpace(header)

	// set Default value
	resp.Default()

	json.NewEncoder(w).Encode(resp)

}