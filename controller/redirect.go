package controller

import "net/http"

func redirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func redirectToOTPLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login/otp", http.StatusFound)
}

func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}
