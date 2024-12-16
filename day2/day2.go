/* day2.go */
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

func check_monotonic_condition(increasing *int, curr_difference int) bool {
	if *increasing == -1 {
		if curr_difference > 0 {
			*increasing = 1
			return true
		} else {
			*increasing = 0
			return true
		}
	} else if *increasing == 1 && curr_difference > 0 {
		return true
	} else if *increasing == 0 && curr_difference < 0 {
		return true
	} else {
		return false
	}
}

func check_steepness_condition(prev_num int, curr_num int, curr_difference *int) bool {
	*curr_difference = curr_num - prev_num
	upper_bound := Abs(*curr_difference) <= 3
	lower_bound := Abs(*curr_difference) != 0
	if !upper_bound || !lower_bound {
		return true
	}
	return false
}

func is_row_safe(nums []string) bool {
	increasing := -1
	prev_num := -1
	for idx := range len(nums) {
		curr_num, err := strconv.Atoi(nums[idx])
		if err != nil {
			fmt.Println(err)
			return false
		}
		var curr_difference int
		if prev_num == -1 {
			prev_num = curr_num
			continue
		}
		is_steep := check_steepness_condition(prev_num, curr_num, &curr_difference)
		if is_steep {
			return false
		}
		is_monotonic := check_monotonic_condition(&increasing, curr_difference)
		if is_monotonic {
			prev_num = curr_num
		} else {
			return false
		}
	}
	return true
}

func part1(file *os.File) int {
	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		safe := is_row_safe(nums)
		if safe == true {
			count++
		}
	}
	return count
}

func part2(file *os.File) int {
	var count int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		safe := is_row_safe(nums)
		if safe == true {
			count++
		} else {
			for i := 0; i < len(nums); i++ {
				curr_slice := []string{}
				curr_slice = append(curr_slice, nums[:i]...)
				curr_slice = append(curr_slice, nums[i+1:]...)
				safe = is_row_safe(curr_slice)
				if safe {
					count++
					break
				}
			}
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
	file2, err := os.Open("day2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(part2(file2))
}
