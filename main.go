package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jangler/minipkg/tool"
)

const (
	apiKey  = "AIzaSyDR5xOXOViVLMUyWWJM1iQefTaRiKkJfqs"
	apiRoot = "https://www.googleapis.com/youtube/v3"
)

func decodeResponse(addr string, v interface{}) {
	resp, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		log.Fatal(err)
	}
}

func main() {
	tool.Description = `
A command-line interface to the YouTube data API.
`
	tool.Main()
}
