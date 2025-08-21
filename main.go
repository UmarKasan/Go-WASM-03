//go:build js && wasm

package main

import "syscall/js"

var (
	board       [9]string
	gameOver    bool
	playerState string
	player1     = "X"
	player2     = "O"
	statusEl    js.Value
	boardEl     js.Value
)

func main() {
	document := js.Global().Get("document")
	cells := document.Call("querySelectorAll", ".cell")
	cells.Set("innerHTML", "Hi")
	// boardEl = document.Call("querySelector", ".board")
	// statusEl = document.Call("querySelector", "#status")
	// restartButton := document.Get("restart")
	// restartButton.Call("addEventListener", "click", restartGame)

	// Initialize game state
	// playerState = player1
	// gameOver = false
	// board = [9]string{}

	// Add click event listeners to all cells

	// for i := 0; i < 9; i++ {
	// 	cell := cells.Index(i)
	// 	index := i // Create a new variable for the closure
	// 	cell.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 		if !gameOver && board[index] == "" {
	// 			cellClick(index)
	// 		}
	// 		return nil
	// 	}))
	// }

	// Keep running
	select {}
}

// func cellClick(index int) {
// 	if gameOver || board[index] != "" {
// 		return
// 	}

// 	board[index] = playerState
// 	boardEl.Get("children").Index(index).Set("innerHTML", playerState)

// 	// Toggle player
// 	if playerState == player1 {
// 		playerState = player2
// 	} else {
// 		playerState = player1
// 	}
// 	statusEl.Set("innerHTML", "Player "+playerState+"'s turn")
// }
