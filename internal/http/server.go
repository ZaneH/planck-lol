package internal

import (
	"net/http"
)

func SetupServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandle)
	mux.HandleFunc("/create", CreateCodeHandle)
	mux.HandleFunc("/:code", GetCodeHandle)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return server
}
