package controller

import (
	"net/http"
)

type LoginController struct{}

func NewLoginController() LoginController {
	return LoginController{}
}

func (lc LoginController) Show(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.ParseFiles("view/index.html")
	// if err != nil {
	// 	panic(err)
	// }

	// tmpl.Execute(w, nil)
}
