# Go WebAssembly Tic-Tac-Toe

This is a simple Tic-Tac-Toe game built with Go and WebAssembly.

## Prerequisites

- Go 1.22.3 or later
- A modern web browser with WebAssembly support

## Building the Project

1. Build the WebAssembly file:

   ```bash
   GOOS=js GOARCH=wasm go build -o main.wasm
   ```

2. Copy the WebAssembly JavaScript support file:

   ```bash
   cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
   ```

## Running the Application

1. Start a local web server. For example, using Python's built-in server:

   ```cmd
   python -m http.server 8080
   ```

   or

   ```cmd
   go run server.go
   ```

2. Open your browser and navigate to:

   ``` link
   http://localhost:8080
   ```

## Project Structure

- `main.go` - Main Go code compiled to WebAssembly
- `index.html` - Web interface for the game
- `main.wasm` - Compiled WebAssembly binary (generated after build)
- `wasm_exec.js` - Go's WebAssembly runtime (copied during setup)

## Development

To test changes, rebuild the WebAssembly file and refresh your browser:

```bash
go build -o main.wasm
```
