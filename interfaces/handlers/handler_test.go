package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/bondhan/godddnews/internal/manager"
)

type newsData struct {
	ID      uint     `json:"id"`
	Title   string   `json:"title"`
	Slug    string   `json:"slug"`
	Content string   `json:"content"`
	Status  string   `json:"status"`
	Version uint     `json:"version"`
	Topics  []string `json:"topic_slugs"`
	Tags    []string `json:"tag_slugs"`
}

func TestMain(m *testing.M) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "54322")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_NAME", "godddnews_db")
	os.Setenv("DB_PASSWORD", "root")
	os.Setenv("DB_SSLMODE", "disable")
	os.Setenv("PRODUCTION_ENV", "false")

	manager.GetContainer()

	os.Exit(m.Run())
}

// TestHealthCheckPing is checking if service is alive
func TestHealthCheckPing(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ping)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `pong`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// TestGetAllNewsHandler is testing if duplicate news is handled
func TestGetAllNewsHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/api/v1/news", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllNews)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestGetAllTopicHandler is testing if duplicate news is handled
func TestGetAllTopicHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/api/v1/topic", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllTopic)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// TestGetAllTagHandler is testing if duplicate news is handled
func TestGetAllTagHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/api/v1/tag", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllTag)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestGetAllNewsByTopicSlugHandler is testing if duplicate news is handled
func TestGetAllNewsByTopicSlugHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/api/v1/news/topic/politics", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getNewsByTopicSlug)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestGetAllNewsByTagSlugHandler is testing if duplicate news is handled
func TestGetAllNewsByTagSlugHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "localhost:8080/api/v1/news/tag/other-tag", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getNewsByTagSlug)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestCreateNewsHandler is creating a new news
func TestCreateNewsHandler(t *testing.T) {

	nData := newsData{
		ID:      uint(rand.Intn(100)),
		Title:   "news test",
		Slug:    "news-test",
		Content: "Title",
		Status:  "draft",
		Version: 1,
		Topics:  []string{"politics", "national"},
		Tags:    []string{"other-tag", "national-tag"},
	}

	n, err := json.Marshal(nData)
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println(string(n))

	req, err := http.NewRequest("POST", "localhost:8080/api/v1/news", strings.NewReader(string(n)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createNews)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status":"Successful","code":201}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
