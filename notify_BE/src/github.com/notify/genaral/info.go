package genaral

import (
	"fmt"
	"net/http"
)

//Information give details about service
func Information(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "OPTIONS" {
		return
	}
	fmt.Fprintf(w, "Welcome home!")
}
