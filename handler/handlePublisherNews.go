package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"news-service/config"
)

func HandlePublisherNews(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var publisherIDs []string

		dc := json.NewDecoder(r.Body)
		err := dc.Decode(&publisherIDs)
		if err != nil {
			http.Error(w, "Invalid request syntax", http.StatusBadRequest)
			return
		}

		arts, err := env.DB.PublisherNews(r.Context(), publisherIDs)
		if err != nil {
			http.Error(w, "An error has occured while retrieving the news from the publishers", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(arts)
	}
}
