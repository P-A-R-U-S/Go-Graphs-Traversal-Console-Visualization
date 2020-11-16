package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func PrintMatrix(title string, matrix [][]bool, visited [][]bool, stack [][2]int, path [][2]int, iteration int, startRow, startCol, exitRow, exitCol int) {
	// Reset to default
	const RESET = "\u001b[0m"

	// Symbols Colors
	const (
		RED     = "\u001b[31m"
		GREEN   = "\u001b[32m"
		BLACK   = "\u001b[30m"
		YELLOW  = "\u001b[33m"
		BLUE    = "\u001b[34m"
		MAGENTA = "\u001b[35m"
		CYAN    = "\u001b[36m"
		WHITE   = "\u001b[37m"

		BRIGHT_BLACK   = "\u001b[30;1m"
		BRIGHT_RED     = "\u001b[31;1m"
		BRIGHT_GREEN   = "\u001b[32;1m"
		BRIGHT_YELLOW  = "\u001b[33;1m"
		BRIGHT_BLUE    = "\u001b[34;1m"
		BRIGHT_MAGENTA = "\u001b[35;1m"
		BRIGHT_CYAN    = "\u001b[36;1m"
		BRIGHT_WHITE   = "\u001b[37;1m"
	)

	// Background Color
	const (
		BACKGROUND_BLACK   = "\u001b[40m"
		BACKGROUND_RED     = "\u001b[41m"
		BACKGROUND_GREEN   = "\u001b[42m"
		BACKGROUND_YELLOW  = "\u001b[43m"
		BACKGROUND_BLUE    = "\u001b[44m"
		BACKGROUND_MAGENTA = "\u001b[45m"
		BACKGROUND_CYAN    = "\u001b[46m"
		BACKGROUND_WHITE   = "\u001b[47m"

		BACKGROUND_BRIGHT_BLACK   = "\u001b[40;1m"
		BACKGROUND_BRIGHT_RED     = "\u001b[41;1m"
		BACKGROUND_BRIGHT_GREEN   = "\u001b[42;1m"
		BACKGROUND_BRIGHT_YELLOW  = "\u001b[43;1m"
		BACKGROUND_BRIGHT_BLUE    = "\u001b[44;1m"
		BACKGROUND_BRIGHT_MAGENTA = "\u001b[45;1m"
		BACKGROUND_BRIGHT_CYAN    = "\u001b[46;1m"
		BACKGROUND_BRIGHT_WHITE   = "\u001b[47;1m"
	)

	// Symbols
	const (
		START    = "*"
		END      = "X"
		VISITED  = "◼"
		PATH     = "◼"
		OBSTACLE = " "
		ENEMY    = ""
		SPACE    = " "
	)

	ROWS := len(visited)
	COLS := len(visited[0])

	// Clean
	fmt.Print("\u001b[1000D")                         //Move left
	fmt.Print("\u001b[" + strconv.Itoa(ROWS+8) + "A") // Move up

	fmt.Println(SPACE, SPACE, title)
	fmt.Println(SPACE, SPACE, strings.Repeat("===", COLS))

	fmt.Print(SPACE, SPACE, SPACE, SPACE, SPACE)
	for col := 0; col < COLS; col++ {
		fmt.Printf(" %2d", col)
	}
	fmt.Println()
	fmt.Print(SPACE, SPACE, SPACE, SPACE, SPACE)
	for col := 0; col < COLS; col++ {
		fmt.Print(SPACE, SPACE, "-")
	}
	fmt.Println()

	for row := 0; row < ROWS; row++ {
		fmt.Print(RESET, fmt.Sprintf("%4d|", row))
		for col := 0; col < COLS; col++ {
			fmt.Print(RESET)

			if !matrix[row][col] {
				fmt.Print(BACKGROUND_WHITE, SPACE, OBSTACLE, SPACE)
				continue
			}

			isInStack := false
			for i := range stack {
				cell := stack[i]
				if cell[0] == row && cell[1] == col {
					isInStack = true
				}
			}
			if isInStack {
				fmt.Print(BACKGROUND_YELLOW)
			}

			if row == startRow && col == startCol {
				fmt.Print(SPACE, START, SPACE) // 😀
				continue
			}

			if row == exitRow && col == exitCol {
				fmt.Print(SPACE, END, SPACE)
				continue
			}

			if visited[row][col] {
				fmt.Print(SPACE)

				isInPath := false
				for i := range path {
					if path[i][0] == row && path[i][1] == col {
						isInPath = true
						break
					}
				}

				if isInPath {
					fmt.Print(GREEN, PATH)
				} else {
					fmt.Print(RED, VISITED)
				}
				fmt.Print(SPACE)
				continue
			}

			fmt.Print(SPACE, SPACE, SPACE)
		}
		fmt.Println(RESET)
	}
	fmt.Println(RESET)
	fmt.Print(SPACE, SPACE)
	fmt.Printf("ITERATION: %4d", iteration)
	fmt.Println()

	//fmt.Println(Background_Yellow, " ", RESET, " - In Stack")
	//fmt.Println(RED, VISITED, RESET, " - Visited")
	//fmt.Println(Background_Bright_White, OBSTACLE, RESET, " - Obstacle")
	//fmt.Println(GREEN, VISITED, RESET, " - Path")

	time.Sleep(time.Millisecond * 10)
}
