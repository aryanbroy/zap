package main

import (
	"log"
	"net/http"

	"github.com/aryanbroy/zap/internal/config"
	"github.com/aryanbroy/zap/internal/http/handlers"
)

// get the access token
//      how to get access token? use oauth to get access token

// use this token to fetch details of users who submitted the form, from google sheets
//      how to fetch data from sheets? use their api "GET https://sheets.googleapis.com/v4/spreadsheets/{spreadsheetId}/values/{range}"
//      send a email to people who have submitted the form

func main() {
	cfg := config.MustLoad()
	// fmt.Println("Client Id: ", cfg.CLIENT_ID)
	// fmt.Println("Client Secret", cfg.CLIENT_SECRET)

	router := http.NewServeMux()

	server := http.Server{
		Addr:    cfg.PORT,
		Handler: router,
	}

	router.HandleFunc("/auth/google/login", handlers.OAuthGoogleLogin(cfg))
	router.HandleFunc("/auth/google/callback", handlers.OAuthGoogleCallback(cfg))
	router.HandleFunc("/admin", handlers.AdminControl())
	router.HandleFunc("/profile", handlers.AdminProfile())
	router.HandleFunc("/tokens", handlers.TokenControl())

	log.Println("Server started at port", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start the server ", err.Error())
	}
}
