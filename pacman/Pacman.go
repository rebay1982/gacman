package pacman

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

type Pacman struct {
	PosX		int
	PosY		int
}

type UserInput struct {
	Up		bool
	Down	bool
	Left	bool
	Right	bool
	Quit	bool
}

var player = Pacman{1, 1}

func Init() {

	err := term.Init()

	if err != nil {
		panic(err)
	}
	// Setup input channel
	/*
	ch := make(chan string)
	go func(ch chan string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			s, err := reader.ReadString('\n')1


		}
	}
	*/
}

func Destroy() {

	term.Close()
}


func Render() {

	// Clear the console screen before rendinering anything.
	//fmt.Printf("\033[H\033[2J")
	term.Sync()

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
}

func isCollision(posX int, posY int) bool {
	if getMapBlock(posX, posY) == MapWall {
		return true
	}

	return false
}

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
