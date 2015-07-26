package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

var idRegexp = regexp.MustCompile(
	`^https?://(www\.)?youtu(be\.com/watch\?v=|\.be/)(.+)$`)

type videosSnippet struct {
	ChannelID, Title, Description, CategoryID string

	Tags []string
}

type videosItem struct {
	Snippet videosSnippet
}

type videosResponse struct {
	Items []videosItem
}

func title(arg string) {
	log.SetPrefix(fmt.Sprintf("%s: title: ", os.Args[0]))
	match := idRegexp.FindStringSubmatch(arg)
	if match == nil {
		log.Fatal("invalid video URL")
	}
	id := match[3]
	addr := fmt.Sprintf("%s/videos?key=%s&id=%s&part=snippet", apiRoot, apiKey,
		url.QueryEscape(id))

	resp, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var v videosResponse
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		log.Fatal(err)
	}

	for _, item := range v.Items {
		fmt.Println(item.Snippet.Title)
	}
}
