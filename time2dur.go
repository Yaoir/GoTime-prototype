package main

import (
	"fmt"
	"os"
	"time"
	)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("time2dur: please provide an argument\n")
		os.Exit(2)
	}

	// The timer command accepts a Duration and counts down to zero.
	// We need to subtract the time of the show from the current time to get a Duration

	now := time.Now().UTC()

	// Parse the text-format show time into show_time, (type Time).
	// The text is scraped from a web page (http://changelog.com/live)
	// by the getgotime shell script, and is _assumed_ to be in RFC3339 format.

	show_time, err := time.Parse(time.RFC3339,os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr,"alarm: bad RFC3339 time string %q\n",os.Args[1])
		os.Exit(2)
	}

	if now.After(show_time) {
		// This week's show has already started or has ended. Print a duration of zero.
		fmt.Printf("0s\n")
		os.Exit(0)
	}

	// print the difference between the show time and now.
	fmt.Printf("%v\n",show_time.Sub(now))
}
