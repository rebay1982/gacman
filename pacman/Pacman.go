package pacman

import (
	"fmt"
)

type Pacman struct {
	PosX		int
	PosY		int
}

var player = Pacman{1, 1}

func Render() {

	// Clear the console screen before rendinering anything.
	fmt.Printf("\033[H\033[2J")

	for i := 0; i < currentLevel.Height; i++ {
		for j := 0; j < currentLevel.Width; j++ {

			if isPacmanHere(i, j) {
				fmt.Printf("P")
			} else {
				fmt.Printf(currentLevel.GetBlock(i, j))
			}
		}

		fmt.Printf("\n")
	}
}

func UpdateGame() {

}

func PollInput() {

	// No concept of this -- we will need to run a go routine that blocks and when
	// input is received, 
}

func isPacmanHere(posX int, posY int) bool {

	return (posX == player.PosX && posY == player.PosY)
}
