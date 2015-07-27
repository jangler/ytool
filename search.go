package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/jangler/minipkg/tool"
)

var flagSearchN uint = 1
var flagSearchType = "video,channel,playlist"

type searchID struct {
	Kind, VideoID, PlaylistID string
}

type searchSnippet struct {
	ChannelTitle string
}

type searchItem struct {
	ID      searchID
	Snippet searchSnippet
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
		"%s/search?key=%s&part=snippet&maxResults=%d&q=%s&type=%s",
		apiRoot, apiKey, flagSearchN, url.QueryEscape(query),
		url.QueryEscape(flagSearchType))

	var v searchResponse
	decodeResponse(addr, &v)

	for _, item := range v.Items {
		switch item.ID.Kind {
		case "youtube#channel":
			fmt.Printf("https://www.youtube.com/user/%s\n",
				item.Snippet.ChannelTitle)
		case "youtube#playlist":
			fmt.Printf("https://www.youtube.com/playlist?list=%s\n",
				item.ID.PlaylistID)
		case "youtube#video":
			fmt.Printf("https://www.youtube.com/watch?v=%s\n", item.ID.VideoID)
		default:
			fmt.Printf("Unknown item type: %s\n", item.ID.Kind)
		}
	}
}

func init() {
	cmd := &tool.Command{
		Name:    "search",
		Summary: "print the URLs of resources matching a query",
		Usage:   "[<option>]... <query>...",
		Description: `
Search YouTube for <query> (joined by spaces if multiple arguments are
given) and print the URLs of the top matches in descending order by
relevance.
`,
		Function: search,
		MinArgs:  1,
		MaxArgs:  -1,
		HasOpts:  true,
	}

	cmd.FlagSet = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	cmd.FlagSet.Usage = tool.UsageFunc(cmd)
	cmd.FlagSet.UintVar(&flagSearchN, "n", flagSearchN,
		"maximum number of results, in the range [0, 50]")
	cmd.FlagSet.StringVar(&flagSearchType, "type", flagSearchType,
		"restrict search to given resource types")

	tool.Commands[cmd.Name] = cmd
}
