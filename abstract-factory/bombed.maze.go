package abstractFactory

type BombedMazeBuilder struct {

}

func (bm BombedMazeBuilder) MakeMaze() Maze{
	return BombedMaze{Rooms: make([]*Room, 4)}
}

// ---------------

type BombedMaze struct {
	Rooms []*Room
}

func (bm BombedMaze) AddRoom(r *Room) {
	bm.Rooms = append(bm.Rooms, r)
}

func (bm BombedMaze) MakeWall() Wall{
	return new(BombedWall)
}

func (bm BombedMaze) MakeRoom(i int) Room {
	return BombedRoom{}
}

func (bm BombedMaze) MakeDoor(r1 *Room, r2 *Room) Door {
	return &CommonDoor{}
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

func (br BombedRoom) SetSide(side RoomSide, wall WallRoomSide){
	br.Walls[side] = wall
}
