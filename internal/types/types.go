package types

import "golang.org/x/oauth2"

type Config struct {
	GoogleAuthCfg *oauth2.Config
	PORT          string
	SHEET_ID      string
	GEMINI_API    string
	AUTHSTATE     string
}

type SheetResponse struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"majorDimension"`
	Values         [][]string `json:"values"`
}

type UserResponse struct {
	Email     string `json:"emailaddress"`
	Feedback  string `json:"feedback"`
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
}
