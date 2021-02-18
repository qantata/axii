package main

var rookDistances [64][]int
var bishopDistances [64][]int

var mvvlvaTable [16][16]int
var mvvlvaVictimScores = []int{0, 100, 400, 200, 300, 500, 600, 0, 0, 100, 400, 200, 300, 500, 600}

// MVVLVA constants
const (
	MVVLVA_PADDING   = 1000000
	KILLER_PRIMARY   = 900000
	KILLER_SECONDARY = 800000
)

const (
	MOVEGEN_DIR_N = 0
	MOVEGEN_DIR_E = 1
	MOVEGEN_DIR_S = 2
	MOVEGEN_DIR_W = 3
)

const piecesToChars = " PRNBQK  prnbqk"

func initMoveGen() {
	var order = []int{WHITE_PAWN, WHITE_KNIGHT, WHITE_BISHOP, WHITE_ROOK, WHITE_QUEEN, WHITE_KING, BLACK_PAWN, BLACK_KNIGHT, BLACK_BISHOP, BLACK_ROOK, BLACK_QUEEN, BLACK_KING}

	for i := 0; i < len(order); i++ {
		for j := 0; j < len(order); j++ {
			attacker := order[i]
			victim := order[j]
			mvvlvaTable[victim][attacker] = (mvvlvaVictimScores[victim] + 6 - (mvvlvaVictimScores[attacker] / 100))
		}
	}
}

func generateMoves(pos Position) ([MAX_MOVES_IN_POS]Move, int) {
	var moves [MAX_MOVES_IN_POS]Move
	var index int = 0

	generatePawnMoves(pos, &moves, &index, false)
	generateSlidingMoves(pos, &moves, &index, SLIDING_GEN_TYPE_ROOK, false)
	generateKnightMoves(pos, &moves, &index, false)
	generateSlidingMoves(pos, &moves, &index, SLIDING_GEN_TYPE_BISHOP, false)
	generateSlidingMoves(pos, &moves, &index, SLIDING_GEN_TYPE_QUEEN, false)
	generateKingMoves(pos, &moves, &index, false)

	return moves, index
}

func generateCaptures(pos Position) ([MAX_MOVES_IN_POS]Move, int) {
	var moves [MAX_MOVES_IN_POS]Move
	var index int = 0

	generatePawnMoves(pos, &moves, &index, true)
	generateSlidingMoves(pos, &moves, &index, SLIDING_GEN_TYPE_ROOK, true)
	generateKnightMoves(pos, &moves, &index, true)
	generateSlidingMoves(pos, &moves, &index, SLIDING_GEN_TYPE_BISHOP, true)
	generateSlidingMoves(pos, &moves, &index, SLIDING_GEN_TYPE_QUEEN, true)
	generateKingMoves(pos, &moves, &index, true)

	return moves, index
}

