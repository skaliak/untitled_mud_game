package mobs

import "github.com/skaliak/untitled_mud_game/navmap"

type Mover interface {
	Move(d navmap.Direction) error
}

type Mob struct {
	room     *navmap.Room
	worldMap *navmap.WorldMap
}
