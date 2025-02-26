package main

import (
	"io"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	expected := "Hello, World!\n"
	if string(body) != expected {
		t.Errorf("Expected %s but got %s", expected, string(body))
	}
}

func TestIndiaHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/india", nil)
	w := httptest.NewRecorder()

	indiaHandler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	expected := "Hello, India!\n"
	if string(body) != expected {
		t.Errorf("Expected %s but got %s", expected, string(body))
	}
}
