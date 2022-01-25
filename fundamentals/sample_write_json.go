package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type User struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

func main() {
	data := User{Name: "Joe Rizal", Job: "Editor in Chief"}
	byteData, err := json.Marshal(data)

	if err != nil {
		log.Fatalf("Failed to Marshal data: %s", err)
	}

	ioutil.WriteFile("sample_write_json_data.json", byteData, 0644)
}
