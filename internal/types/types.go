package types

import "golang.org/x/oauth2"

type Config struct {
	GoogleAuthCfg *oauth2.Config
	PORT          string
}
