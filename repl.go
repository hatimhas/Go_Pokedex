package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var cleanedOutput []string
	lowerCasedOutput := strings.ToLower(text)
	cleanedOutput = strings.Fields(lowerCasedOutput)
	return cleanedOutput
}
