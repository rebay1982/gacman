package pacman

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

type Pacman struct {
	PosX		int
	PosY		int
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


func PollInput() bool {

	quit := false

	switch kev := term.PollEvent(); kev.Type {
	case term.EventKey:
		switch kev.Key {
		case term.KeyArrowUp:
			reset()
			fmt.Println("Up")

		case term.KeyArrowDown:
			reset()
			fmt.Println("Down")

		case term.KeyArrowLeft:
			reset()
			fmt.Println("Left")

		case term.KeyArrowRight:
			reset()
			fmt.Println("Right")

		case term.KeyEsc:
			reset()
			quit = true

		default:
			reset()
			fmt.Println("Some unmapped key.")
		}

	case term.EventError:
		panic(kev.Err)
	}

	return quit
}

func reset() {
	term.Sync()	// Cosmetic purposes ?
}

func isPacmanHere(posX int, posY int) bool {

	return (posX == player.PosX && posY == player.PosY)
}
