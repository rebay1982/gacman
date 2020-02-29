package pacman

const mapWidth int = 10
const mapHeight int = 5

const MapWall string = "#"
const MapPill string = "*"
const mapFree string = " "
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
	{MapWall, mapFree, mapFree, mapFree, mapFree, mapFree, mapFree, mapFree, mapFree, MapWall},
	{MapWall, mapFree, mapFree, mapFree, MapWall, MapWall, MapWall, MapWall, mapFree, MapWall},
	{MapWall, mapFree, mapFree, mapFree, MapWall, mapFree, mapFree, mapFree, mapFree, MapWall},
	{MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall, MapWall}}}
