package main

import (
	"fmt"
	"strings"
)

func printAsciiArt(input string, banner map[rune][]string) {
	lines := strings.Split(input, "\n")
	if len(lines) > 1 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}

		for i := 0; i < 8; i++ {
			for _, ch := range line {
				if art, ok := banner[ch]; ok {
					fmt.Print(art[i])
				}
			}
			fmt.Println()
		}
	}
}
