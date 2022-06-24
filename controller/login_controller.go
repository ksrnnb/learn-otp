package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ksrnnb/otp/model"
	"github.com/ksrnnb/otp/session"
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
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}
	id := r.FormValue("id")
	pwd := r.FormValue("password")

	user := model.FindUserById(id)
	if user == nil {
		redirectToIndex(w, r)
		return
	}

	if !user.EqualsPassword(pwd) {
		redirectToIndex(w, r)
		return
	}

	c := session.NewClient()
	s, err := c.CreateOTPSession(context.Background(), id)
	if err != nil {
		redirectToIndex(w, r)
		return
	}

	setCookie(w, otpCookieName, s)
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
