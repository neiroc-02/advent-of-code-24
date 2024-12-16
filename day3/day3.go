/* day3.go */
/* https://adventofcode.com/2024/day/3 */
/* https://learnxinyminutes.com/go/ */
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func evaluate_muls(valid_muls []string, can_run *bool) int {
	total := 0
	pattern := regexp.MustCompile(`[0-9]+`)
	for idx := range len(valid_muls) {
		if valid_muls[idx] == "do()" {
			*can_run = true
			continue
		} else if valid_muls[idx] == "don't()" {
			*can_run = false
			continue
		} else if *can_run {
			nums := pattern.FindAllString(valid_muls[idx], -1)
			if len(nums) == 2 {
				num1, err1 := strconv.Atoi(nums[0])
				num2, err2 := strconv.Atoi(nums[1])
				if err1 != nil || err2 != nil {
					fmt.Println(err1)
					fmt.Println(err2)
					return -1
				}
				total += num1 * num2
			}
		}
	}
	return total
}

func main() {
	file, err := os.Open("day3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	total := 0
	can_run := true
	scanner := bufio.NewScanner(file)
	//pattern1 := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`) //use this pattern to get part1 answer
	pattern3 := regexp.MustCompile(`(mul\([0-9]+,[0-9]+\))|(do\(\))|(don't\(\))`)
	for scanner.Scan() {
		line := scanner.Text()
		valid_muls := pattern3.FindAllString(line, -1)
		total += evaluate_muls(valid_muls, &can_run)
	}
	fmt.Println(total)
	return
}
