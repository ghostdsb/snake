package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func NextMoveBFS(grid [][]Coordinate) int {
	var start, end Coordinate
	for r, row := range grid {
		for c, col := range row {
			if col.squareType == SnakeHeadType {
				start = Coordinate{X: c, Y: r, squareType: SnakeHeadType}
			}
			if col.squareType == FoodType {
				end = Coordinate{X: c, Y: r, squareType: FoodType}
			}
		}
	}
	q := []Coordinate{start}
	visited := make(map[Coordinate]bool)
	parent := make(map[Coordinate]Coordinate)
	parent[start] = Coordinate{}
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		top := q[0]
		fmt.Println("X", top, start, end)
		q = q[1:]
		visited[top] = true
		if top.squareType == FoodType {
			fmt.Println("last", parent[top])
			break
		}
		for _, d := range dir {
			nx := top.X + d[0]
			ny := top.Y + d[1]
			if nx >= 0 && ny >= 0 && nx < len(grid[0]) && ny < len(grid) {
				if grid[ny][nx].squareType == EmptyType || grid[ny][nx].squareType == FoodType {
					if _, ok := visited[grid[ny][nx]]; !ok {
						visited[grid[ny][nx]] = true
						parent[grid[ny][nx]] = top
						q = append(q, grid[ny][nx])
					}
				}
			}
		}
	}
	curr := end
	var moveTo Coordinate
	for curr != start {
		moveTo = curr
		if curr == parent[curr] {
			best, direction := 0, []int{-1, -1}
			for _, d := range dir {
				nx := start.X + d[0]
				ny := start.Y + d[1]
				if nx >= 0 && ny >= 0 && nx < len(grid[0]) && ny < len(grid) {
					if grid[ny][nx].squareType == EmptyType || grid[ny][nx].squareType == FoodType {
						fill := floodFill(grid, grid[ny][nx])
						if fill > best {
							direction = d
							best = fill
						}
						// if d[0] == 1 && d[1] == 0 {
						// 	return 1
						// }
						// if d[0] == -1 && d[1] == 0 {
						// 	return 3
						// }
						// if d[0] == 0 && d[1] == 1 {
						// 	return 2
						// }
						// if d[0] == 0 && d[1] == -1 {
						// 	return 0
						// }
					}
				}
			}
			if direction[0] == 1 && direction[1] == 0 {
				return 1
			}
			if direction[0] == -1 && direction[1] == 0 {
				return 3
			}
			if direction[0] == 0 && direction[1] == 1 {
				return 2
			}
			if direction[0] == 0 && direction[1] == -1 {
				return 0
			}

			return int(rl.GetRandomValue(0, 3))
		}
		curr = parent[curr]
		fmt.Println(curr)
	}
	if moveTo.X > start.X {
		fmt.Println("GO RIGHT")
		return 1
	}
	if moveTo.X < start.X {
		fmt.Println("GO LEFT")
		return 3
	}
	if moveTo.Y > start.Y {
		fmt.Println("GO DOWN")
		return 2
	}
	if moveTo.Y < start.Y {
		fmt.Println("GO UP")
		return 0
	}
	return -1
}

func floodFill(grid [][]Coordinate, start Coordinate) int {
	count := 0
	q := []Coordinate{start}
	visited := make(map[Coordinate]bool)
	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		visited[top] = true
		for _, d := range dir {
			nx := top.X + d[0]
			ny := top.Y + d[1]
			if nx >= 0 && ny >= 0 && nx < len(grid[0]) && ny < len(grid) {
				if grid[ny][nx].squareType == EmptyType || grid[ny][nx].squareType == FoodType {
					if _, ok := visited[grid[ny][nx]]; !ok {
						count += 1
						visited[grid[ny][nx]] = true
						q = append(q, grid[ny][nx])
					}
				}
			}
		}
	}
	return count
}
