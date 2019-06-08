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

// If the full version of the regex fails, try the simple version

const rfc3339_simple string = `2[0-9][0-9][0-9].*[Zz]`
const rfc3339_full   string = `([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])[Tt]([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]|60)(\.[0-9]+)?(([Zz])|([\+|\-]([01][0-9]|2[0-3]):[0-5][0-9]))`

//const url string = "https://changelog.com/slack/countdown/gotime"
// For development and testing:
const url string = "http://localhost/gotime-mock.html"

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

	id := regexp.MustCompile(rfc3339_full)
	rfc3339 := id.Find([]byte(gotime_json.Data[:]))
	fmt.Printf("%s\n",rfc3339)
}
