package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	if len(os.Args) != 4 {
		fmt.Println("usage is: go run . [website (url)] [maxConcurrency (int)] [maxPages (int)]")
		return
	}

	rawBaseURL := os.Args[1]
	maxConcurrency := os.Args[2]
	maxPages := os.Args[3]

	maxConcurrencyInt, err := strconv.Atoi(maxConcurrency)
	if err != nil {
		fmt.Printf("Error: maxConcurrency must be a number, got %q\n", maxConcurrency)
		return
	}
	maxPagesInt, err := strconv.Atoi(maxPages)
	if err != nil {
		fmt.Printf("Error: maxPages must be a number, got %q\n", maxPages)
		return
	}
	if maxConcurrencyInt <= 0 {
		fmt.Printf("Error: maxConcurrency must be greater than 0, got %d\n", maxConcurrencyInt)
		return
	}
	if maxPagesInt <= 0 {
		fmt.Printf("Error: maxPages must be greater than 0, got %d\n", maxPagesInt)
		return
	}

	cfg, err := configure(rawBaseURL, maxConcurrencyInt, maxPagesInt)
	if err != nil {
		fmt.Printf("Error - configure: %v\n", err)
		return
	}
	fmt.Printf("starting crawl of: %s\n", rawBaseURL)
	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	printReport(cfg.pages, rawBaseURL)
	fmt.Printf("Report generation took %v\n", time.Since(start))
}
