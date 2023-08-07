
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestTrue(t *testing.T) {
	if true != true {
		t.Errorf("Assertion failed: true should equal true")
	}
}

func TestRootURL(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Welcome to WebGo!"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Handler returned unexpected body: got %v, expected to contain %v", rr.Body.String(), expected)
	}
}