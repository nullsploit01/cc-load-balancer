package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	os.Setenv("SERVER_NAME", "test-server")

	router := NewRouter("test-server")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", rr.Code)
	}

	expectedBody := `{"message": "test-server says, sup?"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, rr.Body.String())
	}
}

func TestDefaultServerName(t *testing.T) {
	os.Unsetenv("SERVER_NAME")

	router := NewRouter("")

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", rr.Code)
	}

	expectedBody := `{"message": "default-server says, sup?"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, rr.Body.String())
	}
}
