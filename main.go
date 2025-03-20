package main

import (
	"log"
	"net/http"

	"github.com/aryanbroy/zap/internal/config"
	"github.com/aryanbroy/zap/internal/http/handlers"
	"github.com/aryanbroy/zap/internal/utils/gemini"
)

// done
// get the access token
//      how to get access token? use oauth to get access token

// working...
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

	router.HandleFunc("GET /auth/google/login", handlers.OAuthGoogleLogin(cfg))
	router.HandleFunc("GET /auth/google/callback", handlers.OAuthGoogleCallback(cfg))
	router.HandleFunc("GET /api/form-responses", handlers.FormResponses(cfg))

	gemini.GeminiResponse(cfg)

	log.Println("Server started at port", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start the server ", err.Error())
	}
}
