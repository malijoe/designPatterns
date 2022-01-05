package template

import "fmt"

func PlayGame(start, takeTurn func(), haveWinner func() bool, winningPlayer func() int) {
    for !haveWinner() {
        takeTurn()
    }

    fmt.Printf("Player %d wins. \n", winningPlayer())
}
func FunctionalTemplate() {
    turn, maxTurns, currentPlayer := 1, 10, 0

    start := func() {
        fmt.Println("Starting a game of chess")
    }

    takeTurn := func() {
        fmt.Printf("Turn %d taken by player %d \n", turn, currentPlayer)
        turn ++
        currentPlayer = 1 - currentPlayer
    }

    haveWinner := func() bool {
        return turn == maxTurns
    }

    winningPlayer := func() int {
        return currentPlayer
    }

    PlayGame(start, takeTurn, haveWinner, winningPlayer)
}