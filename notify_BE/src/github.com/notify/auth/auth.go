package auth

import (
	"context"
	"encoding/json"
	"github.com/notify/genaral/models"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JwtVerify Middleware function
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if (*r).Method == "OPTIONS" {
			return
		}
		resp := model.Response{}
	
		// set Default value
		resp.Default()

		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header2 := strings.TrimSpace(header)
		if header2 == "" {
			// Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			resp.Code = 403
			resp.Message = "Restricted no access token"
			json.NewEncoder(w).Encode(resp)
			return
		}
		tk := &Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtSecret")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			resp.Code = 403
			resp.Message = "Restricted : "+err.Error()
			json.NewEncoder(w).Encode(resp)
			log.Println("ERROR: Token encording  process is failed due to. ", err)
			return
		}

		log.Println("INFO: [AU]: Authorized User.")

		ctx := context.WithValue(r.Context(), "user", tk)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
