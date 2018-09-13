package abstractFactory

import (
	"fmt"
	"testing"
)

var factory *MazeGame

func init() {
	factory = &MazeGame{}
}

func TestCommonMaze(t *testing.T) {
	commonBuilder := &CommonMazeBuilder{}
	maze := factory.CreateMaze(commonBuilder)

	var commonDoor *CommonDoor = factory.Door.(*CommonDoor)

	//if commonDoor1, ok := factory.Door.(*abstractFactory.CommonDoor);ok{
	//	commonDoor = commonDoor1
	//	commonDoor.Title = "Back door"
	//}

	commonDoor.Title = "Entrance hall"

	if factory.Door.(*CommonDoor).Title != commonDoor.Title {
		t.Errorf("The door title has different title")
	}

	if _, ok := maze.(*CommonMaze); !ok {
		t.Errorf("Maze is not implement CommonMaze")
	}
}

func TestBombedMaze(t *testing.T) {
	bombedBuilder := BombedMazeBuilder{}
	//bombedMaze := factory.CreateMaze(&bombedBuilder).(abstractFactory.BombedMaze)
	bombedMaze := factory.CreateMaze(&bombedBuilder)

	if _, ok := bombedMaze.(*BombedMaze); !ok {
		t.Errorf("The maze is not a BombedMaze")
	}

	fmt.Printf("%T \n", bombedMaze)
}
