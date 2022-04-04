package navmap

import "errors"

type Point struct {
	x int
	y int
}

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

var (
	n = Point{0, -1}
	s = Point{0, 1}
	e = Point{1, 0}
	w = Point{-1, 0}
)

func (a *Point) Add(b *Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func (p *Point) GetNeighbors(maxp *Point) []Point {
	neighbors := make([]Point, 4)
	dirs := []Point{n, s, e, w}
	for _, d := range dirs {
		neigh := p.Add(&d)
		if neigh.x >= 0 && neigh.x <= maxp.x && neigh.y >= 0 && neigh.y <= maxp.y {
			neighbors = append(neighbors, neigh)
		}
	}
	return neighbors
}

func (p *Point) GetNeighbor(dir Direction, maxp *Point) (*Point, error) {
	var d Point
	switch dir {
	case North:
		d = n
	case South:
		d = s
	case East:
		d = e
	case West:
		d = w
	}
	neigh := p.Add(&d)
	if neigh.x >= 0 && neigh.x <= maxp.x && neigh.y >= 0 && neigh.y <= maxp.y {
		return &neigh, nil
	} else {
		return nil, errors.New("out of bounds")
	}
}
