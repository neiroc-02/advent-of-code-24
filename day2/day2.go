/* day1.go */
/* https://adventofcode.com/2024/day/2 */
/* https://learnxinyminutes.com/go/ */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func part1(file *os.File) int {
	var count int
	var safe bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		safe = true
		increasing := -1 // -1 uninitialized, 1 true, 0 false
		line := scanner.Text()
		nums := strings.Fields(line)
		prev_num := -1
		for idx := range len(nums) {
			curr_num, err := strconv.Atoi(nums[idx])
			if err != nil {
				fmt.Println(err)
				return -1
			}
			if prev_num == -1 {
				prev_num = curr_num
				continue
			}
			curr_difference := curr_num - prev_num
			upper_bound := Abs(curr_difference) <= 3
			lower_bound := Abs(curr_difference) != 0
			if !upper_bound || !lower_bound {
				safe = false
				break
			}
			if increasing == -1 {
				if curr_difference > 0 {
					increasing = 1
					prev_num = curr_num
					continue
				} else {
					increasing = 0
					prev_num = curr_num
					continue
				}
			} else if increasing == 1 && curr_difference > 0 && upper_bound && lower_bound {
				prev_num = curr_num
				continue
			} else if increasing == 0 && curr_difference < 0 && upper_bound && lower_bound {
				prev_num = curr_num
				continue
			} else {
				safe = false
				break
			}

		}
		if safe == true {
			count++
		}
	}
	return count
}

func main() {
	/* Open the file for reading */
	file, err := os.Open("day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(part1(file))
}
