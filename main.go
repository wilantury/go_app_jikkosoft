package main

import (
	"go_app_jikkosoft/app/bundles/sort_array"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	arraySortHandle := sort_array.SortArrayHandle{}

	router := mux.NewRouter()

	router.HandleFunc("/api/sort", arraySortHandle.SortArrayView).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
