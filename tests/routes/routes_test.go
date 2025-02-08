package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EllieZora/TerminalRPG/internal/routes"
)

func TestIndexRoute(t *testing.T) {
	router := routes.NewRouter()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	expectedContentType := "text/plain"
	if contentType := recorder.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("content type of header does not match: got %v want %v", contentType, expectedContentType)
	}

	expectedStatus := http.StatusOK
	if status := recorder.Code; status != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v", status, expectedStatus)
	}

	expectedBody := "Welcome to TerminalRPG!"
	if body := recorder.Body.String(); body != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedBody)
	}
}
