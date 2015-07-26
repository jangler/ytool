ytool
=====
A command-line interface to the YouTube Data API.

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
	Usage: ytool <cmd> [<arg>]...

	A command-line interface to the YouTube data API. If not enough
	command-line arguments are specified for a command, remaining arguments
	are read from standard input.

	Commands:
	  search <query>...  print the top URL matching <query>
	  title <url>        print the title of the video at <url>

Examples
--------
	$ ytool search ilkae | tee >(ytool title)
	https://youtube.com/watch?v=tCIJPYB3xUU
	Ilkae - Ampersand

License
-------
Public domain.
