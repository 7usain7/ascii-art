package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input(str string) error {
	file, err := os.Open("standard.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	var all_results [][]string
	lines := strings.Split(str, "\\n")
	for _, line := range lines {
		var result []string
		i := 0
		for _, rune := range line {
			file.Seek(0, 0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if int(rune)-32 == i/9 {
					line := scanner.Text()
					result = append(result, line)
				}
				i++
			}
			scanner_err := scanner.Err()
			if scanner_err != nil {
				return scanner_err
			}
			i = 0
		}
		if len(result) > 0 {
			all_results = append(all_results, result)
		} else {
			all_results = append(all_results, []string{})
		}
	}

	for _, result := range all_results {
		if len(result) == 0 {
			fmt.Println()
			continue
		}

		for i := 1; i < 9; i++ {
			for j := range len(result) {
				if i+j*9 < len(result) {
					fmt.Print(result[i+j*9])
				}
			}
			fmt.Println()
		}
	}
	return nil
}
