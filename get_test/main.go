package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatalf("Failed to get URL: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	v, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Failed to marshal response body: %v", err)
	}

	fmt.Println(string(v))
}
