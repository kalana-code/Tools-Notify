package genaral

import (
	"github.com/notify/auth"
	"github.com/notify/genaral/models"
	"encoding/json"
	"net/http"
	"os"
	"log"
	"github.com/dgrijalva/jwt-go"
)

// Verify used when User login
func Verify(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}


	resp := model.Response{}
	
	// set Default value
	resp.Default()
	
	var header = r.Header.Get("x-access-token")
	token := &auth.Token{}

	_, err := jwt.ParseWithClaims(header, token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwtSecret")), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		log.Println("ERROR: Token encording  process is failed due to . ", err)
		// set State
		resp.Code = 401
		resp.Message = "Unauthorized: "+err.Error()
	}else{
		// set State
		resp.Code = 200
		resp.Message = "OK"
		resp.Data = map[string]interface{}{"userInfo":token}
	}
	
	json.NewEncoder(w).Encode(token)

}
