package controller

import (
	"context"
	"net/http"

	"github.com/ksrnnb/otp/session"
)

const (
	sessionCookieName = "session_id"
	otpCookieName     = "otp_session_id"
)

func isLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	sid, err := r.Cookie(sessionCookieName)
	if err != nil {
		return false
	}

	c := session.NewClient()
	_, err = c.GetLoginSession(context.Background(), sid.Value)
	return err == nil
}

func isOTPLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	sid, err := r.Cookie(otpCookieName)

	if err != nil {
		return false
	}

	c := session.NewClient()
	_, err = c.GetOTPSession(context.Background(), sid.Value)
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
