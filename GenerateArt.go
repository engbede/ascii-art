package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(path string) (map[rune][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		fmt.Println("empty bannerFile")
	}
	lines := strings.Split(string(data), "\n")

	if len(lines) != 855 {
		fmt.Println("imcomplete bannerFile")
	}

	if len(lines) > 0 && lines[0] == "" {
		lines = lines[1:]
	}

	banner := make(map[rune][]string)

	char := 32
	for i := 0; i < len(lines); i += 9 {
		if i+8 >= len(lines) {
			break
		}

		banner[rune(char)] = lines[i : i+8]
		char++
	}

	return banner, nil
}
