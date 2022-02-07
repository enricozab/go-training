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

	// Checks if database is load successfully
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
	var index, score int
	var usedProblems []int

	// Prints each question and asks user for the answer
	for ctr := 0; ctr < *questionsCount; ctr++ {
		// Question randomizer
		rand.Seed(time.Now().UnixNano())
		index = rand.Intn(*questionsCount-0+1) + 0

		if !findQuestion(usedProblems, index) {
			// Question that was asked already is saved to a slice - checker to avoid repeating of questions
			usedProblems = append(usedProblems, index)

			fmt.Printf("Q: %s = ", rows[index][0])
			fmt.Scan(&input)

			// Checks if the database has the corresponding right answer to the question
			if len(rows[index]) <= 1 {
				log.Fatalf("No correct answer found. Please check your data.")
			}

			// Increments score if answer is correct
			if input == rows[index][1] {
				score++
			}
		} else {
			ctr--
		}
	}

	// Prints user's final score
	println()
	fmt.Println("You answered ", score, " out of ", *questionsCount, " questions correctly.")
}

// findQuestion checks whether a problem has been questioned
func findQuestion(problems []int, index int) bool {
	for _, i := range problems {
		if i == index {
			return true
		}
	}

	return false
}