func generatePawnMoves(pos Position, moves *[MAX_MOVES_IN_POS]Move, index *int, onlyCaptures bool) {
	piece := createPiece(pos.turn, PIECE_PAWN)
	pieceCount := pos.getPieceCount(piece)
	opponentSide := pos.getOpponentSide()

	for i := 0; i < pieceCount; i++ {
		origSquare := pos.pieceList[piece][i]
		rank := getRankOf(origSquare)
		file := getFileOf(origSquare)
		isPromotion := (rank == RANK_7 && pos.turn == SIDE_WHITE) || (rank == RANK_2 && pos.turn == SIDE_BLACK)

		// Captures and EP
		pawnCaptureDir1 := DIR_NE
		pawnCaptureDir2 := DIR_NW

		if pos.turn == SIDE_BLACK {
			pawnCaptureDir1 = DIR_SE
			pawnCaptureDir2 = DIR_SW
		}

		if file < FILE_H {
			pawnCapture1Square := origSquare + pawnCaptureDir1
			pawnCapture1Piece := pos.getPieceOn(pawnCapture1Square)
			if (pawnCapture1Square == pos.state.epSquare) || (pawnCapture1Piece != NO_PIECE && getColorOf(pawnCapture1Piece) == opponentSide) {
				if isPromotion {
					for promotePiece := PROMOTION_PIECE_ROOK; promotePiece <= PROMOTION_PIECE_QUEEN; promotePiece++ {
						move := Move{}
						move.createWithPromotion(origSquare, pawnCapture1Square, promotePiece)
						move.score = (mvvlvaVictimScores[promotePiece+2] - mvvlvaVictimScores[PIECE_PAWN]) + mvvlvaTable[pawnCapture1Piece][PIECE_PAWN]

						moves[*index] = move
						*index++
					}
				} else {
					pawnCapture1Move := Move{}
					pawnCapture1Move.score = mvvlvaTable[pawnCapture1Piece][PIECE_PAWN]

					if pawnCapture1Square == pos.state.epSquare {
						pawnCapture1Move.createWithEnPassant(origSquare, pawnCapture1Square)
					} else {
						pawnCapture1Move.create(origSquare, pawnCapture1Square)
					}

					moves[*index] = pawnCapture1Move
					*index++
				}
			}
		}

		if file > FILE_A {
			pawnCapture2Square := origSquare + pawnCaptureDir2
			pawnCapture2Piece := pos.getPieceOn(pawnCapture2Square)
			if (pawnCapture2Square == pos.state.epSquare) || (pawnCapture2Piece != NO_PIECE && getColorOf(pawnCapture2Piece) == opponentSide) {
				if isPromotion {
					for promotePiece := PROMOTION_PIECE_ROOK; promotePiece <= PROMOTION_PIECE_QUEEN; promotePiece++ {
						move := Move{}
						move.createWithPromotion(origSquare, pawnCapture2Square, promotePiece)
						move.score = (mvvlvaVictimScores[promotePiece+2] - mvvlvaVictimScores[PIECE_PAWN]) + mvvlvaTable[pawnCapture2Piece][PIECE_PAWN]

						moves[*index] = move
						*index++
					}
				} else {
					pawnCapture2Move := Move{}
					pawnCapture2Move.score = mvvlvaTable[pawnCapture2Piece][PIECE_PAWN]

					if pawnCapture2Square == pos.state.epSquare {
						pawnCapture2Move.createWithEnPassant(origSquare, pawnCapture2Square)
					} else {
						pawnCapture2Move.create(origSquare, pawnCapture2Square)
					}

					moves[*index] = pawnCapture2Move
					*index++
				}
			}
		}

		if onlyCaptures {
			continue
		}

		// Pawn pushes
		pawnPushDir := DIR_N

		if pos.turn == SIDE_BLACK {
			pawnPushDir = DIR_S
		}

		// Single push
		singlePushSquare := origSquare + pawnPushDir
		if pos.getPieceOn(singlePushSquare) == NO_PIECE {
			if isPromotion {
				for promotePiece := PROMOTION_PIECE_ROOK; promotePiece <= PROMOTION_PIECE_QUEEN; promotePiece++ {
					move := Move{}
					move.createWithPromotion(origSquare, singlePushSquare, promotePiece)
					moves[*index] = move
					*index++
				}
			} else {
				singlePushMove := Move{}
				singlePushMove.create(origSquare, singlePushSquare)

				moves[*index] = singlePushMove
				*index++
			}

			// Double push
			if (rank == RANK_2 && pos.turn == SIDE_WHITE) || (rank == RANK_7 && pos.turn == SIDE_BLACK) {
				doublePushSquare := origSquare + (pawnPushDir * 2)
				if pos.getPieceOn(doublePushSquare) == NO_PIECE {
					doublePushMove := Move{}
					doublePushMove.create(origSquare, doublePushSquare)

					moves[*index] = doublePushMove
					*index++
				}
			}
		}
	}
}

const (
	SLIDING_GEN_TYPE_BISHOP = 0
	SLIDING_GEN_TYPE_ROOK   = 1
	SLIDING_GEN_TYPE_QUEEN  = 2
)

