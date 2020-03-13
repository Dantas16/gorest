package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"Info": {"id": 01, "service": "Infos from mongo"}}`))
}

func main() {
	fmt.Println("Starting Restful services...")
	fmt.Println("Using port: 8080")
	// http.HandleFunc("/hello", home)
	r := mux.NewRouter()
	r.HandleFunc("/hello", home).Methods(http.MethodGet)
	r.HandleFunc("/info", info).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
	// err := http.ListenAndServe(":8080", nil)
	// log.Print(err)
	// errorHandler(err)
}

/*func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}*/
