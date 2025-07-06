package internal

import (
	"net/http"
	"planck-lol/internal/controller"
)

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("ok"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func GetCodeHandle(controller *controller.LinkController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")
		link, err := controller.GetLink(code)
		if err != nil {
			http.Error(w, "Link Not Found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link, http.StatusTemporaryRedirect)
		return
	}
}

func CreateCodeHandle(controller *controller.LinkController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		defer r.Body.Close()

		// TODO: Read body and create link<->code with controller
		return
	}
}
