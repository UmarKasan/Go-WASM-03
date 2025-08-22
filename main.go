//go:build js && wasm

package main

import "syscall/js"

var (
	board       [9]string
	gameOver    bool
	playerState = "X"
	player1     = "X"
	player2     = "O"
	status      js.Value
	// board     js.Value
)

func main() {
	document := js.Global().Get("document")
	status := document.Call("querySelector", "#status")
	status.Set("innerHTML", "Player "+playerState+"'s turn")
	// Test script
	// cell := document.Call("querySelector", ".cell[data-index=\"0\"]")
	// if !cell.IsUndefined() {
	// 	cell.Set("innerHTML", "Hi")
	// }

	// cell := document.Call("querySelector", ".cell[data-index=\"0\"]")

	// Handle all 9 cells
	for i := 0; i < 9; i++ {
		cell := document.Call("querySelector", ".cell[data-index=\""+string(rune(i+'0'))+"\"]")
		// Click Event
		onClick := js.FuncOf(func(this js.Value, args []js.Value) any {
			cell.Set("innerHTML", playerState)
			// Change Player turn
			if playerState == player1 {
				playerState = player2
			} else {
				playerState = player1
			}
			return nil
		})
		cell.Call("addEventListener", "click", onClick)
	}
	// Keep running
	select {}
}
