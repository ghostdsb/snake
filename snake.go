package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Snake struct {
	Body []Square
	Size int
}

func (sk *Snake) Create(pos, dir rl.Vector2) {
	for i := 0; i < sk.Size; i++ {
		square := Square{X: int(pos.X) - i*int(dir.X), Y: int(pos.Y) - i*int(dir.Y), Dir: dir, Color: SNAKE_HEAD_COLOR}
		sk.Body = append(sk.Body, square)
	}
}

func (sk Snake) Draw() {
	for _, square := range sk.Body {
		square.Draw()
	}
}

func (sk *Snake) Move(dir rl.Vector2) {
	size := len(sk.Body)
	for i := size - 1; i > 0; i-- {
		sk.Body[i].Dir = sk.Body[i-1].Dir
	}
	for i := size - 1; i > 0; i-- {
		sk.Body[i].Move(sk.Body[i].Dir)
	}
	sk.Body[0].Move(dir)
}

func (sk Snake) Head() Square {
	return sk.Body[0]
}

func (sk Snake) Tail() Square {
	length := len(sk.Body)
	return sk.Body[length-1]
}

func (sk Snake) BackTail() (int, int) {
	tailX, tailY := sk.Tail().X, sk.Tail().Y
	tailX -= int(sk.Tail().Dir.X)
	tailY -= int(sk.Tail().Dir.Y)
	return tailX, tailY
}

func (sk *Snake) Expand() {
	backTailX, backTailY := sk.BackTail()
	sk.Body = append(sk.Body, Square{X: backTailX, Y: backTailY, Color: SNAKE_BODY_COLOR, Dir: sk.Tail().Dir})
}

func (s Snake) IsEatingFood(food Square) bool {
	return snake.Head().X == food.X && snake.Head().Y == food.Y
}

func (s Snake) getAllSnakeCoordinates() []Coordinate {
	bodyCoord := make([]Coordinate, 0)
	for _, square := range s.Body {
		bodyCoord = append(bodyCoord, Coordinate{X: square.X, Y: square.Y})
	}
	return bodyCoord
}
