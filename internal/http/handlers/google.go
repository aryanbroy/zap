package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aryanbroy/zap/internal/types"
	"golang.org/x/oauth2"
)

var oauthState = "random-secret"

func OAuthGoogleLogin(cfg *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := cfg.GoogleAuthCfg.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func OAuthGoogleCallback(cfg *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != oauthState {
			http.Error(w, "State mismatch", http.StatusUnauthorized)
			return
		}

		code := r.URL.Query().Get("code")
		token, err := cfg.GoogleAuthCfg.Exchange(context.Background(), code)
		if err != nil {
			http.Error(w, "Failed to get token", http.StatusInternalServerError)
			return
		}
		fmt.Println("Access token: ", token.AccessToken)
		fmt.Println("Refresh token: ", token.RefreshToken)
		fmt.Println("Token type: ", token.TokenType)
		w.Write([]byte("Authentication successfull"))
	}
}

func AdminControl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func AdminProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func TokenControl() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
