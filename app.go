package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/gorilla/handlers"
	"github.com/ahaha0807/ishikari-2017-gorilla/handler"
)

type Config struct {
	Addr string
	Port string
}

func main() {
	config := Config{
		Addr: "127.0.0.1",
		Port: ":9999"}

	router := mux.NewRouter()

	// routing
	router.HandleFunc("/", handler.HomeHandler).Methods("GET")
	router.HandleFunc("/api/price", handler.ZaifHandler).Methods("GET")

	// static files settings
	router.PathPrefix("/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	http.Handle("/", router)

	log.Println("Server is Running on " + config.Addr + config.Port)
	log.Fatal(http.ListenAndServe(config.Port, handlers.CORS()(router)))
}
