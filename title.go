package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"regexp"

	"github.com/jangler/minipkg/tool"
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
	cmd := &tool.Command{
		Name:    "title",
		Summary: "print the title of a video or playlist at a URL",
		Usage:   "<url>",
		Description: `
Print the title of the video or playlist at <url>.
`,
		Function: title,
		MinArgs:  1,
		MaxArgs:  1,
	}

	cmd.FlagSet = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	cmd.FlagSet.Usage = tool.UsageFunc(cmd)

	tool.Commands[cmd.Name] = cmd
}
