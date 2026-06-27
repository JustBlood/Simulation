package service

import (
	"fmt"
	"simulation/internal/model"
)

type PathService interface {
	FindPath(gameMap *model.Map, startPos model.Position, searchType model.OccupierType) ([]model.Position, error)
}

type BreadthSearchPathService struct {
}

func (BreadthSearchPathService) FindPath(gameMap *model.Map, startPos model.Position, searchType model.OccupierType) ([]model.Position, error) {
	if !gameMap.IsInMap(startPos) {
		return nil, fmt.Errorf("start position is out of map")
	}

	queue := []model.Position{startPos}
	visited := make(map[model.Position]bool)
	cameFrom := make(map[model.Position]model.Position)
	visited[startPos] = true

	head := 0

	for head < len(queue) {
		currentPos := queue[head]
		head++

		if occ, exists := gameMap.PosToOcc[currentPos]; exists && currentPos != startPos {
			if occ.GetType() == searchType {
				return reconstructPath(cameFrom, currentPos, startPos), nil
			}
		}

		nextPositions := getNextPositionsToCheck(currentPos, visited, gameMap, searchType)
		for _, nextPos := range nextPositions {
			cameFrom[nextPos] = currentPos
			queue = append(queue, nextPos)
		}
	}
	return []model.Position{startPos}, fmt.Errorf("can't find path to nearest resource")
}

func getNextPositionsToCheck(currentPos model.Position, visited map[model.Position]bool, gameMap *model.Map, searchType model.OccupierType) []model.Position {
	nextPositions := []model.Position{
		model.NewPosition(currentPos.Height+1, currentPos.Width),
		model.NewPosition(currentPos.Height, currentPos.Width+1),
		model.NewPosition(currentPos.Height-1, currentPos.Width),
		model.NewPosition(currentPos.Height, currentPos.Width-1),
	}

	var validPositions []model.Position
	for _, pos := range nextPositions {
		if visited[pos] {
			continue
		}

		if !gameMap.IsInMap(pos) {
			continue
		}

		if occ, exists := gameMap.PosToOcc[pos]; exists {
			if occ.GetType() != searchType {
				continue
			}
		}

		validPositions = append(validPositions, pos)
		visited[pos] = true
	}

	return validPositions
}

func reconstructPath(cameFrom map[model.Position]model.Position, target, start model.Position) []model.Position {
	path := []model.Position{}

	for current := target; current != start; current = cameFrom[current] {
		path = append(path, current)
	}
	path = append(path, start)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	if len(path) == 0 {
		return []model.Position{start}
	}

	return path
}
