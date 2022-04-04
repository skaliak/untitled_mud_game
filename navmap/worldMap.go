package navmap

import "fmt"

type WorldMap struct {
	w       int
	h       int
	rooms   []Room
	roomIdx int
}

type Navigater interface {
	Navigate(d Direction, room *Room) (*Room, error)
	GetCurrentRoom() *Room
}

func MakeMap(w int, h int) WorldMap {
	m := WorldMap{w, h, make([]Room, w*h), 0}
	for i := 0; i < w*h; i++ {
		m.rooms[i] = Room{
			loc:  m.loc(i),
			desc: "room",
		}
		fmt.Printf("%d:%s\n", i, m.rooms[i].GetDesc())
	}

	return m
}

func (wm *WorldMap) Idx(x int, y int) int {
	return (y * (wm.w)) + x
}

func (wm *WorldMap) loc(i int) *Point {
	y := i / wm.w
	x := i % wm.w
	return &Point{x, y}
}

func (wm *WorldMap) GetCurrentRoom() *Room {
	return &wm.rooms[wm.roomIdx]
}

func (wm *WorldMap) Navigate(d Direction, room *Room) (*Room, error) {
	rm := wm.GetCurrentRoom()
	pt, err := rm.loc.GetNeighbor(d, wm.w, wm.h)
	if err != nil {
		fmt.Printf("%e\n", err)
		return nil, err
	}
	fmt.Printf("new point: %d,%d\n", pt.x, pt.y)
	wm.roomIdx = wm.Idx(pt.x, pt.y)
	fmt.Printf("new idx: %d\n", wm.roomIdx)
	return wm.GetCurrentRoom(), nil
}
