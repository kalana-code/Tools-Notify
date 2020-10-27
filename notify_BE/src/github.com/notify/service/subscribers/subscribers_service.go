package subscriber

import (
	"encoding/json"
	"github.com/gorilla/mux"
	model "github.com/notify/genaral/models"
	"github.com/notify/logger"
	"io/ioutil"
	"net/http"
	"strconv"
)

//Subscribe this is used for subscribe application
func Subscribe(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	//get response model
	resp := model.Response{}

	// set Default value
	resp.Default()
	// get applicationId from request context
	appContext, err := getAppContext(r.Context().Value("appContext"))
	if err != nil {
		logger.ErrorPrint(component_name, "Application Context Issue", err)
		resp.RestrictedError()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		var subscriberRequest SubscribeRequest
		reqBody, err := ioutil.ReadAll(r.Body)

		if err != nil {
			logger.ErrorPrint(component_name, "Bad Request", err)
			resp.BadRequest()
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err := json.Unmarshal(reqBody, &subscriberRequest)
			if err != nil {
				logger.ErrorPrint(component_name, "Payload error", err)
				resp.BadRequest()
				w.WriteHeader(http.StatusBadRequest)
			} else {
				err := SubscribeApp(subscriberRequest, appContext.ApplicationID)
				if err != nil {
					logger.ErrorPrint(component_name, "Internal Server Error ", err)
					resp.Code = 500
					resp.Message = err.Error()
					w.WriteHeader(500)
				} else {
					logger.InfoPrint(component_name, "Successfully subscribe the application. appID: "+appContext.ApplicationName)
					resp.Code = 200
					resp.Message = "Successfully subscribe application."
					resp.Data = map[string]interface{}{
						"application_id":  appContext.ApplicationID,
						"aplication_name": appContext.ApplicationName,
					}
					w.WriteHeader(200)
				}
			}
		}
	}

	json.NewEncoder(w).Encode(resp)

}

// dont used this function for event handle
func SubscribersListForGivenLevel(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	//get response model
	resp := model.Response{}

	// set Default value
	resp.Default()
	// get applicationId from request context
	appContext, err := getAppContext(r.Context().Value("appContext"))
	if err != nil {
		logger.ErrorPrint(component_name, "Application Context Issue", err)
		resp.RestrictedError()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		parameters := mux.Vars(r)
		applicationLevel, err := strconv.Atoi(parameters["level"])
		if err != nil {
			logger.ErrorPrint(component_name, "Internal Server Error ", err)
			resp.Code = 500
			resp.Message = err.Error()
			w.WriteHeader(500)
		} else {
			list,err := SubscribersList(applicationLevel, appContext.ApplicationID)
			if err != nil {
				logger.ErrorPrint(component_name, "Internal Server Error ", err)
				resp.Code = 500
				resp.Message = err.Error()
				w.WriteHeader(500)
			} else {
				logger.InfoPrint(component_name, "Successfully retrieved subscribers list : "+appContext.ApplicationName)
				resp.Code = 200
				resp.Message = "Successfully, retrieved subscribers list."
				resp.Data = map[string]interface{}{
					"application_id":  appContext.ApplicationID,
					"aplication_name": appContext.ApplicationName,
					"subscriber_list": list,
				}
				w.WriteHeader(200)
			}
		}

	}

	json.NewEncoder(w).Encode(resp)

}
