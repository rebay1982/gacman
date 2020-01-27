package main

import (
	"fmt"
	"time"
	"github.com/rebay1982/gacman/pacman"
)

func main() {

	fmt.Printf("Welcome to Gacman\n")
	pacman.Init()
	pacman.Render()

mainLoop:
	for {
		quit := pacman.PollInput()	

		if quit {
			break mainLoop
		}

		// import "time"
		time.Sleep(200)
	}

	pacman.Destroy()
}


