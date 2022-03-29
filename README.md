## Axii

A simple chess engine written in go

### Building & running

Go version 1.17 is required.

To build the app, run:

```
go build .
```

This will generate an executable that you can run:

```
./axii
```

### Docker

This app is also dockerized. To build:

```
docker build . -t axii
```

And to run:

```
docker run -it axii
```

### Usage

The engine accepts the most common UCI commands.

To ask the engine what it thinks is the best move in the current position, enter `go`, and wait for the calculation to finish.

You can then execute the move on the board, with the `move` command, for example `move e2e4`.

The `showpos` will print out the current board position to the console.

To play against the engine via the terminal, you can enter your moves via `move` commands and run `go` whenever it's the engine's turn, and then execute its moves similarly.

To exit the engine, type either `exit` or `quit`.
