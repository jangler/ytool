package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	apiKey  = "AIzaSyDR5xOXOViVLMUyWWJM1iQefTaRiKkJfqs"
	apiRoot = "https://www.googleapis.com/youtube/v3"
)

const description = `
A command-line interface to the YouTube data API. If not enough
command-line arguments are specified for a command, remaining arguments
are read from standard input.
`

const commands = `
  search <query>...  print the top URL matching <query>
  title <url>        print the title of the video at <url>
`

func parseFlags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <cmd> [<arg>]...\n", os.Args[0])
		fmt.Fprintln(os.Stderr, description)
		fmt.Fprint(os.Stderr, "Commands:")
		fmt.Fprint(os.Stderr, commands)
		os.Exit(2)
	}
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
	}
}

func readStdin(args []string, max int) []string {
	r := bufio.NewReader(os.Stdin)
	for max < 0 || len(args) < max {
		line, err := r.ReadString('\n')
		if len(line) > 0 {
			args = append(args, string(line[:len(line)-1]))
		}
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return args
}

func getArgs(cmd string, min, max int) []string {
	args := flag.Args()[1:]
	if max >= 0 && len(args) > max {
		fmt.Fprintf(os.Stderr, "%s: %s: too many arguments\n", os.Args[0], cmd)
		os.Exit(2)
	}
	if len(args) < min {
		args = readStdin(args, max)
	}
	if len(args) < min {
		fmt.Fprintf(os.Stderr, "%s: %s: missing query argument\n", os.Args[0],
			cmd)
		os.Exit(2)
	}
	return args
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", os.Args[0]))
	parseFlags()
	switch flag.Arg(0) {
	case "search":
		search(getArgs("search", 1, -1)...)
	case "title":
		title(getArgs("title", 1, 1)[0])
	default:
		fmt.Fprintf(os.Stderr, "%s: no such command: %s\n", os.Args[0],
			flag.Arg(0))
		os.Exit(2)
	}
}
