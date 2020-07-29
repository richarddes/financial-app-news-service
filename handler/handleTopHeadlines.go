// Package handler implements all http handlers.
package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"news-service/config"
	"news-service/internal"
)

// HandleTopHeadlines returns the top headlines saved in the environments datastore.
func HandleTopHeadlines(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		if !internal.IsSupportedLang(lang) {
			lang = "en"
		}

		arts, err := env.DB.TopHeadlines(r.Context(), lang)
		if err != nil {
			http.Error(w, "An unexpected error occured. Please try again later.", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(arts)
	}
}
