package main

import (
	"bufio"
	"fmt"
	"net/http"
)

const (
	BannerHeight = 8
)

var bannerURLs = map[string]string{
	"standard":   "https://learn.01founders.co/api/content/root/public/subjects/ascii-art/standard.txt",
	"shadow":     "https://learn.01founders.co/api/content/root/public/subjects/ascii-art/shadow.txt",
	"thinkertoy": "https://learn.01founders.co/api/content/root/public/subjects/ascii-art/thinkertoy.txt",
}

// func main() {
// 	if len(os.Args) < 2 {
// 		return
// 	}
// 	input := os.Args[1]

// 	// Select banner style here: "standard", "shadow", or "thinkertoy"
// 	// style := "standard"

// 	style := os.Args[2]

// 	bannerMap, err := loadBannerFromURL(bannerURLs[style])
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error loading banner: %v\n", err)
// 		os.Exit(1)
// 	}

// 	lines := strings.Split(input, "\\n")

// 	for _, line := range lines {
// 		printAsciiLine(line, bannerMap)
// 	}
// }

func loadBannerFromURL(url string) (map[rune][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download banner file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to download banner file: HTTP %d", resp.StatusCode)
	}

	scanner := bufio.NewScanner(resp.Body)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	bannerMap := make(map[rune][]string)
	charCount := (len(lines) + 1) / (BannerHeight + 1)

	for i := 0; i < charCount; i++ {
		start := i * (BannerHeight + 1)
		if start+BannerHeight > len(lines) {
			break
		}
		bannerMap[rune(32+i)] = lines[start : start+BannerHeight]
	}

	return bannerMap, nil
}

func printAsciiLine(line string, bannerMap map[rune][]string) {
	for i := 0; i < BannerHeight; i++ {
		for _, ch := range line {
			art, ok := bannerMap[ch]
			if !ok {
				// If character is not found, fallback to space character or a blank string
				if spaceArt, ok := bannerMap[' ']; ok {
					fmt.Print(spaceArt[i])
				} else {
					fmt.Print(" ")
				}
			} else {
				fmt.Print(art[i])
			}
		}
		fmt.Println()
	}
}
