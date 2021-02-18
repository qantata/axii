package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func setPosition(pos *Position, tokens []string) {
	if tokens[1] == "startpos" {
		pos.load(STARTING_POS)

		if len(tokens) == 2 {
			return
		}
	}

	if tokens[2] == "moves" {
		for i := 3; i < len(tokens); i++ {
			handleUciMove(pos, tokens[i])
		}
	} else {
		if len(tokens) < 6 {
			fmt.Println("Fen position string has to have at least 6 space separated parts")
			return
		}

		if tokens[1] != "fen" {
			fmt.Println("Position string is not fen!")
			return
		}

		fen := ""
		for i := 2; i < len(tokens); i++ {
			fen += tokens[i] + " "
		}

		pos.load(fen)
	}
}

func handleMove(pos *Position, tokens []string) {
	moveStr := strings.ToLower(tokens[1])
	handleUciMove(pos, moveStr)
}

func handleUciMove(pos *Position, moveStr string) {
	moves, nrMoves := generateMoves(*pos)

	// Find the move by comparing move strings
	for i := 0; i < nrMoves; i++ {
		move := moves[i]
		if !pos.isMoveLegal(move, false) {
			continue
		}

		if moveStr == move.toStr() {
			pos.makeMove(move)
			return
		}
	}

	fmt.Printf("Invalid move: %s", moveStr)
}

func uciLoop(pos *Position) {
	for {
		in := bufio.NewReader(os.Stdin)
		line, _ := in.ReadString('\n')

		var nltokens = strings.Split(line, "\n")
		var tokens = strings.Split(nltokens[0], " ")
		exit := false

		switch tokens[0] {
		case "uci":
			fmt.Printf("id name Axii id author qantata\nuciok")
		case "position":
			setPosition(pos, tokens)
		case "showpos":
			pos.pretty()
		case "go":
			searchGo(*pos)
		case "ucinewgame":
			pos.load(STARTING_POS)
		case "move":
			handleMove(pos, tokens)
		case "isready":
			fmt.Println("readyok")
		case "exit":
			exit = true
		case "quit":
			exit = true
		default:
			fmt.Printf("Unknown token %s\n", tokens[0])
		}

		fmt.Println()

		if exit {
			fmt.Println("Goodbye.")
			break
		}
	}
}
