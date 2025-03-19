package cookies

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

func ApplyCookie(w http.ResponseWriter, token *oauth2.Token) {
	cookie := http.Cookie{
		Name:    "accessToken",
		Value:   token.AccessToken,
		Path:    "/",
		Expires: token.Expiry,
	}

	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request, cookieName string) (string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			log.Fatalln("Cookie not found")
		default:
			log.Fatalln("Server error")
		}
		return "", err
	}

	fmt.Println("Value of cookie: ", cookie.Value)

	return cookie.Value, nil
}
