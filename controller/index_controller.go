package controller

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/ksrnnb/otp/model"
	"github.com/ksrnnb/otp/session"
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
	if !isLoggedIn(w, r) {
		destroyCookies(w, r)
		redirectToLogin(w, r)
		return
	}

	sid, err := r.Cookie(sessionCookieName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	c := session.NewClient()
	userId, err := c.GetLoginSession(context.Background(), sid.Value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	user := model.FindUserById(userId)
	tmplIndex.Execute(w, map[string]interface{}{"id": user.Id()})
}
