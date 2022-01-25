package main

import "fmt"

func main() {
	var languages = map[string]string{
		"HELLO":        "ENGLISH",
		"HOLA":         "SPANISH",
		"HALLO":        "GERMAN",
		"BONJOUR":      "FRENCH",
		"CIAO":         "ITALIAN",
		"ZDRAVSTVUJTE": "RUSSIAN",
	}

	var language string
	var ctr = 1

	_, err := fmt.Scan(&language)

	for err == nil && language != "#" {
		if _, ok := languages[language]; ok {
			fmt.Printf("Case %v: %v\n", ctr, languages[language])
		} else {
			fmt.Printf("Case %v: %v\n", ctr, "UNKNOWN")
		}

		ctr++
		_, err = fmt.Scan(&language)
	}
}
