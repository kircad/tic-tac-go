// Package main defines the executable entry point.
package main

import (
	"errors"
	"flag" // Standard CLI flag parsing.
	"fmt"  // Formatted I/O for printing output.
	ticatacgo "game-runner/tictacgo"
	"game-runner/types"
	"os" // Access to process exit.
)

// TODO refactor to support multiple game types -- need to increase abstraction significantly
// E.G. game interface should include move, move legality checker, win checker, player list, etc.

const version = "0.1.0" // CLI version string.

type game interface {
	GetStarterPlayer() (bool, bool)
	Move(bool) bool
	CheckEnd(bool) string // TODO ENUM TYPE?
	HandleOutcome(string)
}

func main() {
	cfg := types.Config{}                                                    // Initialize config with zero values.
	showVersion := flag.Bool("version", false, "print the version and exit") // Define a -version flag.
	flag.StringVar(&cfg.Difficulty, "difficulty", "easy", "CPU skill")       // Bind -difficulty to cfg.difficulty.
	flag.StringVar(&cfg.Game, "game", "tic-tac-go", "Game selection")
	flag.IntVar(&cfg.Rows, "rows", 3, "Game selection")
	flag.IntVar(&cfg.Cols, "cols", 3, "Game selection")
	flag.IntVar(&cfg.WinLen, "winLen", 3, "Game selection")
	flag.StringVar(&cfg.P1, "p1", "Player", "Player 1 name")
	flag.StringVar(&cfg.P2, "p2", "CPU", "Player 2 name")
	flag.Parse() // Parse flags from os.Args.

	if *showVersion { // If -version was provided... (checks if not nil)
		fmt.Println(version) // Print version.
		os.Exit(0)           // Exit without doing normal work.
	}

	var err error
	var g game
	if cfg.Game == "tic-tac-go" {
		g, err = ticatacgo.GetGame(&cfg)
	} else {
		err = errors.New("Invalid game selection.")
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if cfg.P2 != "CPU" {
		fmt.Println("Playing Multiplayer " + cfg.Game + ".")
	} else {
		fmt.Println("Playing " + cfg.Game + " on " + cfg.Difficulty + " difficulty.")
	}
	run(g)
	fmt.Println("Thank you for playing " + cfg.Game + "!")
}

func run(g game) {
	var outcome string
	fmt.Println("Enter '!q' at any time to end the game")
	fmt.Println()

	p1turn, quit := g.GetStarterPlayer()
	if quit {
		return
	}

	for {
		if quit := g.Move(p1turn); quit {
			break
		}
		outcome = g.CheckEnd(p1turn)
		if outcome != "" {
			break
		}
		p1turn = !p1turn

	}
	g.HandleOutcome(outcome)
}
