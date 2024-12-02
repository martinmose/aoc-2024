package constants

import "net/url"

// BaseURL is the base URL for the Advent of Code website
var BaseURL = &url.URL{
	Scheme: "https",
	Host:   "adventofcode.com",
	Path:   "/2024/day/",
}
