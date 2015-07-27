package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"regexp"

	"github.com/jangler/minipkg/tool"
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
	cmd := &tool.Command{
		Name:    "playlist",
		Summary: "print the URLs of videos in a playlist",
		Usage:   "<url>",
		Description: `
Print the URLs of the videos in the playlist at <url>, up to a maximum
of 50 videos.
`,
		Function: playlist,
		MinArgs:  1,
		MaxArgs:  1,
	}

	cmd.FlagSet = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	cmd.FlagSet.Usage = tool.UsageFunc(cmd)

	tool.Commands[cmd.Name] = cmd
}
