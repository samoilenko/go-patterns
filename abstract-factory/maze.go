package abstract_factory

type CommonMaze struct {
	rooms []*Room
}

func (cm *CommonMaze) AddRoom(r *Room) {
	cm.rooms = append(cm.rooms, r)
}

type CommonWall struct {
	r1 *Room
	r2 *Room
}

func (cw *CommonWall) BuildWall() {

}
