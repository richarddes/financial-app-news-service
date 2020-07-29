package handler_test

import (
	"net/http"
	"net/http/httptest"
	"news-service/config"
	"news-service/handler"
	"strings"
	"testing"
)

func TestHandleTopHeadlines(t *testing.T) {
	mockEnv := config.NewMockEnv()

	langs := []struct {
		lang             string
		expectedRespCode int
	}{
		{"en", http.StatusOK},
		{"de", http.StatusOK},
		// "us" isn't a 2 letter ISO-639-1 code but it should still work since it should default to "en" when
		// it's not a supported language
		{"us", http.StatusOK},
	}

	for _, l := range langs {
		req, err := http.NewRequest("GET", "/api/news/top-headlines", strings.NewReader(l.lang))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		h := http.HandlerFunc(handler.HandleTopHeadlines(mockEnv))

		h.ServeHTTP(rr, req)

		if rr.Code != l.expectedRespCode {
			t.Errorf("Expected status code %d but got %d instead", l.expectedRespCode, rr.Code)
		}
	}
}
