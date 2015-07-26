ytool
=====
A command-line interface to the YouTube data API.

Status
------
Alpha. Not much is implemented, and breaking changes may occur.

Installation
------------
Install or update via the [go command](http://golang.org/cmd/go/):

	go get -u github.com/jangler/ytool

<!-- TODO: mention binary releases once they're available -->

Usage
-----
`ytool`

	Usage: ytool <cmd> [<arg>]...

	A command-line interface to the YouTube data API. If not enough
	command-line arguments are specified for a command, remaining arguments
	are read from standard input.

	For help regarding a specific command, see 'ytool <cmd> -h'.

	Commands:
	  playlist  print the URLs of videos in a playlist
	  search    print the URLs of videos matching a query
	  title     print the title of a video at a URL

`ytool playlist`

	Usage: ytool playlist <url>

	Print the URLs of the videos in the playlist at <url>, up to a maximum
	of 50 videos.

`ytool search`

	Usage: ytool search [<option>]... <query>...

	Search YouTube for <query> (joined by spaces if multiple arguments are
	given) and print the URLs of the top matches in descending order by
	relevance.

	Options:
	  -n=1: maximum number of results, in the range [0, 50]

`ytool title`

	Usage: ytool title <url>

	Print the title of the video at <url>.

Examples
--------
	$ ytool search ilkae | tee >(ytool title)
	https://youtube.com/watch?v=tCIJPYB3xUU
	Ilkae - Ampersand

License
-------
Public domain.
