package abstract_factory

type MazeGame struct {
}

func (m *MazeGame) CreateMaze(f AbstractFactory) {
	aMaze := f.MakeMaze()
	r1 := f.MakeRoom(1)
	r2 := f.MakeRoom(2)
	door := f.MakeDoor(&r1, &r2)

	doorAsRoomSide := WallRoomSide(door)

	aMaze.AddRoom(&r1)
	aMaze.AddRoom(&r2)

	r1.SetSide(North, f.MakeWall())
	r1.SetSide(East, &doorAsRoomSide)
	r1.SetSide(South, f.MakeWall())
	r1.SetSide(West, f.MakeWall())

	r2.SetSide(North, f.MakeWall())
	r2.SetSide(East, f.MakeWall())
	r2.SetSide(South, f.MakeWall())
	r2.SetSide(West, &door)
}
