package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Square struct {
	X, Y  int
	Color rl.Color
	Dir   rl.Vector2
}

func (s Square) Draw() {
	rl.DrawRectangle(int32(s.X*SQUARESIZE-10), int32(s.Y*SQUARESIZE-10), SQUARESIZE+20, SQUARESIZE+20, rl.Black)
	rl.DrawRectangle(int32(s.X*SQUARESIZE), int32(s.Y*SQUARESIZE), SQUARESIZE, SQUARESIZE, s.Color)
}

func (s *Square) Displace() {
	randonX, randomY := rl.GetRandomValue(GAMEPLAY_X_OFFSET/SQUARESIZE, (GAMEPLAYWIDTH+GAMEPLAY_X_OFFSET)/SQUARESIZE-1),
		rl.GetRandomValue(GAMEPLAY_Y_OFFSET/SQUARESIZE, (GAMEPLAYHEIGHT+GAMEPLAY_Y_OFFSET)/SQUARESIZE-1)
	s.X = int(randonX)
	s.Y = int(randomY)
}

func (s *Square) MoveTo(coordinate Coordinate) {
	s.X = coordinate.X + GAMEPLAY_X_OFFSET/SQUARESIZE
	s.Y = coordinate.Y + GAMEPLAY_Y_OFFSET/SQUARESIZE
}

func (s *Square) Move(dir rl.Vector2) {
	s.Dir = dir
	s.X += int(dir.X)
	s.Y += int(dir.Y)

	if s.X < GAMEPLAY_X_OFFSET/SQUARESIZE {
		s.X = GAMEPLAYWIDTH/SQUARESIZE - 1 + GAMEPLAY_X_OFFSET/SQUARESIZE
	} else if s.X*SQUARESIZE >= GAMEPLAYWIDTH+GAMEPLAY_X_OFFSET {
		s.X = GAMEPLAY_X_OFFSET / SQUARESIZE
	}

	if s.Y < GAMEPLAY_Y_OFFSET/SQUARESIZE {
		s.Y = GAMEPLAYWIDTH/SQUARESIZE - 1 + GAMEPLAY_Y_OFFSET/SQUARESIZE
	} else if s.Y*SQUARESIZE >= GAMEPLAYWIDTH+GAMEPLAY_Y_OFFSET {
		s.Y = GAMEPLAY_Y_OFFSET / SQUARESIZE
	}

}
