package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	gridSize = 10
)

type Position struct {
	x, y int
}

func (p *Position) move(dx, dy int) {
	p.x = (p.x + dx + gridSize) % gridSize
	p.y = (p.y + dy + gridSize) % gridSize
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wolf := Position{rand.Intn(gridSize), rand.Intn(gridSize)}
	hare := Position{rand.Intn(gridSize), rand.Intn(gridSize)}

	for wolf == hare { // Ensure wolf and hare do not start at the same position
		hare = Position{rand.Intn(gridSize), rand.Intn(gridSize)}
	}

	fmt.Println("Welcome to Wolf and Hare Game!")
	fmt.Println("The wolf (W) is chasing the hare (H) on a 10x10 grid.")
	fmt.Println("Use W/A/S/D to move the hare and escape the wolf.")

	gameOver := false

	for !gameOver {
		fmt.Println("\nCurrent Grid:")
		for i := 0; i < gridSize; i++ {
			for j := 0; j < gridSize; j++ {
				if wolf.x == i && wolf.y == j {
					fmt.Print("W ")
				} else if hare.x == i && hare.y == j {
					fmt.Print("H ")
				} else {
					fmt.Print(". ")
				}
			}
			fmt.Println()
		}

		if wolf == hare {
			fmt.Println("\nThe wolf caught the hare! Game Over!")
			gameOver = true
			break
		}

		var move string
		fmt.Print("Move (W/A/S/D): ")
		fmt.Scan(&move)

		switch move {
		case "W", "w":
			hare.move(-1, 0)
		case "A", "a":
			hare.move(0, -1)
		case "S", "s":
			hare.move(1, 0)
		case "D", "d":
			hare.move(0, 1)
		default:
			fmt.Println("Invalid move. Use W/A/S/D.")
			continue
		}

		if wolf.x < hare.x {
			wolf.move(1, 0)
		} else if wolf.x > hare.x {
			wolf.move(-1, 0)
		}

		if wolf.y < hare.y {
			wolf.move(0, 1)
		} else if wolf.y > hare.y {
			wolf.move(0, -1)
		}
	}
}
