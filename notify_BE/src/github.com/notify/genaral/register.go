package genaral

import (
	"github.com/notify/genaral/models"
	"github.com/notify/genaral/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Register used for register a new user
func Register(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	var TempUser model.UserInfo
	TempUser.Role = "a"
	reqBody, err := ioutil.ReadAll(r.Body)
	// message on Response
	resp := model.Response{}
	// set Default value
	resp.Default()

	if err != nil {
		log.Println("ERROR: Payload Error . ", err)
		resp.BadRequest()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		err2 := json.Unmarshal(reqBody, &TempUser)
		if err2 == nil {
			pass, errCrpto := bcrypt.GenerateFromPassword([]byte(TempUser.Password), bcrypt.DefaultCost)
			if errCrpto != nil {
				log.Println("ERROR: Decryption process has been failed. ", errCrpto)
				resp.InternalServerError()
				w.WriteHeader(http.StatusInternalServerError)

			} else {
				// set encripted passward
				TempUser.Password = string(pass)
				insertError := utils.InsertUser(TempUser)
				// set response message
				if resp.Code =200; insertError!=nil {
					resp.Code =500
				} 
				
				if resp.Message ="Insert pocess has been completed successfully"; insertError!=nil {
					resp.Message =" Inserting Process has been failed"
				} 

				w.WriteHeader(resp.Code)
				json.NewEncoder(w).Encode(resp)
			}
		} else {
			log.Println("ERROR: Payload Error . ", err)
			resp.BadRequest()
			w.WriteHeader(http.StatusBadRequest)
		}

	}

}
