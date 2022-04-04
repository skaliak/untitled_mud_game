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
		fmt.Print("navigate: ")
		char, _, err := stdin.ReadRune()
		if err != nil {
			switch char {
			case 'n':

			}
		}
	}
}
