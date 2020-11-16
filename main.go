package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	testCases := []struct {
		maze     [][]bool
		startRow int
		startCol int
		exitRow  int
		exitCol  int
		result   [][2]int
	}{
		{
			[][]bool{
				//0     1      2      3      4      5      6      7      8      9      10     11     12     13     14     15     16     17    18     19     20     21     22     23     24     25     26     27     28
				{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true},                   // 0
				{true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, true, false, false, false},  // 1
				{true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, true},             // 2
				{true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, true, false, true, true, true},             // 3
				{true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true},      // 4
				{false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true},         // 5
				{true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true},         // 6
				{true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, false, true, true, true, false, false, true, false, true, true, true},         // 7
				{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, false, true, true, true, false, true, true, false, true, true, true},            // 8
				{true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true, true, true, true, true, true, true, true},                 // 9
				{true, true, true, false, false, true, false, false, false, false, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, true, false, false, false}, // 10
				{true, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, true, false, true, false, true, true, true, true, true},            // 11
				{true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, true, true, true, true, true, false, true, true, true, false, true, true, true},             // 12
				{true, false, true, true, false, false, true, false, true, true, false, true, false, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true},     // 13
				{false, false, true, true, true, true, true, false, true, false, false, true, false, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true},        // 14
				{true, true, true, true, true, false, false, false, true, true, false, true, false, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true},       // 15
				{true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true},          // 16
				{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true},             // 17
				{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true},             // 18
				{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false, true, true, true},                  // 19
				{true, true, true, false, false, true, false, false, false, true, true, true, false, false, true, false, false, false, false, false, true, true, true, false, false, false, false, false, false}, // 20
				{true, false, true, false, true, true, true, true, true, true, false, true, false, true, true, true, false, true, true, true, true, false, true, false, true, false, true, true, true},           // 21
				{true, false, true, true, true, false, true, true, true, true, false, true, true, true, false, true, false, true, true, true, true, false, true, true, true, false, true, true, true},            // 22
				{true, false, true, true, false, false, true, false, true, true, false, true, true, false, false, true, false, true, false, true, true, false, true, true, false, false, true, false, true},      // 23
				{false, false, true, true, true, true, true, false, true, false, false, true, true, true, true, true, false, true, false, true, false, false, true, true, true, true, true, false, true},         // 24
				{true, true, true, true, true, false, false, false, true, true, true, true, true, true, false, false, false, true, false, true, true, true, true, true, true, false, false, false, true},         // 25
				{true, true, false, false, true, false, true, true, true, true, true, false, false, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true},          // 26
				{true, true, false, true, true, false, true, true, true, true, true, false, true, true, false, true, true, true, true, true, true, true, false, true, true, false, true, true, true},             // 27

			},
			0, 0,
			27, 28,
			[][2]int{},
		},

		{
			[][]bool{
				//0     1      2      3      4
				{true, true, true, true, true},   // 0
				{true, true, true, false, false}, // 1
				{true, false, true, false, true}, // 2
				{true, false, true, true, true},  // 3
				{true, false, true, true, false}, // 4
				{false, false, true, true, true}, // 5
			},
			0, 0,
			5, 4,
			[][2]int{},
		},
	}

	for _, tc := range testCases {

		// DFS
		ClearConsole()
		fmt.Println("--------------------------------------------------------")
		fmt.Println("Death-First-Search")
		fmt.Println("--------------------------------------------------------")
		ClearConsole()
		_ = findPath_DFS(tc.maze, tc.startRow, tc.startCol, tc.exitRow, tc.exitCol)
		PressToContinue()
	}
}

func ClearConsole() {
	fmt.Print("\u001b[H\u001b[2J")
}
func PressToContinue() {
	reader := bufio.NewReader(os.Stdin)
	_, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
}
func Sleep(ms time.Duration) {
	time.Sleep(time.Millisecond * ms)
}
