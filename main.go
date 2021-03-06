package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// route handler packages below
	"github.com/coding-CEO/go-backend-test/routeHandlers/authenticationHandlers/googleAuthenticationHandler"
	"github.com/coding-CEO/go-backend-test/routeHandlers/homeHandler"
)

func main() {
	router := mux.NewRouter()

	// route handlers below
	router.HandleFunc("/", homeHandler.HomeHandler)
	router.HandleFunc("/googleGenerateUserOAuthCode", googleAuthenticationHandler.GoogleGenerateUserOAuthCode)
	router.HandleFunc("/auth/google/verifyUserOAuthCode", googleAuthenticationHandler.GoogleVerifyUserOAuthCode)
    
	// start the server
	fmt.Println("Server is Listening on port 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}