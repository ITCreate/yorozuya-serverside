package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"time"
	"github.com/ahaha0807/yorozuya-serverside/handler"
)

func main() {
	router := mux.NewRouter()

	// routing
	router.HandleFunc("/", handler.HomeHandler).Methods("GET")
	router.HandleFunc("/api/price", handler.ZaifHandler).Methods("GET")

	// static files settings
	router.PathPrefix("/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))

	http.Handle("/", router)

	// server setting
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9999",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
