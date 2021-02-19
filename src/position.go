package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var zobristEpFile [8]uint64
var zobristSideToMove uint64
var zobristPieceMap [16][NR_SQUARES]uint64
var zobristCastlingRights [NR_CASTLING]uint64

type BoardState struct {
	key            uint64
	castlingRights int
	epSquare       int
	rule50         int
	capturedPiece  int
	prevState      *BoardState
}

type Position struct {
	board           []int
	pieces          []uint64
	turn            int
	halfMoveClock   int
	fullMoveCounter int
	pieceCount      []int
	pieceList       [][10]int
	pieceIndex      []int
	piecesByType    []uint64
	piecesByColor   []uint64
	state           *BoardState
}

func (p *Position) reset() {
	p.board = make([]int, NR_SQUARES)
	p.pieces = make([]uint64, MAX_PIECE_NR)
	p.turn = SIDE_WHITE
	p.halfMoveClock = 0
	p.fullMoveCounter = 0
	p.pieceIndex = make([]int, NR_SQUARES)
	p.pieceCount = make([]int, MAX_PIECE_NR)
	p.pieceList = make([][10]int, MAX_PIECE_NR)
	p.piecesByType = make([]uint64, NR_PIECE_TYPES)
	p.piecesByColor = make([]uint64, 2)
	p.state = nil
}

func initPosition() {
	rand.Seed(133742069)

	for piece := PIECE_PAWN; piece <= BLACK_KING; piece++ {
		for square := SQ_A1; square <= SQ_H8; square++ {
			zobristPieceMap[piece][square] = rand.Uint64()
		}
	}

	for file := FILE_A; file <= FILE_H; file++ {
		zobristEpFile[file] = rand.Uint64()
	}

	for cr := NO_CASTLING_RIGHTS; cr <= CASTLING_ALL; cr++ {
		zobristCastlingRights[cr] = rand.Uint64()
	}

	zobristSideToMove = rand.Uint64()
}

