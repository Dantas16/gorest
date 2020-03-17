package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

type User struct {
	ID      int    `json:"ID"`
	name string	   `json:"name"`
}

func info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017/"))
	if err != nil {
		fmt.Println("erro NewClient")
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println("Erro client Connect")
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Println("Erro ping")
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("restapi").Collection("users")

	golang := User{1, "Golang"}
	inserted, err := collection.InsertOne(context.TODO(), golang)
	if err != nil {
		fmt.Println("Erro ao inserir")
		log.Fatal(err)
	}
	fmt.Println("Inserted document: ", inserted.InsertedID)

	filter := bson.M{"id": 1}
	var result User

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Erro ao buscar")
		log.Fatal(err)
	}

	fmt.Printf("Result: ", result)

	// w.Write([]byte(`{"Info_mongo": {"id": 01, "service": "Infos from mongo"}}`))
	json.NewEncoder(w).Encode(result)
}

func main() {
	fmt.Println("Starting Restful services...")
	fmt.Println("Using port: 8080")

	r := mux.NewRouter()
	r.HandleFunc("/hello", home).Methods(http.MethodGet)
	r.HandleFunc("/info", info).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
