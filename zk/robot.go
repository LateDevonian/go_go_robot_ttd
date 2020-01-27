package toys

//based on zak's example

import (
	"errors"
	"fmt"
)

// GameBoard is a 2D grid which can have robots moved around on it
type GameBoard struct {
	XLength int
	YLength int
}

// Position represents the position of on the game board.
// This is relative where (0,0) is the bottom left corner of the board.
type Position struct {
	X int
	Y int
}

// Robot is a toy that can be moved around on the game board
type Robot struct {
	Position  Position
	GameBoard GameBoard
}

func (robot *Robot) Report() error {
	fmt.Println("i am reporting")
	return nil
}

// Move moves the robot around the game board, where deltaX is the number of tiles to move in the X direction
// and deltaY is the number of tiles to move in the Y direction.
func (robot *Robot) Move(deltaX, deltaY int) {

	robot.Position.X += deltaX
	robot.Position.Y += deltaY

	if !robot.isWithinBounds() {
		robot.Position.X -= deltaX
		robot.Position.Y -= deltaY
		return errors.New("Invalid move, must be on board")
	}

	return nil
}

func (robot *Robot) Place(deltaX, deltaY int) error {
	robot.Position.X = deltaX
	robot.Position.Y = deltaY

	if !robot.isWithinBounds() {
		return errors.New("Invalid move, must be on board")
	}
	return nil
}

func (robot *Robot) isWithinBounds() bool {
	position := robot.Position
	board := robot.GameBoard
	if position.X < 0 || position.X >= board.XLength {
		return false
	}
	if position.Y < 0 || position.Y >= board.YLength {
		return false
	}

	return true
}
