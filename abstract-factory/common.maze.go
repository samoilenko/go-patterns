package abstractFactory

type CommonMazeBuilder struct {
}

func (cmb *CommonMazeBuilder) MakeMaze() Maze {
	return &CommonMaze{}
}

func (cmb *CommonMazeBuilder) MakeWall() Wall {
	return new(CommonWall)
}

func (cmb *CommonMazeBuilder) MakeDoor(r1 *Room, r2 *Room) Door {
	return &CommonDoor{Title: "Common door"}
}

func (cmb *CommonMazeBuilder) MakeRoom(i int) Room {
	return &CommonRoom{Walls: make(map[RoomSide]WallRoomSide)}
}

// ------------

type CommonMaze struct {
	Rooms []*Room
}

func (cm *CommonMaze) AddRoom(r *Room) {
	cm.Rooms = append(cm.Rooms, r)
}

//----------------------------------------
type CommonWall struct {
	r1 *Room
	r2 *Room
}

func (cw *CommonWall) BuildWall() {

}

func (cw *CommonWall) Brash(color int) {

}

//----------------------
type CommonDoor struct {
	Title string
}

func (cd *CommonDoor) BuildWall() {

}

//--------------------------------------
type CommonRoom struct {
	Walls map[RoomSide]WallRoomSide
}

func (cr *CommonRoom) SetSide(side RoomSide, wall WallRoomSide) {
	cr.Walls[side] = wall
}
