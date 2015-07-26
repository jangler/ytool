package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
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

func title(args []string) {
	match := idRegexp.FindStringSubmatch(args[0])
	if match == nil {
		log.Fatal("invalid video URL")
	}
	id := match[3]
	addr := fmt.Sprintf("%s/videos?key=%s&id=%s&part=snippet", apiRoot, apiKey,
		url.QueryEscape(id))

	var v videosResponse
	decodeResponse(addr, &v)

	for _, item := range v.Items {
		fmt.Println(item.Snippet.Title)
	}
}

func init() {
	cmd := &command{
		name:    "title",
		summary: "print the title of a video at a URL",
		usage:   "<url>",
		description: `
Print the title of the video at <url>.
`,
		function: title,
		minArgs:  1,
		maxArgs:  1,
	}

	cmd.flagSet = flag.NewFlagSet(cmd.name, flag.ExitOnError)
	cmd.flagSet.Usage = usageFunc(cmd)

	commands[cmd.name] = cmd
}
