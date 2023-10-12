package common

import (
	"fmt"
	"os"
	"strconv"
)

type Prompt struct {
}

func (p Prompt) ReadInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

func (p Prompt) ReadInputWithMessage(message string) string {
	fmt.Println(message)
	return p.ReadInput()
}

func (p Prompt) ReadInputForOptions(message string, options []string) string {
	input := p.ReadInputWithMessage(message)

	if !contains(input, options) {
		fmt.Println("Unaceptable input\"", input, "\", acceptable values are:", options)
		os.Exit(1)
	}

	return input
}

func (p Prompt) ReadInt(message string) int {
	input := p.ReadInputWithMessage(message)
	parsedValue, err := strconv.Atoi(input)
	Error{}.CheckErr(err)
	return int(parsedValue)
}

func (p Prompt) ReadFloat(message string) float32 {
	input := p.ReadInputWithMessage(message)
	parsedValue, err := strconv.ParseFloat(input, 32)
	Error{}.CheckErr(err)
	return float32(parsedValue)
}

func contains(value string, array []string) bool {
	for _, e := range array {
		if e == value {
			return true
		}
	}
	return false
}
