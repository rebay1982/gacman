package main

import (
	"fmt"
	"time"
	"github.com/rebay1982/gacman/pacman"
)

func main() {

	fmt.Printf("Welcome to Gacman\n")
	pacman.Init()

mainLoop:
	for {
		pacman.Render()

		input := pacman.PollInput()

		if input.Quit {
			break mainLoop
		}

		pacman.UpdateGame(input)

		// import "time"
		time.Sleep(200)
	}

	pacman.Destroy()
}


