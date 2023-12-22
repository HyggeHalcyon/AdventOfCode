// --- Day 4: Scratchcards ---
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * =========================================================
 *                       	CONSTANTS
 * =========================================================
 */

/*
 * =========================================================
 *                       	UTILITY
 * =========================================================
 */
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
		scratchcard := strings.Split(line, ": ")[1]

		left := strings.Split(scratchcard, " |")[0]
		right := strings.Split(scratchcard, "| ")[1]

		var wcard []int
		var wins int

		for _, card := range strings.Split(left, " ") {
			if card == "" {
				continue
			}

			num, _ := strconv.Atoi(card)
			wcard = append(wcard, num)
		}

		for _, card := range strings.Split(right, " ") {
			num, _ := strconv.Atoi(card)

			for _, win := range wcard {
				if num == win {
					wins++
					break
				}
			}
		}

		sum += int(math.Pow(2, float64(wins-1)))
	}

	return sum
}

func part_two_sln() int {
	// --- PART TWO ---
	lines := readFile("problem.txt")
	scratchcards := make([]int, len(lines))
	var sum int

	for id, line := range lines {
		scratchcards[id] += 1
		scratchcard := strings.Split(line, ": ")[1]

		left := strings.Split(scratchcard, " |")[0]
		right := strings.Split(scratchcard, "| ")[1]

		var wcard []int
		var wins int

		// record all winning cards
		for _, card := range strings.Split(left, " ") {
			if card == "" {
				continue
			}

			num, _ := strconv.Atoi(card)
			wcard = append(wcard, num)
		}

		// count card that matches winning cards
		for _, card := range strings.Split(right, " ") {
			num, _ := strconv.Atoi(card)

			for _, wcard := range wcard {
				if num == wcard {
					wins++
					break
				}
			}
		}

		for i := 0; i < wins; i++ {
			scratchcards[i+id+1] += scratchcards[id]
		}
	}

	for i := 0; i < len(scratchcards); i++ {
		sum += scratchcards[i]
	}

	return sum
}

func main() {
	var sln int

	sln = part_one_sln()
	fmt.Printf("[+] part-1 solution: %v\n", sln) // 19855

	sln = part_two_sln()
	fmt.Printf("[+] part-2 solution: %v\n", sln) // 10378710
}
