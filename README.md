# Graphs Traversal Console Visualization [![License](https://img.shields.io/dub/l/vibe-d.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/P-A-R-U-S/Golang-CQRS)](https://goreportcard.com/report/github.com/P-A-R-U-S/Go-Graphs-Traversal-Console-Visualization)

Simple function to visualize graph traversal with any algorithms like DFS, BFS, A*, Dijkstra’s or Greedy Best-First.
#

| **Depth-first search (DFS) Algorithm**    | **Breadth First Search (BFS) Algorithm**  |
| ----------------------------------------- | ----------------------------------------- |
| ![](_DFS.gif)                              | ![](_BFS.gif)                              |

Please, find usage ***'PrintMatrix'*** function below. For every iteration you pass

| **PARAMETER** | **DATA TYPE** | **DESCRIPTION**                                       |
| ------------- | ------------- | ----------------------------------------------------- |
| title         | *string*      | anything you want to print as title                   |
| maze          | *[][]bool*    | 2D array representing  graph (or maze)                |
| visited       | *[][]bool*    | 2D array representing  visited cell                   |
| stack         | *[][2]int*    | array representing cell in stack or queue             |
| path          | *[][2]int*    | array representing path                               |
| interation    | *int*         | number of iterations                                  |
| startRow      | *int*         | start row                                             |
| startCol      | *int*         | start column                                          |
| exitRow       | *int*         | end row                                               |
| exitCol       | *int*         | end column                                            |

All you actually need to do, just pass all parameters above to display snapshot of current state. 
The code below are just example of usage this function to visualize BFS Shortest path algorithm.

Enjoy !!!

```Go
func findPath_BFS(maze [][]bool, startRow, startCol, exitRow, exitCol int) [][2]int {
	var interation int
	rand.Seed(time.Now().UnixNano())

	var path [][2]int

	ROWS, COLS := len(maze), len(maze[0])

	directions := [][2]int { {-1, 0}, {0, -1}, {0, 1}, {1,0} }

	q := [][2]int{ { startRow, startCol} }
	visited 	:= make([][]bool, ROWS)
	prev 		:= make([][][2]int, ROWS)
	for row := 0; row < ROWS; row++ {
		visited[row] = make([]bool, COLS)
		prev[row] = make([][2]int, COLS)
	}
	visited[startRow][startCol] = true
	prev[startRow][startCol] = [2]int{-1, -1}

	for len(q) > 0 {
		cell := q[0]
		q = q[1:]

		if cell[0] == exitRow && cell[1] == exitCol {
			row := cell[0]
			col := cell[1]
			for prev[row][col] != [2]int{-1,-1} {
				path = append(path,prev[row][col])
				row, col = prev[row][col][0], prev[row][col][1]
			}
			break
		}

		// Shuffle the directions
		rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

		for _, d := range directions {
			row, col := cell[0] + d[0], cell[1] + d[1]

			if row < 0 || row >= ROWS || col < 0 || col >= COLS { continue }

			if !maze[row][col] { continue }

			if !visited[row][col] {
				visited[row][col] = true
				q = append(q, [2]int{row, col})
				prev[row][col] = [2]int{cell[0],cell[1]}
			}
		}

		interation++
		<b>PrintMatrix("Breadth First Search (BFS)", maze, visited, q, path, interation, startRow, startCol, exitRow, exitCol)<b>
	}

	return path
}
```