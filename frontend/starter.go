package main

import (
	"fmt"
	"go-crawler/config"
	"go-crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(
		http.Dir("frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(
		"frontend/view/template.html"))
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.FrontendPort), nil)
	if err != nil {
		panic(err)
	}
}
