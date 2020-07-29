package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"news-service/config"
	"news-service/fetcher"
	"news-service/handler"
	"news-service/models"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var (
	apiKeyPath = os.Getenv("NEWS_API_KEY")
	dbUser     = os.Getenv("DB_USER")
	dbPass     = os.Getenv("DB_PASSWORD")
	dbPort     = os.Getenv("DB_PORT")
	dbName     = os.Getenv("DB_NAME")
	dbHost     = os.Getenv("DB_HOST")
	devMode    = os.Getenv("DEV_MODE")
	sptLangs   = os.Getenv("SUPPORTED_LANGUAGES")

	apiKey string
)

func init() {
	if apiKeyPath == "" {
		log.Fatal("No environment variable named NEWS_API_KEY present")
	} else {
		content, err := ioutil.ReadFile(apiKeyPath)
		if err != nil {
			log.Fatal(err)
		}

		apiKey = string(content)
	}

	if dbUser == "" {
		log.Fatal("No environment variable named DB_USER present")
	}

	if dbPass == "" {
		log.Fatal("No environment variable named DB_PASSWORD present")
	}

	if dbPort == "" {
		log.Fatal("No environment variable named DB_PORT present")
	}

	if dbName == "" {
		log.Fatal("No environment variable named DB_NAME present")

	}

	if dbHost == "" {
		log.Fatal("No environment variable named DB_HOST present")
	}

	if sptLangs == "" {
		config.SupportedLangs = []string{"en"}
	}

	config.SupportedLangs = strings.Split(sptLangs, ",")
}

var (
	env *config.Env
)

func main() {
	connStr := fmt.Sprintf("port=%s user=%s password=%s dbname=%s host=%s sslmode=disable", dbPort, dbUser, dbPass, dbName, dbHost)
	db, err := models.New(connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	nf, err := fetcher.New(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	env = &config.Env{DB: db, NewsClient: nf}

	go env.NewsClient.FetchAndSave(context.Background(), env, time.Minute*15)

	r := mux.NewRouter()

	api := r.PathPrefix("/api/news").Subrouter()
	api.HandleFunc("/top-headlines", handler.HandleTopHeadlines(env)).Methods("GET")

	fmt.Println("The news server is ready")
	log.Fatal(http.ListenAndServe(":8083", r))
}
