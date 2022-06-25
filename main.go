package main

import (
	"encoding/base32"
	"fmt"
	"log"
	"net/http"

	"github.com/ksrnnb/otp/router"
)

func main() {
	secret := []byte{'h', 'e', 'l', 'l', 'o'}

	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	encSecret := encoder.EncodeToString(secret)
	fmt.Println("base32 secret:", encSecret)

	r := router.NewRouter()
	r.RegisterRoutes()

	fmt.Println("connect to localhost:8080")
	srv := &http.Server{Handler: r.Router(), Addr: "0.0.0.0:8080"}
	log.Fatal(srv.ListenAndServe())
}
