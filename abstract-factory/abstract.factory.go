package abstractFactory

import "fmt"

func init() {
	fmt.Println("A Loading abstract factory.")
}

type RoomSide string

const (
	North RoomSide = "North"
	East  RoomSide = "East"
	South RoomSide = "South"
	West  RoomSide = "West"
)

type Wall interface {
	WallRoomSide
	Brash(color int)
}

type Maze interface {
	AddRoom(r Room)
}

type Room interface {
	SetSide(side RoomSide, wall WallRoomSide)
}

type Door interface {
	WallRoomSide
}

type WallRoomSide interface {
	BuildWall()
}

// AbstractFactory is it an interface for creating mazes
type AbstractFactory interface {
	MakeMaze() Maze
	MakeWall() Wall
	MakeRoom(i int) Room
	MakeDoor(r1 *Room, r2 *Room) Door
}
