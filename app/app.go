package app

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute).Methods(http.MethodGet)
	router.HandleFunc("/tasks", getTasks).Methods(http.MethodGet)
	router.HandleFunc("/tasks", createTask).Methods(http.MethodPost)
	router.HandleFunc("/tasks/{id:[0-9]+}", getTask).Methods(http.MethodGet)
	router.HandleFunc("/tasks/{id:[0-9]+}", deleteTask).Methods(http.MethodDelete)
	router.HandleFunc("/tasks/{id:[0-9]+}", updateTask).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8000", router))
}