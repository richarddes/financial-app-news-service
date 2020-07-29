// +build integration

package models_test

import (
	"context"
	"fmt"
	"log"
	"news-service/config"
	"news-service/models"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASSWORD")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")
	dbHost = os.Getenv("DB_HOST")
)

func init() {
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
		dbHost = "localhost"
	}

	// test values
	config.SupportedLangs = []string{"en", "de"}
}

func TestDefaulImpl(t *testing.T) {
	connStr := fmt.Sprintf("port=%s user=%s password=%s dbname=%s host=%s sslmode=disable", dbPort, dbUser, dbPass, dbName, dbHost)
	db, err := models.New(connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ModelsSuite(t, db)
}

func ModelsSuite(t *testing.T, impl config.Datastore) {
	ctx := context.Background()

	arts := []config.Article{
		config.Article{Author: "John Doe", Content: "Some Content", Description: "A Description", PublishedAt: "1.1.1970", Title: "Title 1", Source: "johndoe.com", URL: "https://johndoe.com/title-1", URLToImage: "https://johndoe.com/title-1/img", Lang: "en"},
		config.Article{Author: "John Doe", Content: "Some Content", Description: "A Description", PublishedAt: "1.1.1970", Title: "Title 2", Source: "johndoe.com", URL: "https://johndoe.com/title-2", URLToImage: "https://johndoe.com/title-2/img", Lang: "en"},
		config.Article{Author: "John Doe", Content: "Some Content", Description: "A Description", PublishedAt: "1.1.1970", Title: "Title 3", Source: "johndoe.com", URL: "https://johndoe.com/title-3", URLToImage: "https://johndoe.com/title-3/img", Lang: "de"},
		config.Article{Author: "John Doe", Content: "Some Content", Description: "A Description", PublishedAt: "1.1.1970", Title: "Title 4", Source: "johndoe.com", URL: "https://johndoe.com/title-4", URLToImage: "https://johndoe.com/title-4/img", Lang: "de"},
	}

	t.Run("test article insertion", func(t *testing.T) {
		for _, a := range arts {
			err := impl.InsertArticle(ctx, a)
			if err != nil {
				t.Fatalf("Unexpected error: %v when article=%v", err, a)
			}
		}
	})

	t.Run("test article reading", func(t *testing.T) {
		cases := []struct {
			lang  string
			valid bool
		}{
			{"en", true},
			{"de", true},
			{"fr", false},
			{"us", false},
		}

		for _, c := range cases {
			as, err := impl.TopHeadlines(ctx, c.lang)

			if c.valid {
				if err != nil {
					t.Errorf("Unexpected error: %v when lang=%s", err, c.lang)
				}

				// We expect to get back 2 articles when the language equals "en" or "de".
				if len(as) != 2 {
					t.Errorf("Originially 2 articles have been inserted but got %d articles back", len(as))
				}
			} else {
				if err == nil {
					t.Errorf("Expected an error but go none when lang=%s", c.lang)
				}
			}
		}
	})

	t.Run("test EntriesExist() with entries", func(t *testing.T) {
		exist, err := impl.EntriesExist(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if !exist {
			t.Fatal("EntriesExist() returned false even though entries exist")
		}
	})

	t.Run("test database clearing", func(t *testing.T) {
		err := impl.Clear(ctx)
		if err != nil {
			t.Fatal(err)
		}

		exist, err := impl.EntriesExist(ctx)
		if err != nil {
			t.Fatal(err)
		}

		if exist {
			t.Errorf("The database wasn't cleared")
		}
	})
}