func (p *Position) load(fen string) {
	p.reset()
	p.state = new(BoardState)

	s := strings.Split(fen, " ")
	square := SQ_A8

	// Pieces
	for i := 0; i < len(s[0]); i++ {
		letter := string(fen[i])

		if num, err := strconv.Atoi(letter); err == nil {
			square += num
		} else {
			var piece int = 0
			var side int = SIDE_BLACK

			if letter == strings.ToUpper(letter) {
				side = SIDE_WHITE
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

			p.putPiece(createPiece(side, piece), square)
			square += 1
		}
	}

	// Side to move
	if s[1] == "w" {
		p.turn = SIDE_WHITE
	} else {
		p.turn = SIDE_BLACK
		p.state.key ^= zobristSideToMove
	}

	// Castling rights
	if strings.Contains(s[2], "K") {
		p.state.castlingRights |= CASTLING_WHITE_OO
	}

	if strings.Contains(s[2], "Q") {
		p.state.castlingRights |= CASTLING_WHITE_OOO
	}

	if strings.Contains(s[2], "k") {
		p.state.castlingRights |= CASTLING_BLACK_OO
	}

	if strings.Contains(s[2], "q") {
		p.state.castlingRights |= CASTLING_BLACK_OOO
	}

	p.state.key ^= zobristCastlingRights[p.state.castlingRights]

	// En passant square
	if s[3] != "-" {
		epSquare := SSM[s[3]]
		p.state.epSquare = epSquare
		p.state.key ^= zobristEpFile[getFileOf(epSquare)]
	}

	// Halfmove clock
	if len(s) > 4 {
		if num, err := strconv.Atoi(string(s[4])); err == nil {
			p.halfMoveClock = num
		}
	}

	// Full move counter
	if len(s) > 5 {
		if num, err := strconv.Atoi(string(s[5])); err == nil {
			p.fullMoveCounter = num
		}
	}
}

func (p Position) isInCheck() bool {
	kingSquare := p.getSquareWithPieceType(p.turn, PIECE_KING)
	attackers := getAttackersBBToSquare(kingSquare, p.getAllPieces(), p)

	opponentPieces := p.getPiecesByColor(p.getOpponentSide())

	return (attackers & opponentPieces) != 0
}

func (p Position) isInsufficientMaterial() bool {
	if p.getPiecesByType(PIECE_PAWN) != 0 || p.getPiecesByType(PIECE_ROOK) != 0 || p.getPiecesByType(PIECE_QUEEN) != 0 {
		return false
	}

	myKnights := p.getPiecesByTypeAndColor(PIECE_KNIGHT, p.turn)
	oppKnights := p.getPiecesByTypeAndColor(PIECE_KNIGHT, p.getOpponentSide())
	myBishops := p.getPiecesByTypeAndColor(PIECE_BISHOP, p.turn)
	oppBishops := p.getPiecesByTypeAndColor(PIECE_BISHOP, p.getOpponentSide())

	// Knight + bishop is not a draw
	if myKnights != 0 && myBishops != 0 {
		return false
	}

	if oppKnights != 0 && oppBishops != 0 {
		return false
	}

	// Bishop + bishop is not a draw
	if p.pieceCount[WHITE_BISHOP] >= 2 || p.pieceCount[BLACK_BISHOP] >= 2 {
		return false
	}

	// More than 2 knights is not a draw
	if p.pieceCount[WHITE_KNIGHT] > 2 || p.pieceCount[BLACK_KNIGHT] > 2 {
		return false
	}

	return true
}

func (p Position) evaluate() int {
	return evaluatePosition(p)
}

func (p Position) getPieceCount(piece int) int {
	return p.pieceCount[piece]
}

func (p Position) getPieceOn(square int) int {
	return p.board[square]
}

func (p Position) getOpponentSide() int {
	return p.turn ^ 1
}

func (p Position) getPiecesByType(pieceType int) uint64 {
	return p.piecesByType[pieceType]
}

func (p Position) getPiecesByTypeAndColor(pieceType int, color int) uint64 {
	return p.getPiecesByType(pieceType) & p.getPiecesByColor(color)
}

func (p Position) getAllPieces() uint64 {
	return p.piecesByType[PIECE_TYPE_ALL]
}

func (p Position) getPiecesByColor(color int) uint64 {
	return p.piecesByColor[color]
}

func (p Position) canCastleShort(color int) bool {
	if color == SIDE_WHITE {
		return (p.state.castlingRights & CASTLING_WHITE_OO) != 0
	} else {
		return (p.state.castlingRights & CASTLING_BLACK_OO) != 0
	}
}

func (p Position) canCastleLong(color int) bool {
	if color == SIDE_WHITE {
		return (p.state.castlingRights & CASTLING_WHITE_OOO) != 0
	} else {
		return (p.state.castlingRights & CASTLING_BLACK_OOO) != 0
	}
}

func (p *Position) makeMove(move Move) {
	from := move.orig()
	to := move.dest()
	piece := p.getPieceOn(from)
	moveType := move.typeOf()
	pieceType := getTypeOf(piece)
	opponentSide := p.getOpponentSide()

	capturedPiece := p.getPieceOn(to)
	capturedSquare := to

	newState := new(BoardState)
	newState.castlingRights = p.state.castlingRights
	newState.capturedPiece = p.state.capturedPiece
	newState.epSquare = p.state.epSquare
	newState.rule50 = p.state.epSquare
	newState.prevState = p.state
	newState.key = p.state.key
	p.state = newState

	p.state.key ^= zobristSideToMove
	p.state.rule50++

	if moveType == MOVE_TYPE_ENPASSANT {
		if p.turn == SIDE_WHITE {
			capturedPiece = p.getPieceOn(to + DIR_S)
			capturedSquare = to + DIR_S
		} else {
			capturedPiece = p.getPieceOn(to + DIR_N)
			capturedSquare = to + DIR_N
		}
	}

	capturedPieceType := getTypeOf(capturedPiece)
	if capturedPieceType == PIECE_KING {
		p.pretty()
		fmt.Println("King was captured!")
		return
	}

	// If a rook was captured, check if it needs to update castling rights
	if capturedPieceType == PIECE_ROOK {
		canOpponentCastleShort := p.canCastleShort(opponentSide)
		canOpponentCastleLong := p.canCastleLong(opponentSide)

		if to == SQ_H8 && p.turn == SIDE_WHITE && canOpponentCastleShort {
			p.state.castlingRights ^= CASTLING_BLACK_OO
			p.state.key ^= zobristCastlingRights[CASTLING_BLACK_OO]
		} else if to == SQ_H1 && p.turn == SIDE_BLACK && canOpponentCastleShort {
			p.state.castlingRights ^= CASTLING_WHITE_OO
			p.state.key ^= zobristCastlingRights[CASTLING_WHITE_OO]
		} else if to == SQ_A8 && p.turn == SIDE_WHITE && canOpponentCastleLong {
			p.state.castlingRights ^= CASTLING_BLACK_OOO
			p.state.key ^= zobristCastlingRights[CASTLING_BLACK_OOO]
		} else if to == SQ_A1 && p.turn == SIDE_BLACK && canOpponentCastleLong {
			p.state.castlingRights ^= CASTLING_WHITE_OOO
			p.state.key ^= zobristCastlingRights[CASTLING_WHITE_OOO]
		}
	}

	// Castling
	if moveType == MOVE_TYPE_CASTLING {
		rookFrom := SQ_A1
		rookTo := SQ_D1

		if to > from {
			rookFrom = SQ_H1
			rookTo = SQ_F1
		}

		if p.turn == SIDE_BLACK {
			rookFrom = flipSquare(rookFrom)
			rookTo = flipSquare(rookTo)
		}

		p.movePiece(rookFrom, rookTo)
	}

	// Update castling rights on king or rook move
	if pieceType == PIECE_KING || pieceType == PIECE_ROOK {
		if p.turn == SIDE_WHITE {
			isFromE1 := from == SQ_E1

			if (isFromE1 || from == SQ_H1) && p.canCastleShort(SIDE_WHITE) {
				p.state.castlingRights ^= CASTLING_WHITE_OO
				p.state.key ^= zobristCastlingRights[CASTLING_WHITE_OO]
			}

			if (isFromE1 || from == SQ_A1) && p.canCastleLong(SIDE_WHITE) {
				p.state.castlingRights ^= CASTLING_WHITE_OOO
				p.state.key ^= zobristCastlingRights[CASTLING_WHITE_OOO]
			}
		} else {
			isFromE8 := from == SQ_E8

			if (isFromE8 || from == SQ_H8) && p.canCastleShort(SIDE_BLACK) {
				p.state.castlingRights ^= CASTLING_BLACK_OO
				p.state.key ^= zobristCastlingRights[CASTLING_BLACK_OO]
			}

			if (isFromE8 || from == SQ_A8) && p.canCastleLong(SIDE_BLACK) {
				p.state.castlingRights ^= CASTLING_BLACK_OOO
				p.state.key ^= zobristCastlingRights[CASTLING_BLACK_OOO]
			}
		}
	}

	if capturedPiece != NO_PIECE {
		p.removePiece(capturedPiece, capturedSquare)

		p.state.rule50 = 0
	}

	if p.state.epSquare != NO_SQUARE {
		p.state.key ^= zobristEpFile[getFileOf(p.state.epSquare)]
		p.state.epSquare = NO_SQUARE
	}

	if pieceType == PIECE_PAWN {
		// En passant
		if (to ^ from) == 16 {
			if p.turn == SIDE_WHITE {
				p.state.epSquare = to + DIR_S
			} else {
				p.state.epSquare = to + DIR_N
			}

			p.state.key ^= zobristEpFile[getFileOf(p.state.epSquare)]
		}

		// Promotion
		if moveType == MOVE_TYPE_PROMOTION {
			promotedPiece := createPiece(p.turn, move.promotionPieceType())

			p.removePiece(piece, from)
			p.putPiece(promotedPiece, to)
		}

		p.state.rule50 = 0
	}

	if moveType != MOVE_TYPE_PROMOTION {
		p.movePiece(from, to)
	}

	p.halfMoveClock++

	p.turn = p.getOpponentSide()
	p.state.capturedPiece = capturedPiece
}

func (p *Position) unmakeMove(move Move) {
	p.turn = p.getOpponentSide()

	from := move.orig()
	to := move.dest()
	piece := p.getPieceOn(to)
	moveType := move.typeOf()

	// Reverse castling
	if moveType == MOVE_TYPE_CASTLING {
		rookFrom := SQ_D1
		rookTo := SQ_A1

		if to > from {
			rookFrom = SQ_F1
			rookTo = SQ_H1
		}

		if p.turn == SIDE_BLACK {
			rookFrom = flipSquare(rookFrom)
			rookTo = flipSquare(rookTo)
		}

		p.movePiece(rookFrom, rookTo)
	}

	// Reverse promotion
	if moveType == MOVE_TYPE_PROMOTION {
		p.removePiece(piece, to)
		p.putPiece(createPiece(p.turn, PIECE_PAWN), from)
	} else {
		// Move piece back
		p.movePiece(to, from)
	}

	// Put back captured piece
	if p.state.capturedPiece != NO_PIECE {
		captureSquare := to

		if moveType == MOVE_TYPE_ENPASSANT {
			if p.turn == SIDE_WHITE {
				captureSquare += DIR_S
			} else {
				captureSquare += DIR_N
			}
		}

		p.putPiece(p.state.capturedPiece, captureSquare)
	}

	p.halfMoveClock--

	oldState := p.state.prevState
	p.state = oldState
}

func (p Position) getSquareWithPieceType(side int, pieceType int) int {
	piece := createPiece(side, pieceType)

	return p.pieceList[piece][0]
}

func (p *Position) isMoveLegal(move Move, makeMoveIfLegal bool) bool {
	from := move.orig()
	to := move.dest()
	moveType := move.typeOf()
	kingSquare := p.getSquareWithPieceType(p.turn, PIECE_TYPE_KING)
	opponentSide := p.getOpponentSide()
	pieceType := getTypeOf(p.getPieceOn(from))

	if (getColorOf(p.getPieceOn(from)) != p.turn) || (p.getPieceOn(from) == NO_PIECE) {
		return false
	}

	// We separately check en passant since only rooks, bishops and queens can
	// have a discovered check after an en passant move, so it's more efficient
	if moveType == MOVE_TYPE_ENPASSANT {
		captureSquare := to

		if p.turn == SIDE_WHITE {
			captureSquare += DIR_S
		} else {
			captureSquare += DIR_N
		}

		// Get new occupancy bitboard after en passant
		occupied := (p.getAllPieces() ^ getBBOfSquare(from) ^ getBBOfSquare(captureSquare)) | getBBOfSquare(to)

		queens := p.getPiecesByTypeAndColor(PIECE_QUEEN, opponentSide)
		rooks := p.getPiecesByTypeAndColor(PIECE_ROOK, opponentSide)
		bishops := p.getPiecesByTypeAndColor(PIECE_BISHOP, opponentSide)

		// Check the attack boards from the king's point of view and if they hit any
		// of the bishops, queens or rooks, then the king is in check
		if (getBishopAttacks(kingSquare, occupied) & (queens | bishops)) != 0 {
			return false
		}

		if (getRookAttacks(kingSquare, occupied) & (queens | rooks)) != 0 {
			return false
		}

		if makeMoveIfLegal {
			p.makeMove(move)
		}

		return true
	}

	opponentPieces := p.getPiecesByColor(opponentSide)

	// Check castling
	if moveType == MOVE_TYPE_CASTLING {
		steps := 0
		nrSteps := 3
		dir := DIR_W

		if to > from {
			dir = DIR_E
		}

		allPieces := p.getAllPieces()
		for square := kingSquare; steps < nrSteps; square, steps = square+dir, steps+1 {
			if (getAttackersBBToSquare(square, allPieces, *p) & opponentPieces) != 0 {
				return false
			}
		}

		if makeMoveIfLegal {
			p.makeMove(move)
		}

		return true
	}

	// Check that the king isn't walking into a check
	if pieceType == PIECE_KING {
		pieces := p.getAllPieces()
		pieces ^= (getBBOfSquare(from) ^ getBBOfSquare(to))

		if (getAttackersBBToSquare(to, pieces, *p) & opponentPieces) != 0 {
			return false
		}

		if makeMoveIfLegal {
			p.makeMove(move)
		}

		return true
	}

	// Otherwise, make the move and check if the king ended up in a check
	p.makeMove(move)

	// Use opponent king since we made a move
	// TODO: CHECK king square
	ownPieces := p.getPiecesByColor(p.turn)
	kingSquareNew := p.getSquareWithPieceType(p.getOpponentSide(), PIECE_KING)

	isLegal := true
	if (getAttackersBBToSquare(kingSquareNew, p.getAllPieces(), *p) & ownPieces) != 0 {
		isLegal = false
	}

	if !isLegal || !makeMoveIfLegal {
		p.unmakeMove(move)
	}

	return isLegal
}

func (p *Position) putPiece(piece int, square int) {
	x := getColorOf(piece)
	bbOfSquare := getBBOfSquare(square)
	p.board[square] = piece                          // Put the piece on the square
	p.pieceList[piece][p.pieceCount[piece]] = square // Update piece list
	p.pieceIndex[square] = p.pieceCount[piece]       // Update piece index list
	p.pieceCount[piece]++                            // Update piece count
	p.piecesByType[getTypeOf(piece)] |= bbOfSquare   // Update piece counts per type
	p.piecesByType[PIECE_TYPE_ALL] |= bbOfSquare     // Also update all types so we can easily get all pieces on the board
	p.piecesByColor[x] |= bbOfSquare                 // Update pieces by color
	p.state.key ^= zobristPieceMap[piece][square]    // Update zobrist key
}

func (p *Position) removePiece(piece int, square int) {
	bbOfSquare := getBBOfSquare(square)

	// Update board
	p.board[square] = NO_PIECE

	// Update bitboard tables
	p.piecesByColor[getColorOf(piece)] ^= bbOfSquare
	p.piecesByType[getTypeOf(piece)] ^= bbOfSquare
	p.piecesByType[PIECE_TYPE_ALL] ^= bbOfSquare

	// Update piece count
	p.pieceCount[piece]--

	/*
		Since we're deleting a piece from the piece list, we need to make sure there are
		no empty gaps left in the pieceList array. We do this by moving the last element of the
		pieceList array into the spot of the removed piece. More comments of what each line does below.
	*/
	lastSquare := p.pieceList[piece][p.pieceCount[piece]] // Get the square of the last element in the piece list array
	p.pieceIndex[lastSquare] = p.pieceIndex[square]       // Set the index of the last pieceList array element to the removed piece index
	p.pieceList[piece][p.pieceIndex[square]] = lastSquare // Move the last element into the spot of the removed one, so that we can continue to grow the array normally
	p.pieceList[piece][p.pieceCount[piece]] = NO_SQUARE   // Now reset the last element since it was moved

	p.state.key ^= zobristPieceMap[piece][square] // Update zobrist key
}

func (p *Position) movePiece(from int, to int) {
	piece := p.getPieceOn(from)
	fromToBB := getBBOfSquare(from) ^ getBBOfSquare(to)

	// Update board
	p.board[to] = piece
	p.board[from] = NO_PIECE

	// Update bitboard tables
	p.piecesByColor[getColorOf(piece)] ^= fromToBB
	p.piecesByType[getTypeOf(piece)] ^= fromToBB
	p.piecesByType[PIECE_TYPE_ALL] ^= fromToBB

	// Update piece index & list
	p.pieceIndex[to] = p.pieceIndex[from] // Don't need reset from square since we're not gonna access it
	p.pieceList[piece][p.pieceIndex[to]] = to
	p.state.key ^= zobristPieceMap[piece][from] ^ zobristPieceMap[piece][to]
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

	result += "\nKey: " + strconv.FormatUint(p.state.key, 16)

	result += "\n"

	fmt.Println(result)
}
