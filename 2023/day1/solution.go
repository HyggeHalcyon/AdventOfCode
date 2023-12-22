// --- Day 1: Trebuchet?! ---
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
 * =========================================================
 *                       	CONSTANTS
 * =========================================================
 */
var WORD_TO_STRING = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

/*
 * =========================================================
 *                       	UTILITY
 * =========================================================
 */
func isdigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	file.Close()

	return lines
}

func startsWith(s string, key string) bool {
	if len(s) < len(key) {
		return false
	}

	for i := 0; i < len(key); i++ {
		if s[i] != key[i] {
			return false
		}
	}

	return true
}

/*
 * =========================================================
 *                       SOLUTIONS
 * =========================================================
 */
func part_one_sln() int {
	// --- PART ONE ---
	lines := readFile("problem.txt")
	var sum int

	for _, line := range lines {
		digit := make([]byte, 2)

		for _, c := range line {
			if isdigit(c) {
				if digit[0] == 0 {
					digit[0] = byte(c)
				}
				digit[1] = byte(c)
			}
		}

		// fmt.Printf("[!] DEBUG: %s -> %v:%v\n", line, digit[0], digit[1])
		res, _ := strconv.Atoi(string(digit))
		sum += res
	}

	return sum
}

func part_two_sln() int {
	// --- PART TWO ---
	lines := readFile("problem.txt")
	var sum int

	for _, line := range lines {
		var firstDigit, lastDigit int

		for idx, c := range line {
			var num int

			if isdigit(c) {
				num = int(c - '0')
			} else {
				for key, value := range WORD_TO_STRING {
					if startsWith(line[idx:], key) {
						num = value
						break
					}
				}

				if num == 0 {
					continue
				}
			}

			if firstDigit == 0 {
				firstDigit = num
			}
			lastDigit = num
		}

		// fmt.Printf("[!] DEBUG: %s -> %v:%v\n", line, firstDigit, lastDigit)
		res, _ := strconv.Atoi(strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit))
		sum += res
	}

	return sum
}

func main() {
	var sln int

	sln = part_one_sln()
	fmt.Printf("[+] part-1 solution: %v\n", sln) // 55816

	sln = part_two_sln()
	fmt.Printf("[+] part-2 solution: %v\n", sln) // 54980
}
