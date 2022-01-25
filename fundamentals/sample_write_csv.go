package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	lines := [][]string{{"Joe Rizal", "Editor in Chief"}, {"Andy Bonifacio", "Supremo"}}
	csvFile, err := os.Create("sample_write_csv_data.csv")

	if err != nil {
		log.Fatalf("Failed creating csv file: %s", err)
	}

	writer := csv.NewWriter(csvFile)

	for _, row := range lines {
		writer.Write(row)
	}

	writer.Flush()
	csvFile.Close()
}
