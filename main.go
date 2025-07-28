package main

import (
	// "errors"
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board     [6][6]rune
	cursor    [2]int
	whiteTurn bool
	step      int
	playerWon rune
}

func initialModel() model {
	var board [6][6]rune
	for i, row := range board {
		for j, _ := range row {
			board[i][j] = '0'
		}
	}

	return model{
		board:     board,
		cursor:    [2]int{0, 0},
		whiteTurn: true,
		step:      1,
		playerWon: '0',
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		case "y":
			if m.playerWon != '0' {
				m = initialModel()
				break
			}
		case "n":
			if m.playerWon != '0' {
				return m, tea.Quit
			}

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor[1] > 0 {
				m.cursor[1]--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			var limit int
			if m.step == 2 {
				limit = 1
			} else {
				limit = 5
			}
			if m.cursor[1] < limit {
				m.cursor[1]++
			}

		case "left", "h":
			if m.cursor[0] > 0 {
				m.cursor[0]--
			}

		case "right", "l":
			var limit int
			if m.step == 2 {
				limit = 1
			} else {
				limit = 5
			}

			if m.cursor[0] < limit {
				m.cursor[0]++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":

			if m.step == 2 {
				break
			}

			if m.board[m.cursor[1]][m.cursor[0]] == '0' {
				if m.whiteTurn {
					m.board[m.cursor[1]][m.cursor[0]] = 'w'
				} else {
					m.board[m.cursor[1]][m.cursor[0]] = 'b'
				}

				if m.step == 1 {
					m.step++
					m.cursor[0] = 0
					m.cursor[1] = 0
				} else {
					m.whiteTurn = !m.whiteTurn
					m.step = 1
				}
			}
		case "e":
			if m.step != 2 {
				break
			}
			rotateLeft(&m)
			// Check for winning board
			win := checkWin(&m)
			if win {
				if m.whiteTurn {
					m.playerWon = 'w'
				} else {
					m.playerWon = 'b'
				}
			} else {
				m.whiteTurn = !m.whiteTurn
			}
			m.step = 1
			m.cursor[0] = 0
			m.cursor[1] = 0

		case "r":
			if m.step != 2 {
				break
			}
			rotateRight(&m)
			// Check for winning board
			win := checkWin(&m)
			if win {
				if m.whiteTurn {
					m.playerWon = 'w'
				} else {
					m.playerWon = 'b'
				}
			} else {
				m.whiteTurn = !m.whiteTurn
			}
			m.step = 1
			m.cursor[0] = 0
			m.cursor[1] = 0

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	s := "Penta-Go!"

	boardString, err := getBoardString(&m)
	if err != nil {
		log.Fatal(err)
	}

	s += "\n\n" + boardString

	var player string
	if m.whiteTurn {
		player = blue + "Player 1" + reset
	} else {
		player = yellow + "Player 2" + reset
	}

	if m.playerWon != '0' {
		s += "Congrats " + player + "!\n"
		s += "Play again? (y/n)"
		return s
	}

	var step string
	if m.step == 1 {
		step = "Choose a space to play (hjkl, enter)"
	} else {
		step = "Choose a quadrant then type direction (hjkl, r for right, e for left)"
	}
	s += "\n" + player + ": " + step

	return s
}

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
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

func getBoardString(m *model) (string, error) {
	str := ""
	charMap := map[rune]string{
		'0': "○",
		'w': blue + "●" + reset,
		'b': yellow + "●" + reset,
	}
	var color string
	if m.whiteTurn {
		color = blue
	} else {
		color = yellow
	}

	for i, row := range m.board {
		for j, c := range row {
			char, ok := charMap[c]
			if !ok {
				return "", fmt.Errorf("invalid board character '%v'", c)
			}
			if m.step == 1 && m.cursor[0] == j && m.cursor[1] == i {
				if m.board[m.cursor[1]][m.cursor[0]] != '0' {
					str += "● "
				} else {
					str += color + "● " + reset
				}
			} else {
				str += string(char) + " "
			}

			if j == 2 {
				if m.step == 2 && m.cursor[1] == i/3 {
					str += color + "| " + reset
				} else {
					str += "| "
				}
			}
		}
		str += "\n"
		if i == 2 {
			if m.step == 2 {
				if m.cursor[0] == 0 {
					str += color + "------" + reset + "-------\n"

				} else {
					str += "-------" + color + "------" + reset + "\n"
				}
			} else {
				str += "-------------\n"
			}
		}
	}

	return str, nil
}

func rotateLeft(m *model) {
	var newBoard [6][6]rune
	for i, row := range newBoard {
		for j, _ := range row {
			if m.cursor[1] == i/3 && m.cursor[0] == j/3 {
				newBoard[i][j] = m.board[j+3*m.cursor[1]-3*m.cursor[0]][(2+3*m.cursor[0])-(i-(3*m.cursor[1]))]

			} else {
				newBoard[i][j] = m.board[i][j]
			}
		}
	}

	m.board = newBoard

}

func rotateRight(m *model) {
	var newBoard [6][6]rune
	for i, row := range newBoard {
		for j, _ := range row {
			if m.cursor[1] == i/3 && m.cursor[0] == j/3 {
				newBoard[i][j] = m.board[(2+3*m.cursor[1])-(j-3*m.cursor[0])][i+3*m.cursor[0]-3*m.cursor[1]]
			} else {
				newBoard[i][j] = m.board[i][j]
			}
		}
	}

	m.board = newBoard
}
func checkWin(m *model) bool {
	rows, cols := 6, 6
	var currentPiece rune
	if m.whiteTurn {
		currentPiece = 'w'
	} else {
		currentPiece = 'b'
	}

	// 8 directions: (rowDelta, colDelta)
	directions := [][2]int{
		{0, 1},  // right
		{1, 0},  // down
		{1, 1},  // down-right
		{1, -1}, // down-left
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if m.board[r][c] != currentPiece {
				continue
			}

			for _, dir := range directions {
				count := 1
				nr, nc := r+dir[0], c+dir[1]
				for count < 5 && nr >= 0 && nr < rows && nc >= 0 && nc < cols && m.board[nr][nc] == currentPiece {
					count++
					nr += dir[0]
					nc += dir[1]
				}
				if count == 5 {
					return true
				}
			}
		}
	}
	return false
}

func playMove(board [6][6]rune, player rune, i, j int) (bool, error) {
	board[i][j] = player
	return true, nil
}

func handleCommand(command string, params ...string) {
	fmt.Println("You typed in command " + command)
}
func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
