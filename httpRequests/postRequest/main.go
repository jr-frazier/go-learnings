package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	url := "https://httpbin.org/post"

	jon := &Person{
		Name: "Jon",
		Age:  30,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(jon); err != nil {
		log.Fatalf("error encoding json: %v", err)
	}

	// Make a POST request to the URL
	// and print the response to the console.
	resp, err := http.Post(url, "application/json", &buf)
	if err != nil {
		log.Fatalf("error: can't call %s", url)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

}
