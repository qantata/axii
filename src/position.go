package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	board           []uint64
	pieces          []uint64
	turn            uint8
	castlingRights  uint8
	epSquare        uint8
	halfMoveClock   int
	fullMoveCounter int
}

func (p *Position) reset() {
	p.board = make([]uint64, NR_SQUARES)
	p.pieces = make([]uint64, 12)
	p.turn = SIDE_WHITE
	p.castlingRights = NO_CASTLING_RIGHTS
	p.epSquare = NO_SQUARE
	p.halfMoveClock = 0
	p.fullMoveCounter = 0
}

func (p *Position) load(fen string) {
	p.reset()

	s := strings.Split(fen, " ")
	square := SQ_A8

	// Pieces
	for i := 0; i < len(s[0]); i++ {
		letter := string(fen[i])

		if num, err := strconv.Atoi(letter); err == nil {
			square += num
		} else {
			var piece uint64 = 0
			var side uint64 = PIECE_BLACK

			if letter == strings.ToUpper(letter) {
				side = PIECE_WHITE
			}

			switch strings.ToLower(letter) {
			case "p":
				piece = PIECE_PAWN
			case "r":
				piece = PIECE_ROOK
			case "n":
				piece = PIECE_KNIGHT
			case "b":
				piece = PIECE_BISHOP
			case "q":
				piece = PIECE_QUEEN
			case "k":
				piece = PIECE_KING
			case "/":
				// Jump to the rank below, to the A file
				square -= 16
				continue
			default:
				fmt.Printf("Error: %s is not recognized as any piece", letter)
			}

			p.board[square] = piece | side
			square += 1
		}
	}

	// Side to move
	if s[1] == "w" {
		p.turn = SIDE_WHITE
	} else {
		p.turn = SIDE_BLACK
	}

	// Castling rights
	if strings.Contains(s[2], "K") {
		p.castlingRights |= CASTLING_WHITE_OO
	}

	if strings.Contains(s[2], "Q") {
		p.castlingRights |= CASTLING_WHITE_OOO
	}

	if strings.Contains(s[2], "k") {
		p.castlingRights |= CASTLING_BLACK_OO
	}

	if strings.Contains(s[2], "q") {
		p.castlingRights |= CASTLING_BLACK_OOO
	}

	// En passant square
	if s[3] != "-" {
		p.epSquare = SSM[s[3]]
	}

	// Halfmove clock
	if num, err := strconv.Atoi(string(s[4])); err == nil {
		p.halfMoveClock = num
	} else {
		fmt.Printf("Error: %s is not a valid halfmove clock value!", s[4])
	}

	// Full move counter
	if num, err := strconv.Atoi(string(s[5])); err == nil {
		p.fullMoveCounter = num
	} else {
		fmt.Printf("Error: %s is not a valid full move counter value!", s[5])
	}

	fmt.Printf("Loaded position %s\n", fen)
}

func (p *Position) pretty() {
	piecesToChars := "         PRNBQK  prnbqk"
	result := ""
	separator := "+---+---+---+---+---+---+---+---+"

	result += separator + "\n"
	result += "| "

	square := SQ_A8
	for {
		if p.board[square] != NO_PIECE {
			result += string(piecesToChars[p.board[square]]) + " | "
		} else {
			result += "  | "
		}

		square++
		if (square % 8) == 0 {
			square -= 16
			result += "\n" + separator + "\n"

			if square < 0 {
				break
			} else {
				result += "| "
			}
		}
	}

	fmt.Println(result)
}
