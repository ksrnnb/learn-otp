package main

import (
	"encoding/base32"
	"fmt"
	"log"
	"net/http"

	"github.com/ksrnnb/otp/router"
	"github.com/ksrnnb/otp/totp"
)

func main() {
	secret := []byte{'h', 'e', 'l', 'l', 'o'}

	fmt.Println(totp.New(secret, 6))

	encoder := base32.StdEncoding.WithPadding(base32.NoPadding)
	encSecret := encoder.EncodeToString(secret)
	fmt.Println("base32 secret:", encSecret)

	r := router.NewRouter()
	r.RegisterRoutes()

	fmt.Println("connect to localhost:8080")
	srv := &http.Server{Handler: r.Router(), Addr: "0.0.0.0:8080"}
	log.Fatal(srv.ListenAndServe())
}
