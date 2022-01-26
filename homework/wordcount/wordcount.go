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
	"time"
)

type WordCounter struct {
	mu    sync.Mutex
	words map[string]int
}

func main() {
	database := flag.String("csv", "words.csv", "database")

	flag.Parse()

	filePointer, err := os.Open(*database)

	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if filepath.Ext(strings.TrimSpace(*database)) != ".csv" {
		log.Fatalf("Incorrect database format. Database should be in .csv format.")
	}

	defer filePointer.Close()

	reader := csv.NewReader(filePointer)
	rows, _ := reader.ReadAll()

	if len(rows) <= 0 {
		log.Fatalf("No data found.")
	}

	wordCount := WordCounter{words: make(map[string]int)}
	c := make(chan string)

	for index := 0; index < len(rows); index++ {
		go WordFormat(rows[index][0], c)
		formattedWord := <-c

		go wordCount.WordCount(formattedWord)
	}

	time.Sleep(time.Second)

	// Sort the words alphabetically
	wordKeys := make([]string, 0, len(wordCount.words))
	for k := range wordCount.words {
		wordKeys = append(wordKeys, k)
	}
	sort.Strings(wordKeys)

	for _, word := range wordKeys {
		fmt.Printf("%v %v\n", word, wordCount.GetWordCount(word))
	}
}

// WordFormat formats the word to lower case and removes punctuation marks and extra white spaces
func WordFormat(word string, c chan string) {
	word = strings.TrimSpace(word)
	word = strings.ToLower(word)
	regex := regexp.MustCompile(`[[:punct:]]`)
	word = regex.ReplaceAllString(word, "")

	c <- word
}

// WordCount safely increases the count of the word
func (wordCount *WordCounter) WordCount(key string) {
	wordCount.mu.Lock()
	wordCount.words[key]++
	wordCount.mu.Unlock()
}

// GetWordCount safely gets the final count/frequency of the word
func (wordCount *WordCounter) GetWordCount(key string) int {
	wordCount.mu.Lock()
	defer wordCount.mu.Unlock()
	return wordCount.words[key]
}
