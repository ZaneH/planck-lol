package main

import (
	http "planck-lol/internal/http"
)

func main() {
	server := http.SetupServer()
	server.ListenAndServe()
}
