package main

// gotime2dur.go
// Convert an RFC3339 time string to a Go time.Duration.
// The duration is the amount of time between the current time and the specified time.

import (
	"fmt"
	"os"
	"time"
	)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("gotime2dur: please provide an RFC3339 time string as the argument\n")
		os.Exit(2)
	}

	// The timer command accepts a Duration and counts down to zero.
	// We need to subtract the time of the show from the current time to get a Duration

	now := time.Now().UTC()

	// Parse the text-format show time into show_time, (type Time).
	// The text is scraped from a web page (http://changelog.com/live)
	// by the gotime-scraper shell script, and is _assumed_ to be in RFC3339 format.

	show_time, err := time.Parse(time.RFC3339,os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr,"gotime2dur: bad RFC3339 time string %q\n",os.Args[1])
		os.Exit(2)
	}

	if now.After(show_time) {
		// The current time (now) is after the start of the show.
		// This week's show has already started or has ended,
		// so print a duration of zero seconds.
		fmt.Printf("0s\n")
		os.Exit(0)
	}

	// Print the difference between the show time and now.
	fmt.Printf("%v\n",show_time.Sub(now))
}
