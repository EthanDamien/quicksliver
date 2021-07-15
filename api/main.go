package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	bindAddress string = ":8000"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/api/health-check", HealthCheck()).Methods(http.MethodGet)

	log.Printf("API started")

	srv := &http.Server{
		Addr:    bindAddress,
		Handler: router,
	}

	log.Println("hosted on ", bindAddress)
	log.Fatal(srv.ListenAndServe())
}

func HealthCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	})
}
