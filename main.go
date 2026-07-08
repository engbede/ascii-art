package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . \"text\" [banner]")
		return
	}

	bannerFile := "banner/standard.txt"

	if len(os.Args) == 3 {
		bannerFile = "banner/" + os.Args[2] + ".txt"
	}

	banner, err := loadBanner(bannerFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	input := strings.ReplaceAll(os.Args[1], "\\n", "\n")

	printAsciiArt(input, banner)
}
