package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/supertokens/supertokens-golang/supertokens"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//apiDomain := os.Getenv("API_DOMAIN")
	clientDomain := os.Getenv("CLIENT_DOMAIN")

	//supertokensDomain := os.Getenv("SUPERTOKENS_DOMAIN")
	//superTokensApiKey := os.Getenv("SUPERTOKENS_API_KEY")

	router := mux.NewRouter()

	// Adding handlers.CORS(options)(supertokens.Middleware(router)))
	if err := http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedHeaders(append([]string{"Content-Type"},
			supertokens.GetAllCORSHeaders()...)),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{clientDomain}),
		handlers.AllowCredentials(),
	)(supertokens.Middleware(router))); err != nil {
		fmt.Println(err)
	}

	log.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, "Running HTTP Server on Port [:8080]"))
}
