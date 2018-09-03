package abstract_factory

type CommonMazeBuilder struct {
}

func (cmb *CommonMazeBuilder) MakeMaze() Maze {
	return new(CommonMaze)
}

func (cmb *CommonMazeBuilder) MakeWall() WallRoomSide {
	return &CommonWall{}
}
