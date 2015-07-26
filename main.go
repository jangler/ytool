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
A command-line interface to the YouTube data API. If no command-line
arguments are specified for a command that requires arguments, arguments
are read from standard input.
`

const commands = `
  search <query>...  print the top URL matching <query>
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

func readStdin() []string {
	r := bufio.NewReader(os.Stdin)
	args := []string{}
	for {
		line, err := r.ReadString('\n')
		if len(line) > 0 {
			args = append(args, string(line))
		}
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return args
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", os.Args[0]))
	parseFlags()
	args := flag.Args()[1:]
	switch flag.Arg(0) {
	case "search":
		if len(args) < 1 {
			args = readStdin()
		}
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "%s: search: missing query argument\n",
				os.Args[0])
			os.Exit(2)
		}
		search(args...)
	default:
		fmt.Fprintf(os.Stderr, "%s: no such command: %s\n", os.Args[0],
			flag.Arg(0))
		os.Exit(2)
	}
}
