package main

import (
	"fmt"
	"net/url"
	"sort"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

type PageInfo struct {
	URL   string
	Links int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func configure(rawBaseURL string, maxConcurrency int, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}

func printReport(pages map[string]int, baseURL string) {
	var slice []PageInfo
	for url, links := range pages {
		slice = append(slice, PageInfo{URL: url, Links: links})
	}
	sort.Slice(slice, func(i, j int) bool {
		if slice[i].Links == slice[j].Links {
			return slice[i].URL < slice[j].URL // Alphabetical order for ties
		}
		return slice[i].Links > slice[j].Links // Descending order of link counts
	})
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")
	for _, page := range slice {
		fmt.Printf("Found %d internal links to %s\n", page.Links, page.URL)
	}
}
