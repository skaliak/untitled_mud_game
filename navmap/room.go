package navmap

import "fmt"

type Room struct {
	loc   *Point
	nExit bool
	sExit bool
	eExit bool
	wExit bool
	desc  string
}

func (room *Room) GetDesc() string {
	return fmt.Sprintf("%d,%d %s", room.loc.x, room.loc.y, room.desc)
}

func NewRoom(location *Point, desc string) Room {
	return Room{loc: location, desc: desc}
}
