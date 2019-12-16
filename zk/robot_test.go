package toys

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRobotMove(t *testing.T) {
	// Bootstrap - setup the board and the robot
	// NOTE: We only have to do this once for our test suite instead doing it once for each test case
	board := GameBoard{
		XLength: 3,
		YLength: 3,
	}

	robot := Robot{
		Position:  Position{X: 0, Y: 0},
		GameBoard: board,
	}

	t.Run("Move up 2", func(t *testing.T) {
		err := robot.Move(0, 2)
		require.NoError(t, err) // If there is an error then this test suite will stop here

		actualPosition := robot.Position
		expectedPosition := Position{X: 0, Y: 2}
		assert.Equal(t, expectedPosition, actualPosition)
	})

	// Sub-tests are executed in chronological order
	t.Run("Move right 2", func(t *testing.T) {
		err := robot.Move(2, 0)
		require.NoError(t, err)

		actualPosition := robot.Position
		expectedPosition := Position{X: 2, Y: 2}
		assert.Equal(t, expectedPosition, actualPosition)
	})

	t.Run("Move down and left at same time", func(t *testing.T) {
		err := robot.Move(-1, -1)
		require.NoError(t, err)

		actualPosition := robot.Position
		expectedPosition := Position{X: 1, Y: 1}
		assert.Equal(t, expectedPosition, actualPosition)

	})

	t.Run("Error if robot told to go out of bounds", func(t *testing.T) {
		//return error
		err := robot.Move(5, 0)
		expectedError := errors.New("robot would fall overboard")
		if assert.Error(t, err) {
			assert.Equal(t, expectedError, err)
		}
	})

	t.Run("Robot does not move if told to go out of bounds", func(t *testing.T) {
		//set up invalid move
		err := robot.Move(5, 0)
		fmt.Println(err)

		//don't move from position if invalid
		actualPosition := robot.Position
		expectedPosition := Position{X: 1, Y: 1}
		assert.Equal(t, expectedPosition, actualPosition)
	})

	// Hint: https://godoc.org/github.com/stretchr/testify/assert#Error
}
