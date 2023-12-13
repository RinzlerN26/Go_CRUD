package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func TestWelcome(t *testing.T) {

	resp, err := http.Get("http://localhost:8080/welcome")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		t.Log("Welcome Successful")
	}
}

func TestGET(t *testing.T) {

	resp, err := http.Get("http://localhost:8080/get_cars")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		t.Log("GET Successful")
	}
}

func TestPOST(t *testing.T) {
	type PostBody struct {
		Type string `json:"type"`
	}

	body := PostBody{
		Type: "Post Request",
	}
	bodyBytes, err1 := json.Marshal(&body)

	if err1 != nil {
		log.Fatal(err1)
	}
	reader := bytes.NewReader(bodyBytes)

	resp, err := http.Post("http://localhost:8080/cars/jaguar/In_garage", "application/json", reader)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		t.Log("POST Successful")
	}
	defer resp.Body.Close()
}

func TestPUT(t *testing.T) {
	type PutBody struct {
		Type string `json:"type"`
	}

	body := PutBody{
		Type: "Put Request",
	}
	bodyBytes, err1 := json.Marshal(&body)

	if err1 != nil {
		log.Fatal(err1)
	}
	reader := bytes.NewReader(bodyBytes)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/cars/3/RepairingDone", reader)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Fatal(err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		t.Log("PUT Successful")
	}
	defer resp.Body.Close()
}

func TestDELETE(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://localhost:8080/cars/delete", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		t.Log("DELETE Successful")
	}
	defer resp.Body.Close()
}
