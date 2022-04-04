package navmap

type WorldMap struct {
	endCorner Point
	rooms     []Room
	roomIdx   int
}

func MakeMap(w int, h int) WorldMap {
	m := WorldMap{Point{w, h}, make([]Room, w*h), 0}
	for i := 0; i < w*h; i++ {
		m.rooms[i] = Room{
			loc:  m.loc(i),
			desc: "room",
		}
	}

	return m
}

func (wm *WorldMap) idx(x int, y int) int {
	return (y*wm.endCorner.x + 1) + x
}

func (wm *WorldMap) loc(i int) *Point {
	y := i / wm.endCorner.x
	x := i % wm.endCorner.x
	return &Point{x, y}
}

func (wm *WorldMap) GetCurrentRoom() *Room {
	return &wm.rooms[wm.roomIdx]
}

func (wm *WorldMap) Navigate(d Direction) error {
	rm := wm.GetCurrentRoom()
	pt, err := rm.loc.GetNeighbor(d, &wm.endCorner)
	if err != nil {
		return err
	}
	wm.roomIdx = wm.idx(pt.x, pt.y)
	return nil
}
