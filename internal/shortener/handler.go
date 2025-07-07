package shortener

import (
	"encoding/json"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("ok"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleCodeGet(s *LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")

		link, err := s.GetLink(code)
		if err != nil {
			http.Error(w, "Link Not Found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.LongUrl, http.StatusTemporaryRedirect)
		return
	}
}

func handleCodeCreate(s *LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()

		var data Link
		json.NewDecoder(r.Body).Decode(&data)

		s.CreateLink(&data.ShortCode, data.LongUrl)

		w.WriteHeader(http.StatusNoContent)
		return
	}
}
