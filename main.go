package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Prateeknandle/Implementing-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	//handles request
	http.ListenAndServe(":9098", sm) // sm = serveMux
}
