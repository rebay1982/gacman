package pacman

const mapWidth int = 5
const mapHeight int = 5

type Level struct {
	Height	int
	Width		int
	Data		[][]string
}

func (l Level) GetBlock(x int, y int) string {
	return l.Data[x][y]
}

var currentLevel = Level {5, 5,[][]string{
	{"*", "*", "*", "*", "*"},
	{"*", " ", " ", " ", "*"},
	{"*", " ", " ", " ", "*"},
	{"*", " ", " ", " ", "*"},
	{"*", "*", "*", "*", "*"}}}
