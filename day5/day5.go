package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* Helper function to check if an int[] has an elem in int[]*/
func contains(rules []int, arr []int) bool {
	for _, rule := range rules {
		for _, elem := range arr {
			if elem == rule {
				return true
			}
		}
	}
	return false
}

/* Helper function to convert a []string -> []int */
func convert_str_to_int_arr(line []string) ([]int, error) {
	ints := make([]int, len(line))
	for idx, item := range line {
		num, err := strconv.Atoi(item)
		if err != nil {
			return nil, err
		}
		ints[idx] = num
	}
	return ints, nil
}

/* Main function to return the middle element if valid, otherwise 0 */
func return_valid_mid_elem(line []string, graph map[int][]int) int {
	/* Convert the []string to []int */
	nums, err := convert_str_to_int_arr(line)
	if err != nil {
		return 0
	}
	for idx, num := range nums {
		rules := graph[num]
		if contains(rules, nums[:idx]) {
			return 0
		}
	}
	return nums[len(nums)/2]
}

func main() {
	/* 0. Open the file for reading */
	graph := make(map[int][]int)
	file, err := os.Open("day5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	/* 1. Building the graph */
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		edges := strings.Split(line, "|")
		node1, err := strconv.Atoi(edges[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		node2, err := strconv.Atoi(edges[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		graph[node1] = append(graph[node1], node2)
	}
	valid_middle_num_sum := 0
	/* 2. Checking if the lists are valid based on the map */
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")
		valid_middle_num_sum += return_valid_mid_elem(nums, graph)
	}
	fmt.Println(valid_middle_num_sum)
	return
}
