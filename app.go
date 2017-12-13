package main

import (
	"net/http"
	"log"
	"time"
	"github.com/rs/cors"
	"github.com/ahaha0807/ishikari-2017-gorilla/handlers"
)

func main() {
	router := http.NewServeMux()

	// routing
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/api/price", handlers.ZaifHandler)

	handler := cors.Default().Handler(router)

	// server setting
	srv := &http.Server{
		Handler: handler,
		Addr:    "127.0.0.1:9999",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server is Running on " + string(srv.Addr))
	log.Fatal(srv.ListenAndServe())
}
