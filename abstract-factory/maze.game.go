package abstractFactory

import "fmt"

func init() {
	fmt.Println("M Loading maze game.")
}

type MazeGame struct {

}

func (m *MazeGame) CreateMaze(f AbstractFactory) Maze {
	aMaze := f.MakeMaze()
	r1 := f.MakeRoom(1)
	r2 := f.MakeRoom(2)
	var door WallRoomSide = f.MakeDoor(&r1, &r2)

	aMaze.AddRoom(r1)
	aMaze.AddRoom(r2)

	wall := f.MakeWall()

	r1.SetSide(North, (wall).(WallRoomSide))
	r1.SetSide(East, door)
	r1.SetSide(South, (WallRoomSide)(f.MakeWall()))
	r1.SetSide(West, (WallRoomSide)(f.MakeWall()))

	r2.SetSide(North, (WallRoomSide)(f.MakeWall()))
	r2.SetSide(East, (WallRoomSide)(f.MakeWall()))
	r2.SetSide(South, (WallRoomSide)(f.MakeWall()))
	r2.SetSide(West, door)

	return aMaze
}
