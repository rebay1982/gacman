package pacman

const mapWidth int = 10
const mapHeight int = 5

const MapWall string = "#"
const MapPill string = "*"
const MapFree string = " "
const mapDott string = "."

type Level struct {
	Height	int
	Width		int
	Data		[][]string
}

func (l Level) GetBlock(x int, y int) string {
	return l.Data[x][y]
}

var currentLevel = Level {5, 10,[][]string{
	{MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall},
	{MapWall, MapFree, MapFree, MapFree, MapFree, MapFree, MapFree, MapFree, MapFree, MapWall},
	{MapWall, MapFree, MapFree, MapFree, MapWall, MapWall, MapWall, MapWall, MapFree, MapWall},
	{MapWall, MapFree, MapFree, MapFree, MapWall, MapFree, MapFree, MapFree, MapFree, MapWall},
	{MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall}}}



