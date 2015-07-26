package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type searchID struct {
	Kind, ChannelID, VideoID, PlaylistID string
}

type searchItem struct {
	ID searchID
}

type searchResponse struct {
	Items []searchItem
}

func search(arg ...string) {
	log.SetPrefix(fmt.Sprintf("%s: search: ", os.Args[0]))
	query := strings.Join(arg, " ")
	addr := fmt.Sprintf(
		"%s/search?key=%s&part=id&maxResults=1&q=%s&type=video",
		apiRoot, apiKey, url.QueryEscape(query))

	resp, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var v searchResponse
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		log.Fatal(err)
	}

	for _, item := range v.Items {
		fmt.Printf("https://youtube.com/watch?v=%s\n", item.ID.VideoID)
	}
}
