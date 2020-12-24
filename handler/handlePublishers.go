package handler

import (
	"encoding/json"
	"net/http"
	"news-service/config"
	"news-service/internal"
)

func HandlePublishers(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		if !internal.IsSupportedLang(lang) {
			lang = "en"
		}

		publishers := config.Publishers[lang]

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(publishers)
	}
}
