package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type IndexController struct{}

var tmplIndex *template.Template

func init() {
	var err error
	tmplIndex, err = template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}
}

func NewIndexController() IndexController {
	return IndexController{}
}

func (lc IndexController) Show(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie(sessionCookieName)

	// user is logged in
	if err == nil {
		tmplIndex.Execute(w, nil)
		return
	}

	// unexpected error
	if err != http.ErrNoCookie {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		return
	}

	// no session_id cookie
	// search otp_session_id cookie
	_, err = r.Cookie(otpCookieName)

	// user is otp logged in
	if err == nil {
		http.Redirect(w, r, "/login/otp", http.StatusFound)
		return
	}

	// unexpected error
	if err != http.ErrNoCookie {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		return
	}

	// user is NOT logged in
	http.Redirect(w, r, "/login", http.StatusFound)
}
