# Tic-Tac-Toe

A command-line Tic-Tac-Toe game written in Go. Play as X against the computer and choose a difficulty:

- **Easy:** chooses a random legal move.
- **Medium:** takes a winning move, blocks an immediate human win, prefers the center and then corners.
- **Hard:** uses minimax and plays optimally.

## Run

```sh
go run xox_game.go
```

Choose a difficulty from 1 to 3, then enter a number from 1 to 9 to choose a square.
