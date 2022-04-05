package navmap

import "fmt"

type Room struct {
	loc     *Point
	blocked map[Direction]interface{}
	desc    string
}

func (room *Room) GetDesc() string {
	return fmt.Sprintf("%d,%d %s", room.loc.x, room.loc.y, room.desc)
}

func (room *Room) IsBlocked(d Direction) bool {
	if room.blocked == nil {
		return false
	}
	_, ok := room.blocked[d]
	return ok
}

func (room *Room) SetBlock(d Direction) {
	if room.blocked == nil {
		room.blocked = make(map[Direction]interface{})
	}
	var s struct{}
	room.blocked[d] = s
}

func (room *Room) DelBlock(d Direction) {
	if room.blocked != nil {
		delete(room.blocked, d)
	}
}

func NewRoom(location *Point, desc string) Room {
	return Room{
		loc:  location,
		desc: desc,
	}
}
