package controller

import (
	"html/template"
	"net/http"
)

type IndexController struct{}

var tmplIndex *template.Template
var tmplLogin *template.Template
var tmplOTPLogin *template.Template

func init() {
	var err error
	tmplIndex, err = template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}

	tmplLogin, err = template.ParseFiles("view/login.html")
	if err != nil {
		panic(err)
	}

	tmplOTPLogin, err = template.ParseFiles("view/otp.html")
	if err != nil {
		panic(err)
	}
}

func NewIndexController() IndexController {
	return IndexController{}
}

func (lc IndexController) Show(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(w, r) {
		tmplIndex.Execute(w, nil)
		return
	}

	destroyCookies(w, r)

	// user is NOT logged in
	redirectToLogin(w, r)
}
