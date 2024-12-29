package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SquareType int

const (
	EmptyType SquareType = iota
	FoodType
	SnakeHeadType
	SnakeTailType
)

type Coordinate struct {
	X, Y       int
	squareType SquareType
}

func GetNextMoveDirection(snake Snake, food Square) rl.Vector2 {
	// index := rl.GetRandomValue(0, 3)
	// directions := []rl.Vector2{LEFT, RIGHT, UP, DOWN}
	grid := GetGrid(snake, food)
	switch NextMoveBFS(grid) {
	case 0:
		return UP
	case 1:
		return RIGHT
	case 2:
		return DOWN
	case 3:
		return LEFT
	}
	return RIGHT
}

func GetGrid(snake Snake, food Square) [][]Coordinate {
	snakeCoodinates := snake.getAllSnakeCoordinates()

	grid := make([][]Coordinate, 0)
	for i := 0; i < GAMEPLAYHEIGHT/SQUARESIZE; i++ {
		row := make([]Coordinate, 0)
		for j := 0; j < GAMEPLAYWIDTH/SQUARESIZE; j++ {
			row = append(row, Coordinate{X: j, Y: i, squareType: EmptyType})
		}
		grid = append(grid, row)
	}
	for i, square := range snakeCoodinates {
		var squareType SquareType
		if i == 0 {
			squareType = SnakeHeadType
		} else {
			squareType = SnakeTailType
		}
		grid[getAdjustedYCoordinate(square.Y)][getAdjustedXCoordinate(square.X)].squareType = squareType
	}
	// fmt.Println(food.Y)
	grid[getAdjustedYCoordinate(food.Y)][getAdjustedXCoordinate(food.X)].squareType = FoodType
	return grid
}

func GetEmptyPosition(grid [][]Coordinate) Coordinate {
	emptyPositions := []Coordinate{}
	for _, row := range grid {
		for _, col := range row {
			if col.squareType == EmptyType {
				emptyPositions = append(emptyPositions, col)
			}
		}
	}
	index := rl.GetRandomValue(0, int32(len(emptyPositions)-1))
	return emptyPositions[index]
}

func PrintBoard() {
	board := GetGrid(snake, food)
	for _, row := range board {
		for _, col := range row {
			fmt.Printf("%d ", col.squareType)
		}
		fmt.Println()
	}
}

func getAdjustedXCoordinate(x int) int {
	return x - GAMEPLAY_X_OFFSET/SQUARESIZE
}

func getAdjustedYCoordinate(y int) int {
	return y - GAMEPLAY_Y_OFFSET/SQUARESIZE
}
