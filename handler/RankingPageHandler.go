package handler

import "net/http"

func RankingPageHandler(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./public/ranking.html")
}
