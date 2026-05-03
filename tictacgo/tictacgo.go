package ticatacgo

import (
	"errors"
	"fmt"
	"game-runner/helpers"
	"game-runner/types"
	"math/rand"
)

type Game struct {
	gameBoard   types.Board
	turnCounter types.Counter
	players     [2]types.Player
}

func GetGame(cfg *types.Config) (Game, error) {
	gameBoard := helpers.ConstructBoard(cfg.Rows, cfg.Cols)
	// TODO MODIFY CODE TO ALLOW FOR VARIABLE WINLEN, ROWS, COLS
	var players [2]types.Player
	players[0].Name = cfg.P1
	players[0].Icon = 'O'
	players[0].Algorithm = helpers.GetUserInput

	players[1].Name = cfg.P2
	players[1].Icon = 'X'
	if cfg.P2 == "CPU" {
		switch cfg.Difficulty {
		case "easy":
			players[1].Algorithm = randomMoves
		case "medium":
			players[1].Algorithm = mixedMoves
		case "hard":
			players[1].Algorithm = optimalMoves
		default:
			return Game{}, errors.New("Difficulty must be one of [easy, medium, hard]")
		}
	} else {
		players[1].Algorithm = helpers.GetUserInput
	}

	g := Game{
		gameBoard:   gameBoard,
		players:     players,
		turnCounter: helpers.CheckTie(len(gameBoard) * len(gameBoard[0])),
	}
	return g, nil
}

func (g Game) GetStarterPlayer() (bool, bool) {
	var choice string
	var choiceInt int32
	fmt.Println("Choosing starter player via coin toss. Heads or tails? Enter !q to quit")
	fmt.Print("Your choice: ")
	for {
		fmt.Scanln(&choice)
		if choice == "!q" {
			return false, true
		}
		if choice == "heads" {
			choiceInt = 0
			break
		}
		if choice == "tails" {
			choiceInt = 1
			break
		}

		fmt.Print("Choice must be within ['heads','tails']: ")
	}

	num := rand.Int31n(2)
	if num == choiceInt {
		fmt.Println(g.players[0].Name + " starts!")
		return true, false
	} else {
		fmt.Println(g.players[1].Name + " starts!")
		return false, false
	}
}

func (g Game) Move(p1Turn bool) (bool, types.Coords) {
	var playerIdx int
	if !p1Turn {
		playerIdx = 1
	}
	fmt.Println(g.players[playerIdx].Name + " Move:")
	helpers.PrintBoard(g.gameBoard, false)
	fmt.Println()

	quitCoords := types.Coords{-1, -1}
	selection := g.players[playerIdx].Algorithm(g.gameBoard)
	if selection == quitCoords {
		return true, types.Coords{}
	}
	g.gameBoard[selection[0]][selection[1]] = g.players[playerIdx].Icon
	return false, selection
}

func (g Game) CheckEnd(p1Turn bool, coords types.Coords) string {
	var playerIdx int
	if !p1Turn {
		playerIdx = 1
	}
	if win := helpers.CheckWin(g.gameBoard, coords, 3); win {
		return g.players[playerIdx].Name
	}
	if end := g.turnCounter(); end {
		return "tie"
	}
	return ""
}

func (g Game) HandleOutcome(outcome string) {
	if !(outcome == "") {
		helpers.PrintBoard(g.gameBoard, false)
		if outcome == "tie" {
			fmt.Println("Tie")
		} else {
			fmt.Println(outcome + " Won!")
		}
	}
}

func randomMoves(gameBoard types.Board) types.Coords {
	for {
		selection := rand.Int31n(9)
		cpuRow := selection / 3 // TODO FIX THIS
		cpuCol := selection % 3
		if !(gameBoard[cpuRow][cpuCol] == '_') {
			continue
		}
		return types.Coords{cpuRow, cpuCol}
	}
}

func mixedMoves(board types.Board) types.Coords {
	// TODO IMPLEMENT -- Play optimally 50% of time OR decise suboptimal algorithm
	return types.Coords{}
}

func optimalMoves(board types.Board) types.Coords {
	// TODO IMPLEMENT -- Should play optimally according to some algorithm -- try to derive!
	return types.Coords{}
}
