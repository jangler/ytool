package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"regexp"
)

var playlistRegexp = regexp.MustCompile(
	`^https://(www\.)?youtube\.com/playlist\?list=(.+)$`)

type playlistContentDetails struct {
	VideoID string
}

type playlistItem struct {
	ContentDetails playlistContentDetails
}

type playlistResponse struct {
	Items []playlistItem
}

func playlist(args []string) {
	match := playlistRegexp.FindStringSubmatch(args[0])
	if match == nil {
		log.Fatal("invalid playlist URL")
	}
	id := match[2]
	addr := fmt.Sprintf("%s/playlistItems?key=%s&maxResults=50&playlistId=%s&"+
		"part=contentDetails", apiRoot, apiKey, url.QueryEscape(id))

	var v playlistResponse
	decodeResponse(addr, &v)

	for _, item := range v.Items {
		fmt.Printf("https://www.youtube.com/watch?v=%s\n",
			item.ContentDetails.VideoID)
	}
}

func init() {
	cmd := &command{
		name:    "playlist",
		summary: "print the URLs of videos in a playlist",
		usage:   "<url>",
		description: `
Print the URLs of the videos in the playlist at <url>, up to a maximum
of 50 videos.
`,
		function: playlist,
		minArgs:  1,
		maxArgs:  1,
	}

	cmd.flagSet = flag.NewFlagSet(cmd.name, flag.ExitOnError)
	cmd.flagSet.Usage = usageFunc(cmd)

	commands[cmd.name] = cmd
}
