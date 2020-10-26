package subscriber

import (
	"encoding/json"
	"fmt"
	model "github.com/notify/genaral/models"
	"io/ioutil"
	"log"
	"net/http"
)

//Subscribe this is used for subscribe application
func Subscribe(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}

	var subscriberRequest SubscribeRequest
	reqBody, err := ioutil.ReadAll(r.Body)

	// get applicationId from request contex
	applicationId:=r.Context().Value("applicationId")
	fmt.Print(applicationId)

	//get response model
	resp := model.Response{}

	// set Default value
	resp.Default()

	if err != nil {
		log.Println("ERROR: Payload Error", err)
		resp.BadRequest()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		err := json.Unmarshal(reqBody, &subscriberRequest)
		if err != nil {
			log.Println("ERROR: Payload Error", err)
			resp.BadRequest()
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err := SubscribeApp(subscriberRequest)
			if err != nil {
				resp.Code = 500
				resp.Message = err.Error()
				w.WriteHeader(500)
			} else {
				resp.Code = 200
				resp.Message = "Successfully subscribe application"
				w.WriteHeader(200)
			}
		}
	}
	json.NewEncoder(w).Encode(resp)

}