package main

import "flag"

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	_ = csvFilename
}

// the flag package is a way of getting values from the user also allows the consumer of the binary to check the arguments passed after flag. lets them know what the project is about
