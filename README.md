# go-tic-tac-toe

A command-line tic-tac-toe game written in Go.

This is a personal learning project — written by hand to practice Go, without AI assistance.

## Usage

```
go run . [flags]
```

**Flags:**

| Flag | Default | Description |
|------|---------|-------------|
| `-difficulty` | `easy` | CPU skill level |
| `-rows` | `3` | Board rows |
| `-cols` | `3` | Board columns |
| `-winLen` | `3` | Consecutive pieces needed to win |
| `-p1` | `Player` | Player 1 name |
| `-p2` | `CPU` | Player 2 name (set to anything other than `CPU` for multiplayer) |
| `-version` | | Print version and exit |

**Examples:**

```sh
# vs CPU
go run .

# multiplayer
go run . -p1 Alice -p2 Bob

# larger board
go run . -rows 5 -cols 5 -winLen 4
```

## Project structure

```
main.go          # Entry point, CLI flags, game loop
tictacgo/        # Game logic
types/           # Shared types and config
helpers/         # Utility functions
```
