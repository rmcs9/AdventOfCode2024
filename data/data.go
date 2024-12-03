package data

import (
	"os"
	"log"
	"bufio"
)

func Get(datfile string) []string {
	file, err := os.Open(datfile)
	if err != nil {
		log.Fatal(err)
	}

	
	scanner := bufio.NewScanner(bufio.NewReader(file))

	var input []string 

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}
