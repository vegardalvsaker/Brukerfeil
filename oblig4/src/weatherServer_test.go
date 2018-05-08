package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	//"fmt"
	"fmt"
	"bytes"
)

func TestIndex (t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/", nil)
	if err != nil {
		t.Error("could not create request %v", err)
	}
	rec := httptest.NewRecorder()

	index(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Error("Expected Status OK: 200; got %v", err)
	}
	html, err := ioutil.ReadAll(res.Body)
	index, err := ioutil.ReadFile("index.html")

	if string(bytes.Trim(html, " ")) != string(bytes.Trim(index, " ")) {
		t.Error("The request gets the wrong index.html file")

	}
	fmt.Printf("%q", html)
	fmt.Printf("%q", index)

}
