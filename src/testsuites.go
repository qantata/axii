package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func runSingleTest(epd string) {
	fmt.Printf("Running test %s...\t", epd)

	parts := strings.Split(epd, " ")
	var bestMoves = []string{}

	fen := ""
	for i := 0; i < len(parts); i++ {
		part := parts[i]

		if part == "bm" {
			bestMoves = strings.Split(parts[i+1], ",")
			break
		} else {
			fen += part + " "
		}
	}

	var pos = Position{}
	pos.load(fen)

	searchGo(pos)
	for i := 0; i < len(bestMoves); i++ {
		if pvMove.toStr() == bestMoves[i] {
			fmt.Printf("passed\n")
			return
		}
	}

	fmt.Printf("failed\n")
}

func runTestSuites() {
	TESTING = true

	var files = []string{"bratko-kopec"}

	file, err := os.Open("testsuites/" + files[0] + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		runSingleTest(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
