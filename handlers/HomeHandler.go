package handlers

import "net/http"

func HomeHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./public/index.html")
}
