package pacman

const mapWidth int = 10
const mapHeight int = 5

const MapWall string = "#"
const mapFree string = " "
const mapDott string = "."
const mapPill string = "*"

type Level struct {
	Height	int
	Width		int
	Data		[][]string
}

func (l Level) GetBlock(x int, y int) string {
	return l.Data[x][y]
}

var currentLevel = Level {5, 10,[][]string{
	{"#", "#", "#", "#", "#", "#", "#", "#", "#", "#"},
	{"#", " ", " ", " ", " ", " ", " ", " ", " ", "#"},
	{"#", " ", " ", " ", "#", "#", "#", "#", "#", "#"},
	{"#", " ", " ", " ", "#", " ", " ", " ", " ", " "},
	{"#", "#", "#", "#", "#", " ", " ", " ", " ", " "}}}
