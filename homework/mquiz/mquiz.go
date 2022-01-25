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

	db, err := os.Open(*database)

	if err != nil {
		log.Fatalf("Failed to load database: %v", err)
	}

	if filepath.Ext(strings.TrimSpace(*database)) != ".csv" {
		log.Fatalf("Incorrect database format. Database should be in .csv format.")
	}

	defer db.Close()

	reader := csv.NewReader(db)
	rows, _ := reader.ReadAll()

	if len(rows) < *questionsCount {
		log.Fatalf("Insufficient questions. Database should contain at least 10 questions.")
	}

	var input string
	var index, score int
	var usedProblems []int

	for ctr := 0; ctr < *questionsCount; ctr++ {
		rand.Seed(time.Now().UnixNano())
		index = rand.Intn(*questionsCount-0+1) + 0

		if !FindQuestion(usedProblems, index) {
			usedProblems = append(usedProblems, index)

			fmt.Printf("Q: %s = ", rows[index][0])
			fmt.Scan(&input)

			if input == rows[index][1] {
				score++
			}
		} else {
			ctr--
		}
	}

	fmt.Println("You answered ", score, " out of ", *questionsCount, " questions correctly.")
}

// FindQuestion checks whether a problem has been questioned
func FindQuestion(problems []int, index int) bool {
	for _, i := range problems {
		if i == index {
			return true
		}
	}

	return false
}
