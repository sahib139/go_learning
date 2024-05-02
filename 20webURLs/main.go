package main

import (
	"fmt"
	"net/url"
)

const URL string = "http://13.60.19.44:3001/html/chattingArea.html?roomId={%22id%22:%226628f7516e7ccf44d0102986%22,%22friendName%22:%22SAUMYA%20SINHA%22}"

func main() {
	fmt.Println("web URL learning")
	rawUrl, err := url.Parse(URL)

	checkNilError(err)

	fmt.Println(rawUrl.Scheme)
	fmt.Println(rawUrl.Host)
	fmt.Println(rawUrl.Path)
	fmt.Println(rawUrl.Hostname())
	fmt.Println(rawUrl.Port())
	fmt.Println(rawUrl.RawQuery)

	parseQuery, err := url.ParseQuery(rawUrl.RawQuery)

	checkNilError(err)

	fmt.Println(parseQuery, parseQuery["roomId"])

	//creating new url
	newUrl := &url.URL{
		Scheme: "https",
		Host:   "google.com",
		Path:   "/search",
	}
	newURL := newUrl.String()
	fmt.Println(newURL)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
