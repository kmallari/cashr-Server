package main

import (
	"encoding/json"
	"github.com/MadAppGang/httplog"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/kmallari/cashr-Server/api"
	"github.com/kmallari/cashr-Server/auth"
	"github.com/kmallari/cashr-Server/db"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/supertokens"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.Use(httplog.LoggerWithConfig(httplog.LoggerConfig{
		RouterName:  "FillBodyFormatter",
		Formatter:   httplog.FullFormatterWithRequestAndResponseHeadersAndBody,
		CaptureBody: true,
	}))

	router.HandleFunc("/sessioninfo", session.VerifySession(nil, auth.SessionInfo)).Methods(http.MethodGet)
	router.HandleFunc("/userinfo", session.VerifySession(nil, auth.UserInfo)).Methods(http.MethodGet)
	router.HandleFunc("/health", checkHealth)

	api.Handler(router.PathPrefix("/api").Subrouter())

	log.Println("Server running @ port 8080")

	http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders(append([]string{"Content-Type"}, supertokens.GetAllCORSHeaders()...)),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowCredentials(),
	)(supertokens.Middleware(router)))
}

func checkHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Add("content-type", "application/json")

	err := db.Client.Ping()
	if err != nil {
		bytes, _ := json.Marshal(map[string]interface{}{
			"message": "Error connecting to database",
		})
		w.WriteHeader(500)
		w.Write(bytes)
		return
	}

	bytes, err := json.Marshal(map[string]interface{}{
		"status": "OK",
	})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error in converting to json"))
		return
	} else {
		w.Write(bytes)
		return
	}
}
