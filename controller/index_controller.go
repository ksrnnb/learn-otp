package controller

import (
	"html/template"
	"net/http"
)

type IndexController struct{}

func NewIndexController() IndexController {
	return IndexController{}
}

func (lc IndexController) Show(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}

	tmpl.Execute(w, nil)
}
