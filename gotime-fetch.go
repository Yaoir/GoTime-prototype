package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"time"
)

type Message struct {
	Data string `json:"data"`
	Text string `json:"text"`
}

// A very simple regex to match the date/time string. Used for testing and development.
// Use it if the RFC 3339 regex fails.

const date_time_regex string = `2[0-9][0-9][0-9].*[Zz]`

// RFC 3339 regex from https://gist.github.com/marcelotmelo/b67f58a08bee6c2468f8

const rfc3339_regex   string = `([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])[Tt]([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]|60)(\.[0-9]+)?(([Zz])|([\+|\-]([01][0-9]|2[0-3]):[0-5][0-9]))`

// ISO 8601 regex from https://stackoverflow.com/questions/28020805/

const iso8601_regex   string = `(?:[1-9]\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\d|2[0-8])|(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\d(?:0[48]|[2468][048]|[13579][26])|(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d(?:Z|[+-][01]\d:[0-5]\d)`

// URL for Go Time API

const url string = "https://changelog.com/slack/countdown/gotime"

// For development and testing:
//const url string = "http://localhost/gotime-mock.html"

func main() {

	gotimeClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr,"%v\n",err)
		os.Exit(2)
	}

	req.Header.Set("User-Agent", "fetch.go")

	res, getErr := gotimeClient.Do(req)
	if getErr != nil {
		fmt.Fprintf(os.Stderr,"%v\n",getErr)
		os.Exit(2)
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		fmt.Fprintf(os.Stderr,"%v\n",readErr)
		os.Exit(2)
	}

	gotime_json := Message{}
	jsonErr := json.Unmarshal(body, &gotime_json)
	if jsonErr != nil {
		fmt.Fprintf(os.Stderr,"%v\n",jsonErr)
		os.Exit(2)
	}

	if len(gotime_json.Data) == 0 {
		// There is currently no show scheduled. Print the message and exit.
		fmt.Fprintf(os.Stderr,"%s\n",gotime_json.Text)
		os.Exit(1)
	}

	/* gotime_json.Data contains an RFC3339 time string. Find it. */

	id := regexp.MustCompile(rfc3339_regex)
	rfc3339 := id.Find([]byte(gotime_json.Data[:]))

	if len(rfc3339) == 0 {
		fmt.Fprintf(os.Stderr,"Failed to extract date and time from JSON.\n")
		os.Exit(2)
	}

	fmt.Printf("%s\n",rfc3339)
}
