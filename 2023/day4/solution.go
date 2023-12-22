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
 *                         UTILITY
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
func sln_one() int {
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

func sln_two() int {
	lines := readFile("problem.txt")
	scratchcards := make([]int, len(lines))
	var sum int

	for i, line := range lines {
		scratchcards[i] += 1
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

		for w := 0; w < wins; w++ {
			scratchcards[w+i+1] += scratchcards[i]
		}
	}

	for s := 0; s < len(scratchcards); s++ {
		sum += scratchcards[s]
	}

	return sum
}

func main() {
	var sln int

	sln = sln_one()
	fmt.Printf("[+] part-1 solution: %v\n", sln) // 19855

	sln = sln_two()
	fmt.Printf("[+] part-2 solution: %v\n", sln) // 10378710
}
