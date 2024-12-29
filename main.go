package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREENWIDTH             = 720
	SCREENHEIGHT            = 780
	GAMEPLAYWIDTH           = 600
	GAMEPLAYHEIGHT          = 600
	SQUARESIZE              = 30
	GAMEPLAY_BOUNDARY_WIDTH = 3
	GAMEPLAY_X_OFFSET       = (SCREENWIDTH - GAMEPLAYWIDTH) / 2
	GAMEPLAY_Y_OFFSET       = SCREENHEIGHT - GAMEPLAY_X_OFFSET - GAMEPLAYHEIGHT
)

var (
	running          = true
	BACKGROUND_COLOR = rl.NewColor(214, 234, 248, 255)
	GAMPLAY_BG_COLOR = rl.NewColor(253, 243, 228, 255)
	SNAKE_HEAD_COLOR = rl.NewColor(142, 243, 154, 255)
	SNAKE_BODY_COLOR = rl.NewColor(169, 223, 191, 255)
	FOOD_COLOR       = rl.NewColor(243, 156, 18, 255)
	SCORE_TEXT_COLOR = rl.NewColor(44, 62, 80, 255)
	food             Square
	snake            Snake
	UP               = rl.Vector2{X: 0, Y: -1}
	DOWN             = rl.Vector2{X: 0, Y: 1}
	LEFT             = rl.Vector2{X: -1, Y: 0}
	RIGHT            = rl.Vector2{X: 1, Y: 0}
	moveDirection    = RIGHT
	moveTimer        = 0.0
	moveInterval     = 0.5
	score            = 0
)

var (
// once = true
)

func drawScene() {
	scoreText := fmt.Sprintf("%d", score)
	rl.DrawText(scoreText, GAMEPLAY_X_OFFSET, 30, 30, SCORE_TEXT_COLOR)
	rl.DrawRectangle(GAMEPLAY_X_OFFSET, GAMEPLAY_Y_OFFSET, GAMEPLAYWIDTH, GAMEPLAYHEIGHT, GAMPLAY_BG_COLOR)
	rl.DrawRectangleLinesEx(
		rl.Rectangle{
			X:      GAMEPLAY_X_OFFSET - GAMEPLAY_BOUNDARY_WIDTH,
			Y:      GAMEPLAY_Y_OFFSET - GAMEPLAY_BOUNDARY_WIDTH,
			Width:  GAMEPLAYWIDTH + 2*GAMEPLAY_BOUNDARY_WIDTH,
			Height: GAMEPLAYHEIGHT + 2*GAMEPLAY_BOUNDARY_WIDTH,
		}, GAMEPLAY_BOUNDARY_WIDTH, SCORE_TEXT_COLOR)
	snake.Draw()
	food.Draw()
}

func input() {
	// if rl.IsKeyPressed(rl.KeyLeft) {
	// 	moveDirection = LEFT
	// } else if rl.IsKeyPressed(rl.KeyRight) {
	// 	moveDirection = RIGHT
	// } else if rl.IsKeyPressed(rl.KeyUp) {
	// 	moveDirection = UP
	// } else if rl.IsKeyPressed(rl.KeyDown) {
	// 	moveDirection = DOWN
	// }
	moveDirection = GetNextMoveDirection(snake, food)
}

func update() {
	running = !rl.WindowShouldClose()
	moveTimer += 0.1
	if moveTimer >= moveInterval {
		// if once {
		// 	grid := GetGrid(snake, food)
		// 	PrintBoard()
		// 	NextMoveBFS(grid)
		// 	once = false
		// 	running = false
		// }
		snake.Move(moveDirection)
		moveTimer = 0
		PrintBoard()
	}
	if snake.IsEatingFood(food) {

		grid := GetGrid(snake, food)
		emptyPosition := GetEmptyPosition(grid)
		food.MoveTo(emptyPosition)
		fmt.Println("empty position", emptyPosition)

		// food.Displace()

		snake.Expand()
		score += 1
	}

}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(BACKGROUND_COLOR)
	drawScene()
	rl.EndDrawing()
}

func init() {
	rl.InitWindow(SCREENWIDTH, SCREENHEIGHT, "snake")
	rl.SetTargetFPS(60)

	food = Square{X: GAMEPLAY_X_OFFSET, Y: GAMEPLAY_Y_OFFSET, Color: FOOD_COLOR}
	snake = Snake{Body: []Square{}, Size: 1}
	snake.Create(rl.Vector2{X: 15, Y: 15}, DOWN)
	food.Displace()
}

func main() {
	defer rl.CloseWindow()

	for running {
		input()
		update()
		render()
	}
}
