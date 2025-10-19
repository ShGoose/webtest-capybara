package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	http.HandleFunc("/", helloHandler)
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200 ОК, но получили %v", res.StatusCode)
	}

	body := w.Body.String()
	expected := "Hello, world!\n"
	if !strings.Contains(body, expected) {
		t.Errorf("Ожидалось тело ответа %q, но получили %q", expected, body)
	}
}

func TestSecondHandler(t *testing.T) {
	http.HandleFunc("/second", secondHandler)
	req := httptest.NewRequest("GET", "/second", nil)
	w := httptest.NewRecorder()

	http.DefaultServeMux.ServeHTTP(w, req)
	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Ожидался статус 200 ОК, но получили %v", res.StatusCode)
	}

	body := w.Body.String()
	expected := "This is the second handler!\n"
	if !strings.Contains(body, expected) {
		t.Errorf("Ожидалось тело ответа %q, но получили %q", expected, body)
	}
}
