package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/skaliak/untitled_mud_game/navmap"
)

func main() {
	fmt.Print("hello world!")
	stdin := bufio.NewReader(os.Stdin)
	world := navmap.MakeMap(10, 10)
	var room *navmap.Room

	fmt.Printf("\n%s\n", world.GetCurrentRoom().GetDesc())

	for {
		fmt.Print("\nnavigate: ")
		str, err := stdin.ReadString('\n')
		if err == nil {
			char := []byte(str)[0]
			switch char {
			case 'n':
				fmt.Print("North...")
				room, err = world.Navigate(navmap.North, room)
			case 's':
				fmt.Print("South...")
				room, err = world.Navigate(navmap.South, room)
			case 'e':
				fmt.Print("East...")
				room, err = world.Navigate(navmap.East, room)
			case 'w':
				fmt.Print("West...")
				room, err = world.Navigate(navmap.West, room)
			default:
				fmt.Printf("invalid direction! [%c]", char)
				continue
			}
			if err != nil {
				fmt.Print(err)
			} else {
				fmt.Printf("%s\n", world.GetCurrentRoom().GetDesc())
			}
		} else {
			fmt.Print(err)
		}
	}
}
