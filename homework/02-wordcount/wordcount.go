package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
)

type WordCounter struct {
	mu    sync.Mutex
	words map[string]int
}

func main() {
	database := flag.String("csv", "words.csv", "database")

	flag.Parse()

	filePointer, err := os.Open(*database)

	// Checks if database is load successfully
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Checks if the format of the database is csv
	if filepath.Ext(strings.TrimSpace(*database)) != ".csv" {
		log.Fatalf("Incorrect database format. Database should be in .csv format.")
	}

	defer filePointer.Close()

	reader := csv.NewReader(filePointer)
	rows, _ := reader.ReadAll()

	// Checks if database contains data
	if len(rows) <= 0 {
		log.Fatalf("No data found.")
	}

	wordCount := WordCounter{words: make(map[string]int)}
	c := make(chan string)

	// Formats each word and counts each word
	for index := 0; index < len(rows); index++ {
		go wordFormat(rows[index][0], c)
		go wordCount.wordCount(<-c)
	}

	// Sort the words alphabetically
	wordKeys := make([]string, 0, len(wordCount.words))
	for k := range wordCount.words {
		wordKeys = append(wordKeys, k)
	}
	sort.Strings(wordKeys)

	// Prints each word and its corresponding count
	for _, word := range wordKeys {
		fmt.Printf("%v %v\n", word, wordCount.getWordCount(word))
	}
}

// wordFormat formats the word to lower case and removes punctuation marks and extra white spaces
func wordFormat(word string, c chan string) {
	word = strings.TrimSpace(word)
	word = strings.ToLower(word)
	regex := regexp.MustCompile(`[[:punct:]]`)
	word = regex.ReplaceAllString(word, "")

	c <- word
}

// wordCount safely increases the count of the word
func (wordCount *WordCounter) wordCount(key string) {
	wordCount.mu.Lock()
	wordCount.words[key]++
	wordCount.mu.Unlock()
}

// getWordCount safely gets the final count/frequency of the word
func (wordCount *WordCounter) getWordCount(key string) int {
	wordCount.mu.Lock()
	defer wordCount.mu.Unlock()
	return wordCount.words[key]
}
