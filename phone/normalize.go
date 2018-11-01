package main

import (
	"fmt"
	"regexp"
)

type example struct {
	input  string
	expect string
}

func normalize(phone string) string {
	re := regexp.MustCompile("\\D")
	return re.ReplaceAllString(phone, "")
}

func main() {
	exs := []example{
		{"1234567890", "1234567890"},
		{"(123) 456 7890", "1234567890"},
	}

	for _, ex := range exs {
		fmt.Println(normalize(ex.input))
		fmt.Println(ex.expect)
	}
}
