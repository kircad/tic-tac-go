package helpers

import (
	"fmt"
	"game-runner/types"
	"strconv"
	"unicode/utf8"
)

// Board Utilities

var gameBoard types.Board = [][]rune{
	{'_', '_', '_'},
	{'_', '_', '_'},
	{'_', '_', '_'},
}

func ConstructBoard(rows int, cols int) types.Board {
	var board types.Board
	for range rows {
		row := make([]rune, cols)
		for j := range row {
			row[j] = '_'
		}
		board = append(board, row)
	}
	return board
}

func PrintBoard(board types.Board, playerMode bool) {
	// TODO MAKE PRETTIER
	for i := range board {
		for j := range board[i] {
			if playerMode {
				fmt.Print(strconv.Itoa(i*len(board)+j+1) + " ")
			} else {
				fmt.Print(string(board[i][j]) + " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Move utilities

func GetUserInput(gameBoard types.Board) types.Coords {
	var userInput string
	for {
		fmt.Println("Press a key (1-9) corresponding to a position on the board:")
		PrintBoard(gameBoard, true)
		fmt.Print("Your Input: ")
		fmt.Scanln(&userInput)

		if userInput == "!q" {
			return types.Coords{-1, -1} // TODO is there a better way to signal a quit
		}

		if utf8.RuneCountInString(userInput) > 1 {
			fmt.Println("Too many characters in input!")
			continue
		}

		val, _ := strconv.Atoi(userInput)
		if val < 1 || val > 9 {
			fmt.Println("Input must be between 1 and 9!")
			continue
		}
		var selection types.Coords = types.Coords{int32((val - 1)) / 3, int32((val - 1) % 3)}

		if !(gameBoard[selection[0]][selection[1]] == '_') {
			fmt.Println("Board occupied!")
			continue
		}
		return selection
	}
}

// End of game utilities

func CheckTie(lim int) func() bool {
	moves := 0
	return func() bool {
		moves += 1
		if moves == lim {
			return true
		}
		return false
	}
}

func CheckWin(gameBoard types.Board, coords types.Coords, winLen int) bool { // TODO clean up
	// win sequence must contain placed piece -- so starting at piece, extend in either direction until you reach an edge or non-player piece, then go in other direction. If win, return true
	height, width := len(gameBoard[0]), len(gameBoard)

	// check row
	fmt.Println("CHECKING ROW")
	winPiece := gameBoard[coords[0]][coords[1]]
	idx := coords[1] + 1
	streak := 1
	for idx < int32(width) {
		if gameBoard[coords[0]][idx] == winPiece {
			streak += 1
			if streak == winLen {
				return true
			}
		} else {
			break
		}
		idx += 1
	}

	idx = coords[1] - 1
	for idx >= 0 {
		if gameBoard[coords[0]][idx] == winPiece {
			streak += 1
			if streak == winLen {
				return true
			}
		} else {
			break
		}
		idx -= 1
	}

	// check col
	streak = 1
	idx = coords[0] + 1
	for idx < int32(height) {
		if gameBoard[idx][coords[1]] == winPiece {
			streak += 1
			if streak == winLen {
				return true
			}
		} else {
			break
		}
		idx += 1
	}

	idx = coords[0] - 1
	for idx >= 0 {
		if gameBoard[idx][coords[1]] == winPiece {
			streak += 1
			if streak == winLen {
				return true
			}
		} else {
			break
		}
		idx -= 1
	}

	// check left diagonal
	streak = 1
	row, col := coords[0]-1, coords[1]-1
	for row >= 0 && col >= 0 {
		if gameBoard[row][col] == winPiece {
			streak += 1
			if streak == winLen {
				return true
			}
		} else {
			break
		}
		col -= 1
		row -= 1
	}

	row, col = coords[0]+1, coords[1]+1
	for col < int32(width) && row < int32(height) {
		if gameBoard[row][col] == winPiece {
			streak += 1
			if streak == winLen {
				return true
			}
		} else {
			break
		}
		col += 1
		row += 1
	}

	// check right diagonal
	streak = 1
	row, col = coords[0]-1, coords[1]+1
	for col < int32(width) && row >= 0 {
		if gameBoard[row][col] == winPiece {
			streak += 1
			if streak == winLen {
				return true
			}
		} else {
			break
		}
		col += 1
		row -= 1
	}

	row, col = coords[0]+1, coords[1]-1
	for col >= 0 && row < int32(height) {
		if gameBoard[row][col] == winPiece {
			streak += 1
			if streak == winLen {
				return true
			}
		} else {
			break
		}
		col -= 1
		row += 1
	}

	return false
}

// Old implementation -- does not use position of most recently placed
// func CheckWin(gameBoard types.Board, winLen int) bool {
// 	width, height := len(gameBoard[0]), len(gameBoard)
// 	// Check rows
// 	for i := range height {
// 		if win := checkSequence(gameBoard[i], winLen); win {
// 			return true
// 		}
// 	}

// 	// Check cols
// 	var col []rune = make([]rune, height)
// 	for i := range width {
// 		for j := range len(gameBoard[i]) {
// 			col[j] = gameBoard[j][i]
// 		} // IMPORTANT x[:][0] == x[0] in Go! [:] just slices the entire outer slice, then [i] takes the ith row -- this is why we have to manually construct col
// 		if win := checkSequence(col, winLen); win {
// 			return true
// 		}
// 	}

// 	// Check diagonals
// 	var diag []rune
// 	maxDiag := min(width, height)
// 	for i := range height {
// 		topCoords, bottomCoords := &types.Coords{0, int32(i)}, &types.Coords{int32(height) - 1, int32(i)}

// 		if i-winLen+1 >= 0 { // Assumes height >= winLen

// 			// top row
// 			diag = constructDiagonal(gameBoard,
// 				topCoords,
// 				func(coords *types.Coords) {
// 					coords[0]++
// 					coords[1]--
// 				},
// 				func(coords types.Coords) bool {
// 					return coords[0] == int32(height) || coords[1] < 0
// 				},
// 				maxDiag)
// 			if win := checkSequence(diag, winLen); win {
// 				return true
// 			}

// 			// bottom row
// 			diag = constructDiagonal(gameBoard,
// 				bottomCoords,
// 				func(coords *types.Coords) {
// 					coords[0]--
// 					coords[1]--
// 				},
// 				func(coords types.Coords) bool {
// 					return coords[0] < 0 || coords[1] < 0
// 				},
// 				maxDiag)
// 			if win := checkSequence(diag, winLen); win {
// 				return true
// 			}

// 		}

// 		if (i+winLen <= width) && (width-i < height) { // Assumes height >= winLen

// 			// top row
// 			diag = constructDiagonal(gameBoard,
// 				topCoords,
// 				func(coords *types.Coords) {
// 					coords[0]++
// 					coords[1]++
// 				},
// 				func(coords types.Coords) bool {
// 					return coords[0] == int32(height) || coords[1] == int32(height)
// 				},
// 				maxDiag)
// 			if win := checkSequence(diag, winLen); win {
// 				return true
// 			}

// 			// bottom row
// 			diag = constructDiagonal(gameBoard,
// 				bottomCoords,
// 				func(coords *types.Coords) {
// 					coords[0]--
// 					coords[1]++
// 				},
// 				func(coords types.Coords) bool {
// 					return coords[0] < 0 || coords[1] < 0
// 				},
// 				maxDiag)
// 			if win := checkSequence(diag, winLen); win {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func constructDiagonal(board types.Board, coords *types.Coords, increment func(*types.Coords), checkEnd func(types.Coords) bool, maxLen int) []rune {
// 	diag := make([]rune, 0, maxLen)
// 	for {
// 		diag = append(diag, board[coords[0]][coords[1]])
// 		increment(coords)
// 		if checkEnd(*coords) {
// 			return diag
// 		}
// 	}
// }

// func checkSequence(sequence []rune, winLen int) bool {
// 	var streak int
// 	var currStart int
// 	for i := range sequence {
// 		if sequence[i] == '_' {
// 			continue
// 		}
// 		if sequence[i] == sequence[currStart] {
// 			streak++
// 			if streak == winLen {
// 				return true
// 			}
// 		} else {
// 			currStart = i
// 			streak = 1
// 		}
// 	}
// 	return false
// }
