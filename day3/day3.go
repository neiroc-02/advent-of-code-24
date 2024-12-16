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

// "strconv"
func evaluate_muls(valid_muls []string) int {
	total := 0
	pattern := regexp.MustCompile(`[0-9]+`)
	for idx := range len(valid_muls) {
		nums := pattern.FindAllString(valid_muls[idx], -1)
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		total += num1 * num2
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
	scanner := bufio.NewScanner(file)
	pattern := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`) //regex for mul(num,num)
	for scanner.Scan() {
		line := scanner.Text()
		valid_muls := pattern.FindAllString(line, -1)
		total += evaluate_muls(valid_muls)
	}
	fmt.Println(total)
	return
}
