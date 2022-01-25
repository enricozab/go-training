package main

import "flag"

func main() {
	csvFilename := flag.String("csv", "sample_problem_flag.csv", "csvin format 'question,answer'")
	flag.Parse()
	_ = csvFilename
}
