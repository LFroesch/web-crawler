package main

import (
	"net/url"
	"strings"
)

func normalizeURL(urlString string) (string, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	host := parsedURL.Host
	path := parsedURL.Path
	if path == "/" || path == "" {
		return host, nil
	}
	path = strings.TrimSuffix(path, "/")
	return host + path, nil
}
