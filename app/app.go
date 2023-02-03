package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()

	//Routes
	router.HandleFunc("/api/time", GetTime).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
