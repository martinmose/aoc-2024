package utils

import (
	"aoc_2024/constants"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// HTTPGet performs a GET request to the Advent of Code website
func HTTPGet(path string) (string, error) {
	fullURL := constants.BaseURL.ResolveReference(&url.URL{Path: path})

	client := &http.Client{}

	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	envName := "AOC_SESSION"
	session := os.Getenv(envName)
	if session == "" {
		return "", fmt.Errorf("error: %s environment variable not set", envName)
	}

	req.Header.Set("Cookie", "session="+session)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error performing request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
