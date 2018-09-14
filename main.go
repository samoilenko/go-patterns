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
	mazeGame := &abstractFactory.MazeGame{}

	commonBuilder := &abstractFactory.CommonMazeBuilder{}
	maze := mazeGame.CreateMaze(commonBuilder)

	//var commonDoor *abstractFactory.CommonDoor = mazeGame.Door.(*abstractFactory.CommonDoor)

	//if commonDoor1, ok := factory.Door.(*abstractFactory.CommonDoor);ok{
	//	commonDoor = commonDoor1
	//	commonDoor.Title = "Back door"
	//}

	//commonDoor.Title = "Entrance hall"

	fmt.Printf("%T \n", maze)

	commonMaze := maze.(*abstractFactory.CommonMaze)

	//doors := [2]*CommonDoor{}

	for _, room := range commonMaze.Rooms {
		fmt.Printf("===== %T\n", room)
		//commonRoom := room.(*CommonRoom)
		//for _, wall := range (*CommonRoom)(room).Walls {
		//	if commonDoor, ok := wall.(*CommonDoor); ok {
		//		doors[i] = commonDoor
		//	}
		//}
	}


	bombedBuilder := abstractFactory.BombedMazeBuilder{}
	//bombedMaze := mazeGame.CreateMaze(&bombedBuilder).(abstractFactory.BombedMaze)
	bombedMaze := mazeGame.CreateMaze(&bombedBuilder)

	fmt.Printf("%T \n", bombedMaze)
}
