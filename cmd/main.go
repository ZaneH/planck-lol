package main

import "planck-lol/internal/shortener"

func main() {
	linkService := &shortener.LinkService{}
	server := shortener.SetupServer(linkService)
	server.ListenAndServe()
}
