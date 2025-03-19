package types

import "golang.org/x/oauth2"

type Config struct {
	GoogleAuthCfg *oauth2.Config
	PORT          string
	SHEET_ID      string
}

type SheetResponse struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"majorDimension"`
	Values         [][]string `json:"values"`
}
