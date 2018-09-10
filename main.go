package main

import (
	"fmt"
	"go-patterns/abstract-factory"
)

func main() {
	commonBuilder := &abstractFactory.CommonMazeBuilder{}
	factory := &abstractFactory.MazeGame{}
	maze := factory.CreateMaze(commonBuilder)

	//a:= factory.Door.(*abstractFactory.CommonDoor)

	var commonDoor *abstractFactory.CommonDoor = factory.Door.(*abstractFactory.CommonDoor)

	//if commonDoor1, ok := factory.Door.(*abstractFactory.CommonDoor);ok{
	//	commonDoor = commonDoor1
	//	commonDoor.Title = "CV"
	//}

	commonDoor.Title = "PPC"

	fmt.Println(maze)



	bombedBuilder := abstractFactory.BombedMazeBuilder{}
	factory.CreateMaze(&bombedBuilder)
}
