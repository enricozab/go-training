package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("sample_read_csv_data.csv")

	if err != nil {
		log.Fatalf("Failed to open csv file: %s", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		log.Fatalf("Failed to parse csv file: %s", err)
	}

	fmt.Println(lines)
}
