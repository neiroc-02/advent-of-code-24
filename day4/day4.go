package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	UNINIT = iota
	RIGHT
	LEFT
	DOWN
	UP
	UPRIGHT
	UPLEFT
	DOWNRIGHT
	DOWNLEFT
)

const WORD = "XMAS"

var count = 0

/* Part 1 Search */
func search(direction int, i int, j int, word_idx int, matrix [][]byte) {
	if word_idx > 3 {
		count++
		return
	}
	if i >= len(matrix) || i < 0 || j >= len(matrix[0]) || j < 0 || WORD[word_idx] != matrix[i][j] {
		return
	}
	if direction == RIGHT || direction == UNINIT {
		search(RIGHT, i+1, j, word_idx+1, matrix)
	}
	if direction == LEFT || direction == UNINIT {
		search(LEFT, i-1, j, word_idx+1, matrix)
	}
	if direction == DOWN || direction == UNINIT {
		search(DOWN, i, j+1, word_idx+1, matrix)
	}
	if direction == UP || direction == UNINIT {
		search(UP, i, j-1, word_idx+1, matrix)
	}
	if direction == UPRIGHT || direction == UNINIT {
		search(UPRIGHT, i+1, j+1, word_idx+1, matrix)
	}
	if direction == UPLEFT || direction == UNINIT {
		search(UPLEFT, i-1, j+1, word_idx+1, matrix)
	}
	if direction == DOWNRIGHT || direction == UNINIT {
		search(DOWNRIGHT, i+1, j-1, word_idx+1, matrix)
	}
	if direction == DOWNLEFT || direction == UNINIT {
		search(DOWNLEFT, i-1, j-1, word_idx+1, matrix)
	}
}

/* Part 2 Search */
func x_mas(i int, j int, matrix [][]byte) int {
	if i-1 < 0 || j-1 < 0 || i+1 >= len(matrix) || j+1 >= len(matrix[0]) {
		return 0
	}
	test1 := [4]byte{
		matrix[i-1][j-1], //top left
		matrix[i+1][j-1], //top right
		matrix[i-1][j+1], //bottom left
		matrix[i+1][j+1], //bottom right
	}
	/*
		TL X TR
		X  X  X
		BL X BR
	*/
	if test1 == [4]byte{'S', 'S', 'M', 'M'} {
		return 1
	} else if test1 == [4]byte{'M', 'M', 'S', 'S'} {
		return 1
	} else if test1 == [4]byte{'S', 'M', 'S', 'M'} {
		return 1
	} else if test1 == [4]byte{'M', 'S', 'M', 'S'} {
		return 1
	}
	return 0
}

func main() {
	file, err := os.Open("day4.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var matrix [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		characters := []byte(line)
		matrix = append(matrix, characters)
	}
	part2 := 0
	SIZE := len(matrix)
	for i := range SIZE {
		for j := range SIZE {
			if matrix[i][j] == 'X' {
				search(UNINIT, i, j, 0, matrix)
			}
			if matrix[i][j] == 'A' {
				part2 += x_mas(i, j, matrix)
			}
		}
	}
	fmt.Println(count)
	fmt.Println(part2)
	return
}
