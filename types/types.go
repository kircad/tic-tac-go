package types

type Board [][]rune
type Coords [2]int32
type Algorithm func(Board) Coords
type Counter func() bool

// config holds parsed CLI options.
type Config struct {
	Game       string
	Difficulty string
	P1         string
	P2         string
	Rows       int
	Cols       int
	WinLen     int
}

type Player struct {
	Name      string
	Icon      rune
	Algorithm Algorithm
}
