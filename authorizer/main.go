package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	rootPath = "/authorize"
	Port     = 8090
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(rootPath, HTTPHandler).Methods(http.MethodGet)

	err := http.ListenAndServe(fmt.Sprintf(":%d", Port), router)
	fmt.Printf("%v", err)
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("This is authorization")

	okStatus := r.Header.Get("returnOK")
	if okStatus == "200" {
		log.Println(w, "200!")
		w.Header().Set("notfound", "false")

		w.WriteHeader(http.StatusOK)
	} else if okStatus == "401" {
		log.Println(w, "401!")
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		log.Println(w, "404")
		w.Header().Set("notfound", "true")
		w.Header().Set("Expect", "1")

		w.WriteHeader(http.StatusNotFound)

	}

	return
}
