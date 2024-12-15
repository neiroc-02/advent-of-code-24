/* day1.go */
/* https://adventofcode.com/2024/day/1 */
/* https://learnxinyminutes.com/go/ */
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* No abs function in Go for ints */
func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	/* Setup file-io for Go */
	file, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	/* Declare the two lists to take the difference of */
	var list1, list2 []int
	count := map[int]int{}
	_ = count
	/* Read the file */
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)
		if len(nums) == 2 {
			num1, err1 := strconv.Atoi(nums[0])
			num2, err2 := strconv.Atoi(nums[1])
			count[num2]++
			if err1 != nil || err2 != nil {
				fmt.Println(err1)
				fmt.Println(err2)
				return
			}
			list1 = append(list1, num1)
			list2 = append(list2, num2)
		}
	}

	/* Sort the lists of integers */
	sort.Ints(list1[:])
	sort.Ints(list2[:])

	/* Calculate answers */
	var part1, part2 int
	for i := range list1 {
		part1 += Abs(list2[i] - list1[i])
		part2 += count[list1[i]] * list1[i]
	}
	fmt.Println(part1)
	fmt.Println(part2)
	return
}
