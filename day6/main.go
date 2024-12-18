package main

import (
	"fmt"
	"strings"
)

// Position represents a point in the grid
type Position struct {
	x, y int
}

// Direction represents the cardinal directions the guard can face
type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

// String returns a string representation of the direction
func (d Direction) String() string {
	return [...]string{"up", "right", "down", "left"}[d]
}

// parseMap converts the input string into a grid and finds the guard's starting position
// and direction. It returns the grid as a 2D slice of runes, the starting position,
// and the initial direction.
func parseMap(input string) ([][]rune, Position, Direction) {
	// Split input into lines and create the grid
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	var startPos Position
	var startDir Direction

	// Convert each line into a slice of runes and find the guard
	for y, line := range lines {
		grid[y] = []rune(strings.TrimSpace(line))
		for x, ch := range grid[y] {
			if ch == '^' {
				startPos = Position{x, y}
				startDir = Up
				grid[y][x] = '.' // Clear the starting position
			}
		}
	}

	return grid, startPos, startDir
}

// getNextPosition calculates the next position based on current position and direction
func getNextPosition(pos Position, dir Direction) Position {
	switch dir {
	case Up:
		return Position{pos.x, pos.y - 1}
	case Right:
		return Position{pos.x + 1, pos.y}
	case Down:
		return Position{pos.x, pos.y + 1}
	case Left:
		return Position{pos.x - 1, pos.y}
	}
	return pos // Should never happen
}

// turnRight returns the new direction after turning 90 degrees clockwise
func turnRight(dir Direction) Direction {
	return (dir + 1) % 4
}

// isObstacle checks if the given position contains an obstacle (#)
func isObstacle(pos Position, grid [][]rune) bool {
	return grid[pos.y][pos.x] == '#'
}

// willTouchBorder checks if the guard is at the border and facing outward
func willTouchBorder(pos Position, dir Direction, grid [][]rune) bool {
	height := len(grid)
	width := len(grid[0])

	switch dir {
	case Up:
		return pos.y == 0
	case Right:
		return pos.x == width-1
	case Down:
		return pos.y == height-1
	case Left:
		return pos.x == 0
	}
	return false
}

// trackGuardPath simulates the guard's movement and returns the number of
// distinct positions visited before leaving the mapped area
func trackGuardPath(input string) int {
	// Parse the input map and get initial state
	grid, currentPos, direction := parseMap(input)

	// Use a map to track visited positions
	visited := make(map[Position]bool)
	visited[currentPos] = true

	// Continue until the guard reaches a border while facing outward
	for {
		if willTouchBorder(currentPos, direction, grid) {
			break
		}

		// Calculate next position based on current direction
		nextPos := getNextPosition(currentPos, direction)

		// If there's an obstacle ahead, turn right
		if isObstacle(nextPos, grid) {
			direction = turnRight(direction)
		} else {
			// Move forward
			currentPos = nextPos
			visited[currentPos] = true
		}
	}

	return len(visited)
}

