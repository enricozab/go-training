package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	database := flag.String("csv", "problems.csv", "database")
	questionsCount := flag.Int("n", 10, "number of questions")

	flag.Parse()

	filePointer, err := os.Open(*database)

	// Checks if database loaded successfully
	if err != nil {
		log.Fatalf("Failed to load database: %v", err)
	}

	// Checks if the format of the database is csv
	if filepath.Ext(strings.TrimSpace(*database)) != ".csv" {
		log.Fatalf("Incorrect database format. Database should be in .csv format.")
	}

	defer filePointer.Close()

	reader := csv.NewReader(filePointer)
	rows, _ := reader.ReadAll()

	// Checks if database has a least 10 questions
	if len(rows) < *questionsCount {
		log.Fatalf("Insufficient questions. Database should contain at least 10 questions.")
	}

	var input string
	var score int

	// Shuffles the questions
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(rows), func(i, j int) {
		rows[i], rows[j] = rows[j], rows[i]
	})

	// Prints each question and asks user for the answer
	for ctr := 0; ctr < *questionsCount; ctr++ {
		fmt.Printf("Q: %s = ", rows[ctr][0])
		fmt.Scan(&input)

		// Adds 1 to the score if user's answer is correct
		if input == rows[ctr][1] {
			score++
		}
	}

	// Prints user's final score
	println()
	fmt.Println("You answered ", score, " out of ", *questionsCount, " questions correctly.")
}
