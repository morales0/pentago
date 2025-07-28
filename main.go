package main

import (
	// "errors"
	"fmt"
	"log"
)

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

type Board struct {
	tl [][]rune
	tr [][]rune
	br [][]rune
	bl [][]rune
}

type cell struct {
	color string
}

func printBoard(board [6][6]rune) (string, error) {
	str := ""
	charMap := map[rune]string{
		'0': "○",
		'w': green + "●" + reset,
		'b': yellow + "●" + reset,
	}
	for i, row := range board {
		for j, c := range row {
			char, ok := charMap[c]
			if !ok {
				return "", fmt.Errorf("invalid board character '%v'", c)
			}
			str += string(char) + " "

			if j == 2 {
				str += "| "
			}
		}
		str += "\n"
		if i == 2 {
			str += "-------------\n"
		}
	}

	return str, nil
}

func playMove(board [6][6]rune, player rune, i, j int) (bool, error) {
	board[i][j] = player
	return true, nil
}

func handleCommand(command string, params ...string) {
	fmt.Println("You typed in command " + command)
}

func startGameLoop(board [6][6]rune) {
	playing := true
	var command string
	for playing {
		boardString, err := printBoard(board)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\n" + boardString)
		fmt.Println("\n\n\n" + "Command (h or help) > ")
		fmt.Scan(&command)
		handleCommand(command)
	}
}

func main() {
	var grid [6][6]rune
	for i, row := range grid {
		for j, _ := range row {
			grid[i][j] = '0'
		}
	}
	startGameLoop(grid)
	// fmt.Println("○ ● ○ | ○ ○ ○")
	// fmt.Println("○ ● ○ | ○ ○ ○")
	// fmt.Println("○ ● ○ | ○ ○ ○")
	// fmt.Println("-------------")
	// fmt.Println("○ ● ○ | ○ ○ ○")
	// fmt.Println("○ ● ○ | ○ ○ ○")
	// fmt.Println("○ ● ○ | ○ ○ ○")
}
