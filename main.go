package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

const BannerHeight = 8

var bannerURLs = map[string]string{
	"standard":   "https://learn.01founders.co/api/content/root/public/subjects/ascii-art/standard.txt",
	"shadow":     "https://learn.01founders.co/api/content/root/public/subjects/ascii-art/shadow.txt",
	"thinkertoy": "https://learn.01founders.co/api/content/root/public/subjects/ascii-art/thinkertoy.txt",
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	// connect css file
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	text := r.FormValue("text")
	style := r.FormValue("banner")
	if text == "" || style == "" {
		http.Error(w, "Missing text or banner style", http.StatusBadRequest)
		return
	}
	bannerMap, err := loadBannerFromURL(bannerURLs[style])
	if err != nil {
		http.Error(w, "Internal Server Error: could not load banner", http.StatusInternalServerError)
		return
	}
	ascii := generateAscii(text, bannerMap)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, template.HTML(ascii))
}

func loadBannerFromURL(url string) (map[rune][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d", resp.StatusCode)
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

func generateAscii(input string, bannerMap map[rune][]string) string {
	lines := strings.Split(input, "\n") // FIXED newline handling
	var result strings.Builder

	for _, line := range lines {
		for i := 0; i < BannerHeight; i++ {
			for _, ch := range line {
				if art, ok := bannerMap[ch]; ok {
					result.WriteString(art[i])
					// fmt.Printf("Char: %c, Line: %s\n", ch, art[i])
				} else {
					// fallback: use space character
					result.WriteString(bannerMap[' '][i])
				}
			}
			result.WriteString("\n")
		}
	}
	// fmt.Printf(result.String())
	return result.String()
}
