package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input(str string) error {
	bannerMap, err := loadBanner("models/standard.txt")
	if err != nil {
		return err
	}

	lines := strings.Split(str, "\\n")

	allEmpty := true
	for _, line := range lines {
		if line != "" {
			allEmpty = false
			break
		}
	}

	if allEmpty {
		for i := 0; i < len(lines)-1; i++ {
			fmt.Println()
		}
		return nil
	}

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}

		for row := range 8 {
			for _, r := range line {
				if art, exists := bannerMap[r]; exists {
					fmt.Print(art[row])
				}
			}
			fmt.Println()
		}
	}

	return nil
}

func loadBanner(filename string) (map[rune][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bannerMap := make(map[rune][]string)

	currentChar := rune(32)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		if len(lines) == 9 {
			bannerMap[currentChar] = lines[1:]
			currentChar++
			lines = nil
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return bannerMap, nil
}
