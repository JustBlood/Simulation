package service

import (
	"fmt"
	"simulation/internal/model"
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
)

const cellWidth = 2

type Renderer interface {
	Render(gameMap *model.Map)
	Close()
}

type ConsoleRenderer struct {
}

func NewConsoleRenderer() *ConsoleRenderer {
	fmt.Print("\033[2J\033[H")
	fmt.Print("\033[?25l")

	return &ConsoleRenderer{}
}

func (r *ConsoleRenderer) Close() {
	fmt.Print("\033[?25h") // Показываем курсор обратно
	fmt.Println()          // Переход на новую строку после карты
}

func (c *ConsoleRenderer) Render(gameMap *model.Map) {
	var sb strings.Builder
	totalHeight := gameMap.GetHeight() + 2
	sb.WriteString(fmt.Sprintf("\033[%dA", totalHeight+3))
	sb.WriteString("\033[H")

	sb.WriteString("┌")
	for w := 0; w < gameMap.GetWidth(); w++ {
		sb.WriteString(strings.Repeat("─", cellWidth))
	}
	sb.WriteString("┐\n")

	for h := 0; h < gameMap.GetHeight(); h++ {
		sb.WriteString("│")
		for w := 0; w < gameMap.GetWidth(); w++ {
			pos := model.NewPosition(h, w)

			var symbol string
			if occ, exists := gameMap.PosToOcc[pos]; exists {
				switch occ.GetType() {
				case model.GRASS:
					symbol = "🌿"
				case model.ROCK:
					symbol = "🗻"
				case model.TREE:
					symbol = "🌳"
				case model.HERBIVORE:
					symbol = "🐒"
				case model.PREDATOR:
					symbol = "🐯"
				}
			} else {
				symbol = "."
			}

			symbolWidth := runewidth.StringWidth(symbol)
			padding := cellWidth - symbolWidth

			sb.WriteString(symbol)
			if padding > 0 {
				sb.WriteString(strings.Repeat(" ", padding))
			}
		}
		sb.WriteString("│\n")
	}

	sb.WriteString("└")
	for w := 0; w < gameMap.GetWidth(); w++ {
		sb.WriteString(strings.Repeat("─", cellWidth))
	}
	sb.WriteString("┘\n")

	fmt.Print(sb.String())
	time.Sleep(time.Duration(1) * time.Second)
}
