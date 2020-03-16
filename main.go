package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

type infoType struct {
	ID      int    `json:"ID"`
	Service string `json:"alias"`
}

func info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	w.Write([]byte(`{"Info_mongo": {"id": 01, "service": "Infos from mongo"}}`))
}

func main() {
	fmt.Println("Starting Restful services...")
	fmt.Println("Using port: 8080")

	r := mux.NewRouter()
	r.HandleFunc("/hello", home).Methods(http.MethodGet)
	r.HandleFunc("/info", info).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