func generateSlidingMoves(pos Position, moves *[MAX_MOVES_IN_POS]Move, index *int, genType int, onlyCaptures bool) {
	piece := PIECE_BISHOP

	if genType == SLIDING_GEN_TYPE_ROOK {
		piece = PIECE_ROOK
	} else if genType == SLIDING_GEN_TYPE_QUEEN {
		piece = PIECE_QUEEN
	}

	piece = createPiece(pos.turn, piece)
	pieceCount := pos.getPieceCount(piece)

	for i := 0; i < pieceCount; i++ {
		origSquare := pos.pieceList[piece][i]

		var attacks uint64 = 0
		occupancy := pos.getAllPieces()

		// Get attacks bitboard
		if genType == SLIDING_GEN_TYPE_BISHOP {
			attacks |= getBishopAttacks(origSquare, occupancy)
		} else if genType == SLIDING_GEN_TYPE_ROOK {
			attacks |= getRookAttacks(origSquare, occupancy)
		} else {
			attacks |= (getBishopAttacks(origSquare, occupancy) | getRookAttacks(origSquare, occupancy))
		}

		// Only use relevant square & pieces
		if onlyCaptures {
			attacks &= pos.getPiecesByColor(pos.getOpponentSide())
		} else {
			attacks &= ^pos.getPiecesByColor(pos.turn)
		}

		// Go through the attack moves
		for attacks != 0 {
			square := bbPopLSB(&attacks)
			move := Move{}

			capPiece := pos.getPieceOn(square)
			if capPiece != NO_PIECE {
				move.score = mvvlvaTable[capPiece][piece]
			}

			move.create(origSquare, square)

			moves[*index] = move
			*index++
		}
	}
}

func generateKnightMoves(pos Position, moves *[MAX_MOVES_IN_POS]Move, index *int, onlyCaptures bool) {
	piece := createPiece(pos.turn, PIECE_KNIGHT)
	pieceCount := pos.getPieceCount(piece)

	for i := 0; i < pieceCount; i++ {
		origSquare := pos.pieceList[piece][i]

		attacks := getKnightAttacks(origSquare)

		if onlyCaptures {
			attacks &= pos.getPiecesByColor(pos.getOpponentSide())
		} else {
			attacks &= ^pos.getPiecesByColor(pos.turn)
		}

		for attacks != 0 {
			square := bbPopLSB(&attacks)
			move := Move{}

			capPiece := pos.getPieceOn(square)
			if capPiece != NO_PIECE {
				move.score = mvvlvaTable[capPiece][piece]
			}

			move.create(origSquare, square)

			moves[*index] = move
			*index++
		}
	}
}

func generateKingMoves(pos Position, moves *[MAX_MOVES_IN_POS]Move, index *int, onlyCaptures bool) {
	piece := createPiece(pos.turn, PIECE_KING)
	pieceCount := pos.getPieceCount(piece)

	for i := 0; i < pieceCount; i++ {
		origSquare := pos.pieceList[piece][i]

		// Regular moves
		attacks := getKingAttacks(origSquare)

		if onlyCaptures {
			attacks &= pos.getPiecesByColor(pos.getOpponentSide())
		} else {
			attacks &= ^pos.getPiecesByColor(pos.turn)
		}

		for attacks != 0 {
			square := bbPopLSB(&attacks)
			move := Move{}

			capPiece := pos.getPieceOn(square)
			if capPiece != NO_PIECE {
				move.score = mvvlvaTable[capPiece][piece]
			}

			move.create(origSquare, square)

			moves[*index] = move
			*index++
		}

		if onlyCaptures {
			continue
		}

		// Castling
		if pos.canCastleShort(pos.turn) {
			if pos.getPieceOn(origSquare+1) == NO_PIECE && pos.getPieceOn(origSquare+2) == NO_PIECE {
				move := Move{}
				move.createWithCastling(origSquare, origSquare+2)

				moves[*index] = move
				*index++
			}
		}

		if pos.canCastleLong(pos.turn) {
			if pos.getPieceOn(origSquare-1) == NO_PIECE && pos.getPieceOn(origSquare-2) == NO_PIECE && pos.getPieceOn(origSquare-3) == NO_PIECE {
				move := Move{}
				move.createWithCastling(origSquare, origSquare-2)

				moves[*index] = move
				*index++
			}
		}
	}
}
