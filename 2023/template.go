// --- Day ?: ??? ---
package main

import (
	"bufio"
	"fmt"
	"os"
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
	// lines := readFile("sample_one.txt")
	var sum int

	return sum
}

func part_two_sln() int {
	// --- PART TWO ---
	// lines := readFile("problem.txt")
	var sum int

	return sum
}

func main() {
	var sln int

	sln = part_one_sln()
	fmt.Printf("[+] part-1 solution: %v\n", sln) //

	sln = part_two_sln()
	fmt.Printf("[+] part-2 solution: %v\n", sln) //
}
