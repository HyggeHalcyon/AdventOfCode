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
 *                         UTILITY
 * =========================================================
 */
func readProblem() []string {
	file, err := os.Open("problem.txt")
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

// calculate the amount distance traveled based upon
// taking consideration of acceleration caused by the
// button press and the rest of the race time
func distanceTraveled(pressed_t int, race_t int) int {
	distance := pressed_t * (race_t - pressed_t)

	return distance
}

// [70 15 30] returns 701530
func concatToInteger(s []string) int {
	format := strings.Repeat("%v", len(s))
	values := make([]interface{}, len(s))

	for i, v := range s {
		values[i] = v
	}

	output := fmt.Sprintf(format, values...)
	ret, _ := strconv.Atoi(output)

	return ret
}

/*
 * =========================================================
 *                       SOLUTIONS
 * =========================================================
 */
func sln_one() int {
	lines := readProblem()
	var sum int = 1
	var line string

	line = strings.Split(lines[0], ":")[1]
	time := strings.Fields(line)

	line = strings.Split(lines[1], ":")[1]
	distance := strings.Fields(line)

	// enumerate through the races
	for r := 0; r < len(time); r++ {
		var ways int

		t, _ := strconv.Atoi(time[r])
		d, _ := strconv.Atoi(distance[r])

		for j := 0; j <= t; j++ {
			if distanceTraveled(j, t) > d {
				ways++
			}
		}

		// fmt.Printf("time:%v distance:%v ways: %v\n", time[i], distance[i], ways)
		sum *= ways
	}

	return sum
}

func sln_two() int {
	lines := readProblem()
	var sum int
	var line string

	line = strings.Split(lines[0], ":")[1]
	time := strings.Fields(line)
	t := concatToInteger(time)

	line = strings.Split(lines[1], ":")[1]
	distance := strings.Fields(line)
	d := concatToInteger(distance)

	for i := 0; i <= t; i++ {
		if distanceTraveled(i, t) > d {
			sum++
		}
	}

	return sum
}

func main() {
	var sln int

	sln = sln_one()
	fmt.Printf("[+] part-1 solution: %v\n", sln) // 140220

	sln = sln_two()
	fmt.Printf("[+] part-2 solution: %v\n", sln) // 39570185
}
