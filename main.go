package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
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

For help regarding a specific command, see '%s <cmd> -h'.
`

type command struct {
	name        string
	summary     string
	usage       string
	description string
	function    func([]string)
	flagSet     *flag.FlagSet
	minArgs     int
	maxArgs     int
	hasOpts     bool
}

var commands = make(map[string]*command)

func usageFunc(cmd *command) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "Usage: %s %s %s\n",
			os.Args[0], cmd.name, cmd.usage)
		fmt.Fprint(os.Stderr, cmd.description)
		if cmd.hasOpts {
			fmt.Fprintln(os.Stderr, "\nOptions:")
			cmd.flagSet.PrintDefaults()
		}
	}
}

func decodeResponse(addr string, v interface{}) {
	resp, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		log.Fatal(err)
	}
}

func printCommands() {
	maxlen := 0
	for _, cmd := range commands {
		if len(cmd.name) > maxlen {
			maxlen = len(cmd.name)
		}
	}
	space := "        "

	for _, cmd := range commands {
		fmt.Fprintf(os.Stderr, "  %s%s  %s\n", cmd.name,
			space[:maxlen-len(cmd.name)], cmd.summary)
	}
}

func parseFlags() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <cmd> [<arg>]...\n", os.Args[0])
		fmt.Fprintf(os.Stderr, description, os.Args[0])
		fmt.Fprint(os.Stderr, "\nCommands:\n")
		printCommands()
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

func getArgs(cmd string, min, max int, args []string) []string {
	if max >= 0 && len(args) > max {
		fmt.Fprintf(os.Stderr, "%s: %s: too many arguments\n", os.Args[0], cmd)
		os.Exit(2)
	}

	if len(args) < min {
		args = readStdin(args, max)
	}
	if len(args) < min {
		fmt.Fprintf(os.Stderr, "%s: %s: missing argument\n", os.Args[0],
			cmd)
		os.Exit(2)
	}

	return args
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", os.Args[0]))

	parseFlags()

	if cmd := commands[flag.Arg(0)]; cmd != nil {
		log.SetPrefix(fmt.Sprintf("%s: %s: ", os.Args[0], cmd.name))
		if err := cmd.flagSet.Parse(flag.Args()[1:]); err == flag.ErrHelp {
			cmd.flagSet.Usage()
			os.Exit(2)
		} else if err != nil {
			log.Fatal(err)
		}
		args := getArgs(cmd.name, cmd.minArgs, cmd.maxArgs, cmd.flagSet.Args())
		cmd.function(args)
	} else {
		fmt.Fprintf(os.Stderr, "%s: no such command: %s\n", os.Args[0],
			flag.Arg(0))
		os.Exit(2)
	}
}
