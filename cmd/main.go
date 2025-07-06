package main

import (
	"planck-lol/internal/controller"
	http "planck-lol/internal/http"
)

func main() {
	controller := &controller.LinkController{}
	server := http.SetupServer(controller)
	server.ListenAndServe()
}
