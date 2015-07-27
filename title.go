package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"regexp"
)

var videoRegexp = regexp.MustCompile(
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

type playlistsSnippet struct {
	ChannelID, Title, Description, ChannelTitle, PublishedAt string
}

type playlistsItem struct {
	Snippet playlistsSnippet
}

type playlistsResponse struct {
	Items []playlistsItem
}

func getPlaylistTitle(id string) string {
	addr := fmt.Sprintf("%s/playlists?key=%s&id=%s&part=snippet", apiRoot,
		apiKey, url.QueryEscape(id))

	var v playlistsResponse
	decodeResponse(addr, &v)

	if len(v.Items) != 1 {
		log.Fatal("bad playlists query response")
	}
	return v.Items[0].Snippet.Title
}

func getVideoTitle(id string) string {
	addr := fmt.Sprintf("%s/videos?key=%s&id=%s&part=snippet", apiRoot, apiKey,
		url.QueryEscape(id))

	var v videosResponse
	decodeResponse(addr, &v)

	if len(v.Items) != 1 {
		log.Fatal("bad videos query response")
	}
	return v.Items[0].Snippet.Title
}

func title(args []string) {
	if m := videoRegexp.FindStringSubmatch(args[0]); m != nil {
		fmt.Println(getVideoTitle(m[3]))
	} else if m := playlistRegexp.FindStringSubmatch(args[0]); m != nil {
		fmt.Println(getPlaylistTitle(m[2]))
	} else {
		log.Fatal("invalid URL")
	}
}

func init() {
	cmd := &command{
		name:    "title",
		summary: "print the title of a video or playlist at a URL",
		usage:   "<url>",
		description: `
Print the title of the video or playlist at <url>.
`,
		function: title,
		minArgs:  1,
		maxArgs:  1,
	}

	cmd.flagSet = flag.NewFlagSet(cmd.name, flag.ExitOnError)
	cmd.flagSet.Usage = usageFunc(cmd)

	commands[cmd.name] = cmd
}
