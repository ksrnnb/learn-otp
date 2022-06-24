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

func setCookie(w http.ResponseWriter, key string, value string) {
	c := &http.Cookie{
		Name:  key,
		Value: value,
		Path:  "/",
	}
	http.SetCookie(w, c)
}

func destroyCookies(w http.ResponseWriter, r *http.Request) {
	for _, c := range r.Cookies() {
		c.MaxAge = -1
		http.SetCookie(w, c)
	}
}
