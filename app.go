package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/ahaha0807/ishikari-2017-gorilla/handler"
	"github.com/gorilla/handlers"
)

func main() {
	router := mux.NewRouter()

	// routing
	router.HandleFunc("/", handler.HomeHandler).Methods("GET")
	router.HandleFunc("/api/price", handler.ZaifHandler).Methods("GET")

	// static files settings
	router.PathPrefix("/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	http.Handle("/", router)

	log.Println("Server is Running on " + string("127.0.0.1:9999"))
	log.Fatal(http.ListenAndServe(":9999", handlers.CORS()(router)))
}
