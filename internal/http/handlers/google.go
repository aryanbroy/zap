package handlers

import (
	"context"
	"net/http"

	"github.com/aryanbroy/zap/internal/types"
	"github.com/aryanbroy/zap/internal/utils/cookies"
	"github.com/aryanbroy/zap/internal/utils/response"
	"github.com/aryanbroy/zap/internal/workflows/google"
	"golang.org/x/oauth2"
)

type SuccessResponse struct {
	Successful bool   `json:"successful"`
	Status     int    `json:"status"`
	Message    string `json:"message"`
}

var oauthState = "random-secret"

// var accessToken string

func OAuthGoogleLogin(cfg *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := cfg.GoogleAuthCfg.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func OAuthGoogleCallback(cfg *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != oauthState {
			response.WriteJson(w, http.StatusUnauthorized, response.CustomError("State mismatched", http.StatusUnauthorized))
			return
		}

		code := r.URL.Query().Get("code")
		token, err := cfg.GoogleAuthCfg.Exchange(context.Background(), code)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err, http.StatusInternalServerError))
			return
		}

		// accessToken = token.AccessToken

		// fmt.Println("Expiries in: ", token.ExpiresIn)
		// fmt.Println("Expiry: ", token.Expiry)

		cookies.ApplyCookie(w, token)

		res := SuccessResponse{
			Successful: true,
			Status:     http.StatusOK,
			Message:    "Authentication successful!",
		}
		response.WriteJson(w, http.StatusOK, res)
	}
}

func FormResponses(cfg *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sheetId := cfg.SHEET_ID

		accessToken, err := cookies.GetCookie(r, "accessToken")

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.CustomError("Error getting cookie", http.StatusInternalServerError))
			return
		}

		sheetData, err := google.FetchSheets(sheetId, accessToken)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err, http.StatusInternalServerError))
			return
		}

		response.WriteJson(w, http.StatusOK, sheetData)
	}
}
