package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func main() {
	file, err := os.Open("sample_read_json_data.json")

	if err != nil {
		log.Fatalf("Failed to open json file: %s", err)
	}

	defer file.Close()

	byteData, err := ioutil.ReadAll(file)

	if err != nil {
		log.Fatalf("Failed to read csv file: %s", err)
	}

	var result User
	json.Unmarshal([]byte(byteData), &result)

	fmt.Println(result)
}
