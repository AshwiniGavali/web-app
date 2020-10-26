package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Welcome to the world of CI/CD with Jenkins CI</h1>")
	})

	http.ListenAndServe(":9090", r)
}
