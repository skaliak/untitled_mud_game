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

	fmt.Printf("\n%s\n", world.GetCurrentRoom().GetDesc())

	for {
		fmt.Print("\nnavigate: ")
		str, err := stdin.ReadString('\n')
		if err == nil {
			char := []byte(str)[0]
			switch char {
			case 'n':
				fmt.Print("North...")
				err = world.Navigate(navmap.North)
			case 's':
				fmt.Print("South...")
				err = world.Navigate(navmap.South)
			case 'e':
				fmt.Print("East...")
				err = world.Navigate(navmap.East)
			case 'w':
				fmt.Print("West...")
				err = world.Navigate(navmap.West)
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
