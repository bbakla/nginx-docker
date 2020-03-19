package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	rootPath        = "/"
	nginxPath       = "/nginx"
	TenantCheckPort = 8088
)

//with mozilla
/*
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(rootPath, HTTPHandler)

	log.Fatal(http.ListenAndServe(strconv.Itoa(TenantCheckPort), router))
}
*/


func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc(nginxPath, HelloHandler).Methods(http.MethodGet)
	router.HandleFunc(rootPath, HTTPHandler).Methods(http.MethodGet)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/all", getAllEventsWithTitle).Methods("GET")
	router.HandleFunc("/events/All", getAllEventsWithTitle2).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	err := http.ListenAndServe(":8099", router)
	fmt.Printf("%v", err)
}

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("This is root folder. dont expect nothing more than welcome")
	fmt.Fprintf(w, "Welcome home!")
}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	log.Println("lets seee")

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	requestInfo := make(map[string]string)

	url := fmt.Sprintf("%v %v %v", request.Method, request.URL, request.Proto)
	requestInfo["url"] = url

	for k, v := range request.Header {
		requestInfo[k] = fmt.Sprintf("%q", v)
	}

	requestInfo["host"] = fmt.Sprintf("%q", request.Host)
	requestInfo["remoteAddress"] = fmt.Sprintf("%q", request.RemoteAddr)

	mapAsJson, _ := json.Marshal(requestInfo)

	fmt.Fprintf(writer, "%s", mapAsJson)
}


