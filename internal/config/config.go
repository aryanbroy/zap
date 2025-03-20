package config

import (
	"log"
	"os"

	"github.com/aryanbroy/zap/internal/types"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func MustLoad() *types.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading env file", err)
		return nil
	}
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectUrl := os.Getenv("REDIRECT_URI")
	port := os.Getenv("PORT")
	sheetId := os.Getenv("SHEET_ID")
	geminiApi := os.Getenv("GEMINI_API")

	googleOAuthConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
		Scopes: []string{
			"https://www.googleapis.com/auth/forms.responses.readonly",
			"https://www.googleapis.com/auth/spreadsheets",
			"https://www.googleapis.com/auth/gmail.send",
		},
		Endpoint: google.Endpoint,
	}

	response := types.Config{
		GoogleAuthCfg: googleOAuthConfig,
		PORT:          port,
		SHEET_ID:      sheetId,
		GEMINI_API:    geminiApi,
	}
	return &response
}
