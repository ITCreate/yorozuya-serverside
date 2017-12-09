package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"time"
)

func main() {
	router := mux.NewRouter()

	// routing
	router.HandleFunc("/", HomeHandler).Methods("GET")

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

// ぜひともハンドルファンクションは別ファイルで作りましょう
func HomeHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./public/index.html")
}
