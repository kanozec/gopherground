package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type problem struct {
	question string
	ans      string
}

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Failed to open the CSV file")
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the CSV file")
	}
	quiz := make([]problem, len(lines))
	for i, line := range lines {
		quiz[i] = problem{
			question: line[0],
			ans:      line[1],
		}

		fmt.Println(quiz[i])
	}
}
