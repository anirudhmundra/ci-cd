package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	expected := "Hello, World!\n"
	if string(body) != expected {
		t.Errorf("Expected %s but got %s", expected, string(body))
	}
}
