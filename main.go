//go:build js && wasm

package main

import "syscall/js"

var (
	board       [9]string
	gameOver    bool
	playerState = "X"
	player1     = "X"
	player2     = "O"
	statusEl    js.Value
	boardEl     js.Value
)

func main() {
	document := js.Global().Get("document")
	// Test script
	// cell := document.Call("querySelector", ".ccell[data-index=\"0\"ell")
	// if !cell.IsUndefined() {
	// 	cell.Set("innerHTML", "Hi")
	// }
	cell := document.Call("querySelector", ".cell[data-index=\"0\"]")
	status := document.Call("querySelector", "#status")
	status.Set("innerHTML", "Player "+playerState+"'s turn")

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

	// Keep running
	select {}
}
