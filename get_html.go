package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	httpResponse, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	defer httpResponse.Body.Close()
	if httpResponse.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP status code error: %d", httpResponse.StatusCode)
	}
	mediaType := httpResponse.Header.Get("Content-Type")
	if !strings.Contains(mediaType, "text/html") {
		return "", fmt.Errorf("wrong content-type - text/html required")
	}

	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
