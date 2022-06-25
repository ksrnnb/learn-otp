package controller

import (
	"net/http"
)

type LogoutController struct{}

func NewLogoutController() LogoutController {
	return LogoutController{}
}

func (lc LogoutController) Logout(w http.ResponseWriter, r *http.Request) {
	destroyCookies(w, r)
	redirectToIndex(w, r)
}
