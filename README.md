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
**ytool**

	Usage: ytool <cmd> [<arg>]...

	A command-line interface to the YouTube data API. If not enough
	command-line arguments are specified for a command, remaining arguments
	are read from standard input.

	For help regarding a specific command, see 'ytool <cmd> -h'.

	Commands:
	  playlist  print the URLs of videos in a playlist
	  search    print the URLs of videos matching a query
	  title     print the title of a video at a URL

**ytool playlist**

	Usage: ytool playlist <url>

	Print the URLs of the videos in the playlist at <url>, up to a maximum
	of 50 videos.

**ytool search**

	Usage: ytool search [<option>]... <query>...

	Search YouTube for <query> (joined by spaces if multiple arguments are
	given) and print the URLs of the top matches in descending order by
	relevance.

	Options:
	  -n=1: maximum number of results, in the range [0, 50]
	  -type="video,channel,playlist": restrict search to given resource types

**ytool title**

	Usage: ytool title <url>

	Print the title of the video at <url>.

Examples
--------
	$ ytool search -type=video ilkae | tee >(ytool title)
	https://youtube.com/watch?v=tCIJPYB3xUU
	Ilkae - Ampersand
	$ ytool search -type=playlist this heat deceit | ytool playlist
	https://www.youtube.com/watch?v=IDAr24H-tKU
	https://www.youtube.com/watch?v=MDq9YHcIip0
	https://www.youtube.com/watch?v=DX-3H3QqwAc
	https://www.youtube.com/watch?v=P0eVTeQi06c
	https://www.youtube.com/watch?v=qu2T1JJTN1U
	https://www.youtube.com/watch?v=n0ySaCs-zJk
	https://www.youtube.com/watch?v=lzZMhAM2SqU
	https://www.youtube.com/watch?v=kAH9u5pXY_Y
	https://www.youtube.com/watch?v=NLMoDU9Tl_E
	https://www.youtube.com/watch?v=e8ztxKXnqwI

License
-------
Public domain.
