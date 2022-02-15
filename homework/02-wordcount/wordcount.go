package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	if len(os.Args) < 2 {
		log.Fatalln("Missing parameter, provide file name/s.")
		return
	}

	wordCount := WordCounter{words: make(map[string]int)}

	// Read each file in the argument
	for ctr := 1; ctr < len(os.Args); ctr++ {
		data, err := ioutil.ReadFile(os.Args[ctr])
		c := make(chan string)

		// Throws an error if file cannot be read or if file is not existing
		if err != nil {
			log.Fatalln("Cannot read file or file is missing:", os.Args[ctr])
			return
		}

		// Converts data read from the file from bytes to string of slice
		slicedData := strings.Split(string(data), "\n")

		// Formats each word and counts each word
		for _, word := range slicedData {
			go WordFormat(word, c)
			go wordCount.WordCount(<-c)
		}
	}

	time.Sleep(time.Second)

	// Sort the words alphabetically
	wordKeys := make([]string, 0, len(wordCount.words))
	for k := range wordCount.words {
		wordKeys = append(wordKeys, k)
	}
	sort.Strings(wordKeys)

	// Prints each word and its corresponding count
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
