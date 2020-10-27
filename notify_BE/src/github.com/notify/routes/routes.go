package routes

import (
	"github.com/notify/auth"
	"github.com/notify/genaral"
	app "github.com/notify/service/app"
	event "github.com/notify/service/events"
	subscriber "github.com/notify/service/subscribers"
	"net/http"

	"github.com/gorilla/mux"
)

// Handlers function Used for arrange routes
func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(CommonMiddleware)
	
	r.HandleFunc("/info", genaral.Information).Methods("GET", "OPTIONS")
	
	r.HandleFunc("/login", genaral.Login).Methods("POST", "OPTIONS")

	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	// use middleware
	s.Use(auth.JwtVerify)
	s.HandleFunc("/verify", genaral.Verify).Methods("GET", "OPTIONS")
	s.HandleFunc("/app", app.AddApplication).Methods("POST", "OPTIONS")

	// Tokens routes
	d := r.PathPrefix("/app").Subrouter()

	// use middleware
	d.Use(auth.AppVerify)
	d.HandleFunc("/", genaral.Information).Methods("GET", "OPTIONS")
	d.HandleFunc("/levels", app.GetApplicationLevels).Methods("GET", "OPTIONS")
	d.HandleFunc("/events", event.Events).Methods("POST", "OPTIONS")
	d.HandleFunc("/subscribe", subscriber.Subscribe).Methods("POST", "OPTIONS")

	d.HandleFunc("/subscriber/{level:[0-9]+}", subscriber.SubscribersListForGivenLevel).Methods("GET", "OPTIONS")
	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding,X-Access-Token")
		next.ServeHTTP(w, r)
	})
}
