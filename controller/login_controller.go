package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

type LoginController struct{}

func NewLoginController() LoginController {
	return LoginController{}
}

func (lc LoginController) Show(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/login.html")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, nil)
}

func (lc LoginController) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: redirect
	tmpl, err := template.ParseFiles("view/otp.html")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, nil)
}

func (lc LoginController) OTPLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: success -> index, fail -> otp.html
	fmt.Println("OTP Log in")
}
