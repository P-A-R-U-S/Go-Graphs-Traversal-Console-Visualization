package main

import (
	"math/rand"
	"time"
)

func findPath_BFS(maze [][]bool, startRow, startCol, exitRow, exitCol int) [][2]int {
	var interation int
	rand.Seed(time.Now().UnixNano())

	var path [][2]int

	ROWS := len(maze)
	COLS := len(maze[0])

	directions := [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

	q := [][2]int{{startRow, startCol}}
	visited := make([][]bool, ROWS)
	prev := make([][][2]int, ROWS)
	for row := 0; row < ROWS; row++ {
		visited[row] = make([]bool, COLS)
		prev[row] = make([][2]int, COLS)
	}
	visited[startRow][startCol] = true
	prev[startRow][startCol] = [2]int{-1, -1}

	for len(q) > 0 {
		cell := q[0]
		q = q[1:]

		//
		if cell[0] == exitRow && cell[1] == exitCol {
			row := cell[0]
			col := cell[1]
			for prev[row][col] != [2]int{-1, -1} {
				path = append(path, prev[row][col])
				row, col = prev[row][col][0], prev[row][col][1]
			}
			interation++
			PrintMatrix("Breadth First Search (BFS) Algorithm | Shortest Path", maze, visited, q, path, interation, startRow, startCol, exitRow, exitCol) //PressToContinue()
			break
		}

		// Shuffle the directions
		rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

		for _, d := range directions {
			row := cell[0] + d[0]
			col := cell[1] + d[1]

			if row < 0 || row >= ROWS || col < 0 || col >= COLS {
				continue
			}

			if !maze[row][col] {
				continue
			}

			if !visited[row][col] {
				visited[row][col] = true
				q = append(q, [2]int{row, col})
				prev[row][col] = [2]int{cell[0], cell[1]}
			}
		}

		interation++
		PrintMatrix("Breadth First Search (BFS) Algorithm | Shortest Path", maze, visited, q, path, interation, startRow, startCol, exitRow, exitCol) //PressToContinue()
	}

	return path
}
