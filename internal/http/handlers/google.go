package handlers

import "net/http"

func OAuthGoogleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// perform user login here
	}
}

func OAuthGoogleCallback() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// handle callback
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
