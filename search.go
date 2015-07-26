package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"
)

var flagSearchN uint = 1

type searchID struct {
	Kind, ChannelID, VideoID, PlaylistID string
}

type searchItem struct {
	ID searchID
}

type searchResponse struct {
	Items []searchItem
}

func search(args []string) {
	if flagSearchN > 50 {
		log.Fatal("-n option must be in the range [0, 50]")
	}

	query := strings.Join(args, " ")
	addr := fmt.Sprintf(
		"%s/search?key=%s&part=id&maxResults=%d&q=%s&type=video",
		apiRoot, apiKey, flagSearchN, url.QueryEscape(query))

	var v searchResponse
	decodeResponse(addr, &v)

	for _, item := range v.Items {
		fmt.Printf("https://www.youtube.com/watch?v=%s\n", item.ID.VideoID)
	}
}

func init() {
	cmd := &command{
		name:    "search",
		summary: "print the URLs of videos matching a query",
		usage:   "[<option>]... <query>...",
		description: `
Search YouTube for <query> (joined by spaces if multiple arguments are
given) and print the URLs of the top matches in descending order by
relevance.
`,
		function: search,
		minArgs:  1,
		maxArgs:  -1,
		hasOpts:  true,
	}

	cmd.flagSet = flag.NewFlagSet(cmd.name, flag.ExitOnError)
	cmd.flagSet.Usage = usageFunc(cmd)
	cmd.flagSet.UintVar(&flagSearchN, "n", flagSearchN,
		"maximum number of results, in the range [0, 50]")

	commands[cmd.name] = cmd
}
