package main

import "fmt"

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
		fmt.Println(top, start, end)
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
