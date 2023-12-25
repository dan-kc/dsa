package main

import "fmt"

var directions = []position{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func main() {
	test := [][]byte{
		{'#', '#', '#', '#', '#', 'E', '#'},
		{'#', ' ', ' ', ' ', ' ', ' ', '#'},
		{'#', 'S', '#', '#', '#', '#', '#'},
	}

	path := solve(test)
	fmt.Println(path)
}

func solve(maze [][]byte) []position {
	startingPos := position{x: 1, y: 2}
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	var path []position
	walk(maze, startingPos, seen, &path)
	return path
}

type position struct {
	x, y int
}

func walk(maze [][]byte, curr position, seen [][]bool, path *[]position) bool {
	// are we off the map?
	if curr.x >= len(maze[0]) || curr.x < 0 {
		return false
	}
	if curr.y >= len(maze) || curr.y < 0 {
		return false
	}

	// are we at a wall?
	if maze[curr.y][curr.x] == '#' {
		return false
	}

	// have we arrived
	if maze[curr.y][curr.x] == 'E' {
		*path = append(*path, curr)
		return true
	}

	// have we already been here
	if seen[curr.y][curr.x] {
		return false
	}

	// pre
	seen[curr.y][curr.x] = true
	*path = append(*path, curr)

	// RECURSE TIME
	for i := 0; i < 4; i++ {
		if walk(maze, position{
			x: curr.x + directions[i].x,
			y: curr.y + directions[i].y,
		}, seen, path) {
			return true
		}
	}

	// post
	*path = (*path)[:len(*path)-1]

	return false
}
