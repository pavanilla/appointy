package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pavanilla/middleware"
)

func TestCreatePosts(t *testing.T) {
	request, err := http.NewRequest("POST", "/posts", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(middleware.CreatePost)
	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := []byte(`<h1>post creation successful</h1>`)

	if rr.Body.ReadBytes() != expected {
		t.Errorf("Error while creating the posts  ")
	}
}
