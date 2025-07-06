package internal

import (
	"net/http"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("ok"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetCodeHandle(w http.ResponseWriter, r *http.Request) {
}

func CreateCodeHandle(w http.ResponseWriter, r *http.Request) {
}
