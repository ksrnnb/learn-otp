package router

import (
	"github.com/gorilla/mux"
	"github.com/ksrnnb/otp/controller"
)

type Router struct {
	mr *mux.Router
}

func NewRouter() *Router {
	return &Router{mr: mux.NewRouter()}
}

var ic = controller.NewIndexController()
var lc = controller.NewLoginController()

const (
	Get  = "GET"
	Post = "POST"
)

func (r Router) RegisterRoutes() {
	r.mr.HandleFunc("/", ic.Show).Methods(Get)
	r.mr.HandleFunc("/login", lc.Show).Methods(Get)
	r.mr.HandleFunc("/login", lc.Login).Methods(Post)
	r.mr.HandleFunc("/login/otp", lc.ShowOTPLogin).Methods(Get)
	r.mr.HandleFunc("/login/otp", lc.OTPLogin).Methods(Post)
}

func (r Router) Router() *mux.Router {
	return r.mr
}
