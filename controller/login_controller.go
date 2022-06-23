package controller

import (
	"fmt"
	"net/http"
)

type LoginController struct{}

func NewLoginController() LoginController {
	return LoginController{}
}

func (lc LoginController) Show(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(w, r) {
		redirectToIndex(w, r)
		return
	}

	if isOTPLoggedIn(w, r) {
		redirectToOTPLogin(w, r)
	}

	tmplLogin.Execute(w, nil)
}

func (lc LoginController) Login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login/otp", http.StatusFound)
}

func (lc LoginController) ShowOTPLogin(w http.ResponseWriter, r *http.Request) {
	if !isOTPLoggedIn(w, r) {
		redirectToLogin(w, r)
	}

	tmplOTPLogin.Execute(w, nil)
}

func (lc LoginController) OTPLogin(w http.ResponseWriter, r *http.Request) {
	// TODO: success -> index, fail -> otp.html
	fmt.Println("OTP Log in")
	redirectToIndex(w, r)
}
