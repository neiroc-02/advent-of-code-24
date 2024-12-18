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
	SIZE := len(matrix)
	for i := range SIZE {
		for j := range SIZE {
			if matrix[i][j] == 'X' {
				search(UNINIT, i, j, 0, matrix)
			}
		}
	}
	fmt.Println(count)
	return
}
