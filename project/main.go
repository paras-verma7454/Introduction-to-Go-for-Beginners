package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PieceType int
type Color int

const (
	Normal PieceType = iota
	Queen
)

const (
	Black Color = iota
	Red
	Empty
)

// piece struct
type Piece struct {
	Type  PieceType
	Color Color
}

// position struct
type Position struct {
	X, Y int
}

// move struct
type Move struct {
	Start    Position
	End      Position
	Captures Position
}

// board struct
type Board struct {
	Grid [8][8]*Piece
	Turn Color
}

func newBoard() *Board {
	board := &Board{
		Turn: Black,
	}
	for y := 0; y < 3; y++ {
		for x := 0; x < 8; x++ {
			if (x+y)%2 != 0 {
				board.Grid[y][x] = &Piece{
					Type:  Normal,
					Color: Black,
				}
			}
		}
	}

	for y := 5; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if (x+y)%2 != 0 {
				board.Grid[y][x] = &Piece{
					Type:  Normal,
					Color: Red,
				}
			}
		}
	}

	return board
}

func (b *Board) VisualizeBoard() {
	fmt.Println("\n  +---+---+---+---+---+---+---+---+")
	for y := 0; y < 8; y++ {
		fmt.Printf("%d |", y)
		for x := 0; x < 8; x++ {
			// This space is empty
			if b.Grid[y][x] == nil {
				if (x+y)%2 == 0 {
					fmt.Printf(" . |")
				} else {
					fmt.Printf("   |")
				}
			} else { // This space is NOT empty
				piece := b.Grid[y][x]

				if piece.Color == Black {
					if piece.Type == Queen {
						fmt.Print(" B |") // Black Queen
					} else {
						fmt.Print(" b |") // Black Normal
					}
				} else {
					if piece.Type == Queen {
						fmt.Print(" R |") // Red Queen
					} else {
						fmt.Print(" r |") // Red Normal
					}
				}
			}
		}
		fmt.Println("\n  +---+---+---+---+---+---+---+---+")
	}
}

func (b *Board) IsValidMove(move Move) (bool, string) {
	if !IsInBounds(move.Start) || !IsInBounds(move.End) {
		return false, "Move is Out of bounds"
	}

	piece := b.Grid[move.Start.Y][move.Start.X]
	if piece == nil {
		return false, "No piece at start position"
	}

	if b.Grid[move.End.Y][move.End.X] != nil {
		return false, "Destination square is occupied"
	}

	if piece.Color != b.Turn {
		return false, fmt.Sprintf("It's %s's turn", colorToString(b.Turn))
	}

	dx := move.End.X - move.Start.X
	dy := move.End.Y - move.Start.Y

	if abs(dx) != abs(dy) {
		return false, "Move must be diagonal"
	}

	if piece.Type == Normal {
		if piece.Color == Black && dy <= 0 {
			return false, "Normal Black pieces can only move down"
		}
		if piece.Color == Red && dy >= 0 {
			return false, "Normal Red pieces can only move up"
		}
	}

	if abs(dx) == 1 && abs(dy) == 1 {
		return true, ""
	} else if abs(dx) == 2 && abs(dy) == 2 {
		captureX := move.Start.X + dx/2
		captureY := move.Start.Y + dy/2
		capturedPiece := b.Grid[captureY][captureX]
		if capturedPiece == nil {
			return false, "No piece to capture"
		}
		if capturedPiece.Color == piece.Color {
			return false, "Cannot capture your own piece"
		}
		return true, "" // Valid capture move
	}

	return false, "Invalid move distance"
}

func (b *Board) MakeMove(move Move) bool {

	valid, reason := b.IsValidMove(move)
	if !valid {
		fmt.Printf("Invalid move: %s\n", reason)
		return false
	}
	// Execute the move
	piece := b.Grid[move.Start.Y][move.Start.X]
	b.Grid[move.End.Y][move.End.X] = piece
	b.Grid[move.Start.Y][move.Start.X] = nil

	// Handle the capture
	if abs(move.End.X-move.Start.X) == 2 {
		captureX := (move.Start.X + move.End.X) / 2
		captureY := (move.Start.Y + move.End.Y) / 2
		capturePiece := b.Grid[captureY][captureX]
		b.Grid[captureY][captureX] = nil
		fmt.Printf("Captured %s piece at %d, %d\n", colorToString(capturePiece.Color), captureX, captureY)
	}

	// Handle promotion
	if piece.Type == Normal {
		if (piece.Color == Black && move.End.Y == 7) || (piece.Color == Red && move.End.Y == 0) {
			piece.Type = Queen
			fmt.Printf("%s piece Promoted to Queen at %d, %d\n", colorToString(piece.Color), move.End.X, move.End.Y)
		}
	}

	// Switch the turn
	b.Turn = opponents(b.Turn)
	return true
}

func IsInBounds(position Position) bool {
	return position.X >= 0 && position.X < 8 && position.Y >= 0 && position.Y < 8
}

func colorToString(c Color) string {
	if c == Black {
		return "Black"
	}
	return "Red"
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func opponents(c Color) Color {
	if c == Black {
		return Red
	}
	return Black
}

func GameLoop(board *Board) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Checkers Game")
	fmt.Println("Instructions: ")
	fmt.Println("-Black pieces are represented by b and Black Queen pieces are represented by B")
	fmt.Println("-Red pieces are represented by r and Red Queen pieces are represented by R")
	fmt.Println("-Black pieces can only move down and Red pieces can only move up")
	fmt.Println("-Enter the coordinates in the format startX startY endX endY")
	fmt.Println("-For example, '2 2 3 3' moves from (2,2) to (3,3) ")

	for {
		board.VisualizeBoard()
		fmt.Printf("Enter move (start X and start Y , end X and end Y): ")
		if !scanner.Scan() {
			return
		}

		input := scanner.Text()
		coords := strings.Split(input, " ")
		if len(coords) != 4 {
			fmt.Println("Invalid input format. Please use startX and startY, endX and endY")
			continue
		}

		numbers := make([]int, 4)
		valid := true

		for i, coords := range coords {
			num, err := strconv.Atoi(coords)
			if err != nil || num < 0 || num > 7 {
				fmt.Println("Invalid input Please enter numbers between 0 and 7")
				valid = false
				break
			}
			numbers[i] = num
		}

		if !valid {
			continue
		}

		move := Move{
			Start: Position{X: numbers[0], Y: numbers[1]},
			End:   Position{X: numbers[2], Y: numbers[3]},
		}

		board.MakeMove(move)
	}
}
func main() {
	board := newBoard()
	GameLoop(board)
}
