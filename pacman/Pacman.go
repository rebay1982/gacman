package pacman

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

type Pacman struct {
	PosX		int
	PosY		int
	Score		int
}

type UserInput struct {
	Up		bool
	Down	bool
	Left	bool
	Right	bool
	Quit	bool
}

var player = Pacman{1, 1, 0}

func Init() {

	err := term.Init()

	if err != nil {
		panic(err)
	}
}

func Destroy() {

	term.Close()
}


func Render() {

	// Cleans up the terminal.
	term.Sync()

	fmt.Printf("Score: %d\n", player.Score)

	for y := 0; y < currentLevel.Height; y++ {
		for x := 0; x < currentLevel.Width; x++ {

			if isPacmanHere(x, y) {
				fmt.Printf("P")
			} else {
				fmt.Printf(getMapBlock(x, y))
			}
		}

		fmt.Printf("\n")
	}
}

func UpdateGame(input UserInput) {

	// Move player around depending on input.
	if input.Up {
		if !isCollision(player.PosX, player.PosY - 1) {
			player.PosY--
		}
	}

	if input.Down {
		if !isCollision(player.PosX, player.PosY + 1) {
			player.PosY++
		}
	}

	if input.Left {
		if !isCollision(player.PosX - 1, player.PosY) {
			player.PosX--
		}
	}

	if input.Right {
		if !isCollision(player.PosX + 1, player.PosY) {
			player.PosX++
		}
	}

	// Update points.
	if isPoints, points := isPoints(player.PosX, player.PosY); isPoints {
		player.Score += points
	}
}

func isCollision(posX int, posY int) bool {
	if getMapBlock(posX, posY) == MapWall {
		return true
	}

	return false
}

func isPoints(posX int, posY int) (bool, int) {

	if getMapBlock(posX, posY) == MapPill {
		return true, 10
	}

	return false, 0
}

// PollInput Polls for user input, returns a user input structure.
func PollInput() UserInput {

	var input UserInput

	switch kev := term.PollEvent(); kev.Type {
	case term.EventKey:
		switch kev.Key {
		case term.KeyArrowUp:
			input.Up = true

		case term.KeyArrowDown:
			input.Down = true

		case term.KeyArrowLeft:
			input.Left = true

		case term.KeyArrowRight:
			input.Right = true

		case term.KeyEsc:
			input.Quit = true

		default:

		}

	case term.EventError:
		panic(kev.Err)
	}

	return input
}

func isPacmanHere(posX int, posY int) bool {
	return (posX == player.PosX && posY == player.PosY)
}

func getMapBlock(posX int, posY int) string {
	return currentLevel.GetBlock(posY, posX)
}
