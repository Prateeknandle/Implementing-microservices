package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//handles request
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("hello world")
		d, err := ioutil.ReadAll(r.Body) //read data from body
		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", d) // write data to user, uses http.ResponseWriter
	})

	http.HandleFunc("/GoodBye", func(http.ResponseWriter, *http.Request) {
		log.Printf("GoodBye World")
	})

	http.ListenAndServe(":9098", nil) // handler = defaultServeMux, as handlers is not mentioned
}
