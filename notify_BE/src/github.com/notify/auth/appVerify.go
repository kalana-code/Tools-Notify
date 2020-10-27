package auth

import (
	"context"
	"encoding/json"
	"github.com/notify/db"
	"github.com/notify/genaral/models"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// AppVerify Middleware function
func AppVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (*r).Method == "OPTIONS" {
			return
		}
		resp := model.Response{}

		// set Default value
		resp.Default()

		var header = r.Header.Get("x-app-access-token") //Grab the token from the header

		accessToken := strings.TrimSpace(header)
		if accessToken == "" {
			// Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			resp.Code = 403
			resp.Message = "Restricted no access token"
			json.NewEncoder(w).Encode(resp)
			return
		}

		_, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtSecret")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			resp.Code = 403
			resp.Message = "Restricted : " + err.Error()
			json.NewEncoder(w).Encode(resp)
			log.Println("ERROR: Token encording  process is failed due to. ", err)
			return
		}

		var applicationId string
		var applicationName string
		dbConnection, err := db.GetMySQLConnection()
		if err != nil {
			resp.Code = 500
			resp.Message = "Internal Server Error: " + err.Error()
			json.NewEncoder(w).Encode(resp)
			log.Println("ERROR: Error when establishing db connections.", err)
			return
		}
		stmtOut, err := dbConnection.Prepare("SELECT application_id,application_name FROM app_keys WHERE application_access_key	 = ?")
		if err != nil {
			resp.Code = 500
			resp.Message = "Internal Server Error: " + err.Error()
			json.NewEncoder(w).Encode(resp)
			log.Println("ERROR: Error when establishing db connections.", err)
			return
		}

		err = stmtOut.QueryRow(accessToken).Scan(
			&applicationId,
			&applicationName,
		)
		if err != nil {
			resp.Code = 500
			resp.Message = "Internal Server Error: " + err.Error()
			json.NewEncoder(w).Encode(resp)
			log.Println("ERROR: Error when executing query.", err)
			return
		}

		if applicationId == "" {
			resp.Code = 403
			resp.Message = "Restricted: No application for given access token"
			json.NewEncoder(w).Encode(resp)
			log.Println("ERROR: Application not found for given app-access token.")
			return
		}

		log.Println("INFO: [AU]: Valid Application.")

		ctx := context.WithValue(r.Context(), "appContext",
			map[string]string{
				"ApplicationID" :   applicationId,
				"ApplicationName" : applicationName,
			},
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
