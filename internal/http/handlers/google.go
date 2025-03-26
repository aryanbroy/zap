package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aryanbroy/zap/internal/types"
	"github.com/aryanbroy/zap/internal/utils/cookies"
	"github.com/aryanbroy/zap/internal/utils/gemini"
	"github.com/aryanbroy/zap/internal/utils/response"
	"github.com/aryanbroy/zap/internal/workflows/google"
	"golang.org/x/oauth2"
)

type SuccessResponse struct {
	Successful bool   `json:"successful"`
	Status     int    `json:"status"`
	Message    string `json:"message"`
}

func OAuthGoogleLogin(cfg *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := cfg.GoogleAuthCfg.AuthCodeURL(cfg.AUTHSTATE, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func OAuthGoogleCallback(cfg *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != cfg.AUTHSTATE {
			response.WriteJson(w, http.StatusUnauthorized, response.CustomError("State mismatched", http.StatusUnauthorized))
			return
		}

		code := r.URL.Query().Get("code")
		token, err := cfg.GoogleAuthCfg.Exchange(context.Background(), code)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err, http.StatusInternalServerError))
			return
		}

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
			response.WriteJson(w, http.StatusUnauthorized, response.CustomError("User not authorized", http.StatusUnauthorized))
			return
		}

		sheetData, err := google.FetchSheets(sheetId, accessToken)
		if err != nil {
			response.WriteJson(w, http.StatusUnauthorized, response.GeneralError(err, http.StatusUnauthorized))
			return
		}

		values := sheetData.Values

		if len(values) == 0 {
			response.WriteJson(w, http.StatusLengthRequired, response.CustomError("No data present in sheets", http.StatusLengthRequired))
			return
		}

		headers := values[0]

		var mappedValues []map[string]string

		for _, row := range values[1:] {
			rowMap := make(map[string]string)
			for i, val := range row {
				rowMap[strings.ReplaceAll(strings.ToLower(headers[i]), " ", "")] = val
			}
			mappedValues = append(mappedValues, rowMap)
		}

		response.WriteJson(w, http.StatusOK, mappedValues)
	}
}

func MailHandler(cfg *types.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var formResponses []types.UserResponse

		err := json.NewDecoder(r.Body).Decode(&formResponses)
		if err != nil {
			status := http.StatusBadRequest
			response.WriteJson(w, status, response.GeneralError(err, status))
			return
		}

		reply, err := gemini.GeminiResponse(cfg, formResponses[0].Feedback)
		if err != nil {
			status := http.StatusServiceUnavailable
			response.WriteJson(w, status, response.GeneralError(err, status))
			return
		}

		accessToken, err := cookies.GetCookie(r, "accessToken")

		if err != nil {
			status := http.StatusInternalServerError
			response.WriteJson(w, status, response.GeneralError(err, status))
			return
		}

		err = google.SendMail(accessToken)
		if err != nil {
			status := http.StatusInternalServerError
			response.WriteJson(w, status, response.GeneralError(err, status))
			return
		}

		response.WriteJson(w, http.StatusOK, reply)
	}
}
