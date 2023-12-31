package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * =========================================================
 *                        CONSTANTS
 * =========================================================
 */
const MAXRED = 12
const MAXGREEN = 13
const MAXBLUE = 14

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

	// iterating over the games
	for idx, line := range lines {
		var red, green, blue int

		// parsing into a list of cubes
		game := strings.Split(line, ": ")[1]
		game = strings.Replace(game, ", ", ",", -1)
		game = strings.Replace(game, "; ", ",", -1)

		// iterating over the cubes
		for _, cube := range strings.Split(game, ",") {
			subs := strings.Split(cube, " ")
			num, _ := strconv.Atoi(subs[0])
			color := subs[1]

			switch color {
			case "red":
				if num > red {
					red = num
				}
			case "green":
				if num > green {
					green = num
				}
			case "blue":
				if num > blue {
					blue = num
				}
			default:
				fmt.Println("[X] Error: unknown color")
			}
		}

		// check if any of the cube set goes over the ceiling
		if red > MAXRED || green > MAXGREEN || blue > MAXBLUE {
			continue // skip
		} else {
			sum += idx + 1
		}
	}

	return sum
}

func sln_two() int {
	lines := readFile("problem.txt")
	var sum int

	// iterating over the games
	for _, line := range lines {
		var red, green, blue int
		game := strings.Split(line, ": ")[1]
		game = strings.Replace(game, ", ", ",", -1)
		game = strings.Replace(game, "; ", ",", -1)

		// iterating over the cubes
		for _, cube := range strings.Split(game, ",") {
			subs := strings.Split(cube, " ")
			num, _ := strconv.Atoi(subs[0])
			color := subs[1]

			switch color {
			case "red":
				if num > red {
					red = num
				}
			case "green":
				if num > green {
					green = num
				}
			case "blue":
				if num > blue {
					blue = num
				}
			default:
				fmt.Println("[X] Error: unknown color")
			}
		}

		sum += red * green * blue
	}

	return sum
}

func main() {
	var sln int

	sln = sln_one()
	fmt.Printf("[+] part-1 solution: %v\n", sln) // 2727

	sln = sln_two()
	fmt.Printf("[+] part-2 solution: %v\n", sln) // 56580
}
