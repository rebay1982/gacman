package pacman

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

const ST_FREE int = 0
const ST_PILL int = 10

type GameState struct {
	PosX			int
	PosY			int
	Score			int
	ObjState	[][]int
}

type UserInput struct {
	Up		bool
	Down	bool
	Left	bool
	Right	bool
	Quit	bool
}

var gameState GameState = *NewGameState()

// NewGameState Initialize a new pacman state.
func NewGameState() *GameState {
	return &GameState{1, 1, 0, [][]int{
	{ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE},
	{ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE},
	{ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE},
	{ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_PILL, ST_FREE},
	{ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE, ST_FREE}}}
}

// Init The init function
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

	fmt.Printf("Score: %d\n", gameState.Score)

	for y := 0; y < currentLevel.Height; y++ {
		for x := 0; x < currentLevel.Width; x++ {

			if isPacmanHere(x, y) {
				fmt.Printf("P")

			} else if isObjectPresent(x, y) {
				object := getObject(x, y)

				if object == ST_PILL {
					fmt.Printf(MapPill)
				}

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
		if !isCollision(gameState.PosX, gameState.PosY - 1) {
			gameState.PosY--
		}
	}

	if input.Down {
		if !isCollision(gameState.PosX, gameState.PosY + 1) {
			gameState.PosY++
		}
	}

	if input.Left {
		if !isCollision(gameState.PosX - 1, gameState.PosY) {
			gameState.PosX--
		}
	}

	if input.Right {
		if !isCollision(gameState.PosX + 1, gameState.PosY) {
			gameState.PosX++
		}
	}

	// Update points.
	switch object := getObject(gameState.PosX, gameState.PosY); object {
	case ST_PILL:
		gameState.Score += 10
		pickupObject(gameState.PosX, gameState.PosY)

	default:

	}
}

func isCollision(posX int, posY int) bool {
	if getMapBlock(posX, posY) == MapWall {
		return true
	}

	return false
}

func pickupObject(posX int, posY int) {
	gameState.ObjState[posY][posX] = ST_FREE
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
	return (posX == gameState.PosX && posY == gameState.PosY)
}

func isObjectPresent(posX int, posY int) bool {
	return gameState.ObjState[posY][posX] > ST_FREE
}

func getMapBlock(posX int, posY int) string {
	return currentLevel.GetBlock(posY, posX)
}

func getObject(posX int, posY int) int {
	return gameState.ObjState[posY][posX]
}


