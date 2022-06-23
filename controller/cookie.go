package controller

import "net/http"

const (
	sessionCookieName = "session_id"
	otpCookieName     = "otp_session_id"
)

func isLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie(sessionCookieName)

	return err == nil
}

func isOTPLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie(otpCookieName)

	return err == nil
}
