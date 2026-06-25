package service

import (
	"fmt"
	"simulation/internal/model"
	"time"
)

type Render func(gameMap *model.Map)

func RenderInConsole(gameMap *model.Map) {
	fmt.Print("\033[2J")
	fmt.Println()
	for range gameMap.GetWidth() {
		fmt.Print("_")
	}
	fmt.Println()
	for i := range gameMap.GetHeight() {
		fmt.Print("|")
		for j := range gameMap.GetWidth() {
			occ, exists := gameMap.PosToOcc[model.NewPosition(i, j)]
			if !exists {
				fmt.Print(" ")
			} else {
				switch occ.GetType() {
				case model.GRASS:
					fmt.Print("🌿")
				case model.ROCK:
					fmt.Print("🗻")
				case model.TREE:
					fmt.Print("🌳")
				case model.HERBIVORE:
					fmt.Print("🐒")
				case model.PREDATOR:
					fmt.Print("🐯")
				}
			}
		}
		fmt.Print("|\n")
	}
	for range gameMap.GetWidth() {
		fmt.Print("_")
	}
	time.Sleep(time.Duration(1) * time.Second)
}