func main() {
	// Example input for testing
	input := `......#........#..........#.................##......................#.............#..................#............#............#..
.....#............................................#.......................#.......##.................#..#...................#.#...
......#...................#............................#.....#.........#................................................#.........
...........................................#.......#..........#..#......#.........#.#.#.#................#.##........##...........
........#.........................................#.......#.....#..........#...#....##....#..#.......##................#......#...
...................#..#.....................................................#.................................................#...
....#..................#..........#........#.............#................................................................#.......
......#..............#..................#...........#......#.....#...........................#..#.#.....................#.....#...
......#...#...............................................#.....#.........................#..........#............................
......#......#..........................................#..................................#....#..#................#..##...#.....
........................................................................................#.........................................
................#..........#..........................#.##......................#..........#......................................
.............#.............................#......#..............................#....................#...........#.....#...#.....
.......#.................#..................#..........................#...........................................#..............
.......#.....................................................#..........................................#....#.#.....#......#.....
...#......#............#........#......................#...........................#.#....#..#....................#.............##
...........................##..#..#..........#....................#................................................#......#.......
.....................#...#.......................#...#..#.....#................#...........................................#....#.
.................#................#.............#...#...........................................................#..#.#............
#.....................#....................##.......#...................#..........#..........................#...................
.#..........##......#...#......#................#.....................................#.#.....#..............#....................
...#...#..#..........................................................#..#..#......................................................
..................#.#..............#...#...............................#...#....#.#.........#.........................##..........
.......#.........................#..............................................#........#....#....#.........#....................
...........#......#......#...........#........#.............................##.................................................#..
............................#...................................#......#........................................#.................
......................................##...........................................#.............#..#...........#...#.....#.......
...................................................#.................................#............................................
..........................#..........................................#..#.....#.................#.#......#........................
..................#...........#.............................#....#.#..................#................................#..........
...................#....#..#.........................................................#....#.......................................
..#..........................................................##..................................#......................#....#.#..
#.........................#.....#...............................#........................#................#.......................
......................#......#...#..........#......#..............................#...#......#.................................#..
...............#.....#............................................................................#...............................
#.....................#........#.#........................................................#.......................#.............#.
.......#...#..........#........#...........#.................#..................#....................................##.....#.....
.#......#..................................#.......................................................................#..............
.................#.....#...........................................................#..............................................
.........#....................................^......#..................#..............................#..#.............#...#.....
.............................................................#.....#....................................#.........................
....................#.................#........#........................................................#..#............#........#
.....#..........................................................#...........................#.................#....##.............
..............#.........#.................#............................#.............#......................#.#............#......
.............#.................#.........................................#...#................#.............#.....#...............
....#.....##..#......#........#............................................#..........#.......#.................................##
..........................................................................#.......................................................
.......#..........#.............................................................................................##..#.............
............................#...............#...............................................................................#.....
....#......#...............................................................................................................#......
..................................#......................#.....#............................................#.....#........#......
...............................#..........................................#.......................................................
.....#.............#...#..............#.#............#....................................................................#.......
..............................................#...............................#......##.......#................#..........#.#.....
........#.....##...............................................#.........................#........................................
.....#.........................#.#........#.............................................#.............#..............##...........
........................................#..##.........................#............#...........#.............#....................
.....................#.#................#............................#.................#.....#..............................#.....
................#.....................................#..#....................#..................................#................
...........#...................#.................#..........................................................#.......#...........#.
...........................#..................................#......#..............#............#.........#....#................#
..................................................................................................................#...............
..............#....#.#............................................................................#.....#.........#........#......
..#...........#.................#.#......#................................#..........................#...........................#
............#...........##...............................#..................................................#.....................
..............##.#..................................................................................#.........#......#......#.....
................................................................#...............#....#....##................#.....................
.........#..#.....................#........................................................#......#................#..............
...#..........#..........#..................#.......................#........................#...#................#.........#.....
...................................#..#...............................#..........#................................................
.........................#........................................................................................................
......................#.........#....................................................##............................#.....#........
.#....#.....................#.........................................................#....#........#.............................
..........................................#....................................#........................................#.........
.................................................#.......................................#........................................
#..#...............#......................#....#.............................................#................#........##.........
.........#....#........##.........#.....#.........#...............................................#..................#....#...#...
..#............................................................................#...............................................#..
.......................................................#.........#......#..........................#.#............................
.........#.............#.................##.....#.....#.........#............#.........#..................#.....#.............#...
#..........#.........#...............................................................................................#............
.......................#....................#..................................................#..................................
.#................#.................................#........................................#...........#........................
.........#........................................................................................................................
..................................................................................#..........#.....................#..............
..#...............................#.......#.....................#......................#.....#..........................#.........
......#....#..........................................................#..........................#................................
...........................................##..................................................##..................#...#..........
...............#....#...........................#........................................#..............#....................##...
#......#......................................................#..........#.........................#.........#....................
...#...................................#......#..........................................#.........................#.............#
#.....#.........#.............#.......................................#...............................#...........................
...................#.#..........................................................................#...............#.................
.......#......##..........#....................................................#..................#..........#....................
.............................................#.##..............#......................#......#..#..................#..........#...
..##...................#...................................#.........................................#...........#................
...................................................#.....................................#.......................#...............#
..........#.......##...................#...............#...............................................................#..........
.......................#............#............##.............................................#......##.................#...#...
.........#.................................................................................#............#....................#....
.......##.........#..................................................................#................#...........................
................#.............................................................................................#...................
......................#.......#....#.................................................#......#.#...............................#...
.........#.........#.#.....##.#..................................##...................................................#...........
.#......................................................#.......#......##........#..........#.........#....#......................
........#.#......................#................................................................................................
..................#.............................#.................#..................#......................#......##.#...........
................##.............................................#....#..........##.#.#....#........................................
..............#...................#.......#............................#.#.....................................................#..
............#..............................................#....#.....................#...#.......................#.....#......#..
..............#........................................................#............................................#.............
.......#.....#.........#..............#..............................#.......................#..#................#....#..#........
.......................#......#........................................#.................#..........#.............................
...........#..................#.............................#............................................................#........
......#....#......#.....................#.....................#........#.....................#...#.....#.#...................#....
#..............................................##...#.....#...........#..........................................#................
..........#..........#.#........................................#..#....#.........................................................
..................#......................................#...#.......#............................................................
...#.............................#..................................................#...........#.................................
................................##.#.............#........................#...#...............................................#...
...#...................................#.......................#........................#............#...............#....#.#.....
..............................................................##.......#...............#.........#..................#.............
....#............................#..................................#..#..#...............#..............##.......................
#...#.................##...............................#...........#...........#..............#.#.........#........#...#..........
..........................#.....#....................#..............#.............#..........#....................#.........#.....
......#.......................................................#........#.....#..........................#...#.........#...........
...............................#............#.....#...........................#...............#........#..........................
............#..#......#...#..#.....................#...............#.........#...........................................#.....#..
.......#..#..................#............#...........#..............................................................#............
....#......#.##..#......##..........#.......#............#...............#....#....................................#..........#.#.`

	result := trackGuardPath(input)
	fmt.Printf("Number of distinct positions visited: %d\n", result)
}
