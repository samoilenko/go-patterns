package abstract_factory

type RoomSide string

const (
	North RoomSide = "North"
	East RoomSide = "East"
	South RoomSide = "South"
	West RoomSide = "West"
)

type Wall interface {
	Build()
}

type Maze interface {
	AddRoom(r *Room)
}

type Room interface {
	SetSide(side RoomSide, wall *WallRoomSide)
}

type Door interface {

}

type WallRoomSide interface {

}

// AbstractFactory is it an interface for creating mazes
type AbstractFactory interface {
	MakeMaze()Maze
	MakeWall()Wall
	MakeRoom(i int)Room
	MakeDoor(r1 *Room, r2 *Room)Door
}



type MazeGame struct {

}

func (m *MazeGame) CreateMaze(f AbstractFactory){
	aMaze:= f.MakeMaze()
	r1 := f.MakeRoom(1)
	r2 := f.MakeRoom(2)
	door := f.MakeDoor(&r1, &r2)

	aMaze.AddRoom(&r1)
	aMaze.AddRoom(&r2)

	r1.SetSide(North, f.MakeWall())
	r1.SetSide(East, door)
	r1.SetSide(South, f.MakeWall())
	r1.SetSide(West, f.MakeWall())

	r2.SetSide(North, f.MakeWall())
	r2.SetSide(East, f.MakeWall())
	r2.SetSide(South, f.MakeWall())
	r2.SetSide(West, door)
}
