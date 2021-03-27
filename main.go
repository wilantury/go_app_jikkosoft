package main

import (
	"go_app_jikkosoft/app/bundles/sort_array"
	"go_app_jikkosoft/app/bundles/users"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	arraySortHandle := sort_array.SortArrayHandle{}
	userHandle := users.Controller{}
	router := mux.NewRouter()

	router.HandleFunc("/api/sort", arraySortHandle.SortArrayView).Methods("POST")
	router.HandleFunc("/api/users", userHandle.GetUsers).Methods("GET")
	router.HandleFunc("/api/users", userHandle.PostUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
