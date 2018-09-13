package main

import (
	"fmt"
	"go-patterns/abstract-factory"
)

func init() {
	// init functions calls in the alphabetic orders
	// https://medium.com/golangspec/init-functions-in-go-eac191b3860a
	fmt.Println("M Starting the program.")
}

func main() {
	factory := &abstractFactory.MazeGame{}

	commonBuilder := &abstractFactory.CommonMazeBuilder{}
	maze := factory.CreateMaze(commonBuilder)

	var commonDoor *abstractFactory.CommonDoor = factory.Door.(*abstractFactory.CommonDoor)

	//if commonDoor1, ok := factory.Door.(*abstractFactory.CommonDoor);ok{
	//	commonDoor = commonDoor1
	//	commonDoor.Title = "Back door"
	//}

	commonDoor.Title = "Entrance hall"

	fmt.Printf("%T \n", maze)

	bombedBuilder := abstractFactory.BombedMazeBuilder{}
	//bombedMaze := factory.CreateMaze(&bombedBuilder).(abstractFactory.BombedMaze)
	bombedMaze := factory.CreateMaze(&bombedBuilder)

	fmt.Printf("%T \n", bombedMaze)
}
