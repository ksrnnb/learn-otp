package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ksrnnb/otp/router"
)

func main() {
	r := router.NewRouter()
	r.RegisterRoutes()

	fmt.Println("connect to localhost:8080")
	srv := &http.Server{Handler: r.Router(), Addr: "0.0.0.0:8080"}
	log.Fatal(srv.ListenAndServe())
}
