package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	board           []int
	pieces          []uint64
	turn            int
	castlingRights  int
	epSquare        int
	halfMoveClock   int
	fullMoveCounter int
	pieceCount      []int
	pieceList       [][10]int
}

func (p *Position) reset() {
	p.board = make([]int, NR_SQUARES)
	p.pieces = make([]uint64, MAX_PIECE_NR)
	p.turn = SIDE_WHITE
	p.castlingRights = NO_CASTLING_RIGHTS
	p.epSquare = NO_SQUARE
	p.halfMoveClock = 0
	p.fullMoveCounter = 0
	p.pieceCount = make([]int, MAX_PIECE_NR)
	p.pieceList = make([][10]int, MAX_PIECE_NR)
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
			var piece int = 0
			var side int = PIECE_BLACK

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

			p.putPiece(piece|side, square)
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

func (p Position) getPieceCount(piece int) int {
	return p.pieceCount[piece]
}

func (p Position) getPieceOn(square int) int {
	return p.board[square]
}

func (p *Position) putPiece(piece int, square int) {
	p.board[square] = piece                          // Put the piece on the square
	p.pieceList[piece][p.pieceCount[piece]] = square // Update piece list
	p.pieceCount[piece]++                            // Update piece count
}

func (p *Position) pretty() {
	piecesToChars := " PRNBQK  prnbqk"
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

	result += "\nSide to move: "

	if p.turn == SIDE_WHITE {
		result += "white"
	} else {
		result += "black"
	}

	result += "\n"

	fmt.Println(result)
}
