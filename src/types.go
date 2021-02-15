package main

/*
	256 here is used as a safe number for the maximum
	possible number of legal moves in any given position.
*/
const MAX_MOVES_IN_POS = 256

const NO_MOVE = 0

// Sides
const (
	SIDE_WHITE = 0
	SIDE_BLACK = 1
)

// Pieces
const (
	NO_PIECE     = 0
	PIECE_PAWN   = 1
	PIECE_ROOK   = 2
	PIECE_KNIGHT = 3
	PIECE_BISHOP = 4
	PIECE_QUEEN  = 5
	PIECE_KING   = 6
	PIECE_WHITE  = 0
	PIECE_BLACK  = 8
	MAX_PIECE_NR = 15
)

// Promotion pieces
const (
	PROMOTION_PIECE_ROOK   = 0
	PROMOTION_PIECE_KNIGHT = 1
	PROMOTION_PIECE_BISHOP = 2
	PROMOTION_PIECE_QUEEN  = 3
)

// Move types
const (
	MOVE_TYPE_NORMAL    = 0
	MOVE_TYPE_PROMOTION = 1
	MOVE_TYPE_ENPASSANT = 2
	MOVE_TYPE_CASTLING  = 3
)

// Directions
const (
	DIR_N  = 8
	DIR_S  = -8
	DIR_E  = 1
	DIR_W  = -1
	DIR_NE = 9
	DIR_NW = 7
	DIR_SE = -7
	DIR_SW = -9
)

// Castling rights
const (
	NO_CASTLING_RIGHTS = 0
	CASTLING_WHITE_OO  = 1
	CASTLING_WHITE_OOO = 2
	CASTLING_BLACK_OO  = 4
	CASTLING_BLACK_OOO = 8
)

// Ranks
const (
	RANK_1 = 0
	RANK_2 = 1
	RANK_3 = 2
	RANK_4 = 3
	RANK_5 = 4
	RANK_6 = 5
	RANK_7 = 6
	RANK_8 = 7
)

// Files
const (
	FILE_A = 0
	FILE_B = 1
	FILE_C = 2
	FILE_D = 3
	FILE_E = 4
	FILE_F = 5
	FILE_G = 6
	FILE_H = 7
)

// Squares
const (
	SQ_A1      = 0
	SQ_B1      = 1
	SQ_C1      = 2
	SQ_D1      = 3
	SQ_E1      = 4
	SQ_F1      = 5
	SQ_G1      = 6
	SQ_H1      = 7
	SQ_A2      = 8
	SQ_B2      = 9
	SQ_C2      = 10
	SQ_D2      = 11
	SQ_E2      = 12
	SQ_F2      = 13
	SQ_G2      = 14
	SQ_H2      = 15
	SQ_A3      = 16
	SQ_B3      = 17
	SQ_C3      = 18
	SQ_D3      = 19
	SQ_E3      = 20
	SQ_F3      = 21
	SQ_G3      = 22
	SQ_H3      = 23
	SQ_A4      = 24
	SQ_B4      = 25
	SQ_C4      = 26
	SQ_D4      = 27
	SQ_E4      = 28
	SQ_F4      = 29
	SQ_G4      = 30
	SQ_H4      = 31
	SQ_A5      = 32
	SQ_B5      = 33
	SQ_C5      = 34
	SQ_D5      = 35
	SQ_E5      = 36
	SQ_F5      = 37
	SQ_G5      = 38
	SQ_H5      = 39
	SQ_A6      = 40
	SQ_B6      = 41
	SQ_C6      = 42
	SQ_D6      = 43
	SQ_E6      = 44
	SQ_F6      = 45
	SQ_G6      = 46
	SQ_H6      = 47
	SQ_A7      = 48
	SQ_B7      = 49
	SQ_C7      = 50
	SQ_D7      = 51
	SQ_E7      = 52
	SQ_F7      = 53
	SQ_G7      = 54
	SQ_H7      = 55
	SQ_A8      = 56
	SQ_B8      = 57
	SQ_C8      = 58
	SQ_D8      = 59
	SQ_E8      = 60
	SQ_F8      = 61
	SQ_G8      = 62
	SQ_H8      = 63
	NR_SQUARES = 64
	NO_SQUARE  = 65
)

// Squares with string names
var SSM = map[string]int{
	"a1": 0,
	"b1": 1,
	"c1": 2,
	"d1": 3,
	"e1": 4,
	"f1": 5,
	"g1": 6,
	"h1": 7,
	"a2": 8,
	"b2": 9,
	"c2": 10,
	"d2": 11,
	"e2": 12,
	"f2": 13,
	"g2": 14,
	"h2": 15,
	"a3": 16,
	"b3": 17,
	"c3": 18,
	"d3": 19,
	"e3": 20,
	"f3": 21,
	"g3": 22,
	"h3": 23,
	"a4": 24,
	"b4": 25,
	"c4": 26,
	"d4": 27,
	"e4": 28,
	"f4": 29,
	"g4": 30,
	"h4": 31,
	"a5": 32,
	"b5": 33,
	"c5": 34,
	"d5": 35,
	"e5": 36,
	"f5": 37,
	"g5": 38,
	"h5": 39,
	"a6": 40,
	"b6": 41,
	"c6": 42,
	"d6": 43,
	"e6": 44,
	"f6": 45,
	"g6": 46,
	"h6": 47,
	"a7": 48,
	"b7": 49,
	"c7": 50,
	"d7": 51,
	"e7": 52,
	"f7": 53,
	"g7": 54,
	"h7": 55,
	"a8": 56,
	"b8": 57,
	"c8": 58,
	"d8": 59,
	"e8": 60,
	"f8": 61,
	"g8": 62,
	"h8": 63,
}
