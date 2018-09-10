package abstractFactory

type BombedMazeBuilder struct {
}

func (bmb BombedMazeBuilder) MakeMaze() Maze {
	return BombedMaze{Rooms: make([]*Room, 4)}
}

func (bmb BombedMazeBuilder) MakeWall() Wall {
	return new(BombedWall)
}

func (bmb BombedMazeBuilder) MakeRoom(i int) Room {
	return BombedRoom{Walls: make(map[RoomSide]WallRoomSide, 4)}
}

func (bmb BombedMazeBuilder) MakeDoor(r1 *Room, r2 *Room) Door {
	return &CommonDoor{}
}

// ---------------

type BombedMaze struct {
	Rooms []*Room
}

func (bm BombedMaze) AddRoom(r *Room) {
	bm.Rooms = append(bm.Rooms, r)

}

// -----Wall----
type BombedWall struct {
}

func (bw BombedWall) BuildWall() {

}

func (bw BombedWall) Brash(color int) {

}

// --------Room------
type BombedRoom struct {
	Walls map[RoomSide]WallRoomSide
}

func (br BombedRoom) SetSide(side RoomSide, wall WallRoomSide) {
	br.Walls[side] = wall
}
