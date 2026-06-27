package service

import (
	"fmt"
	"log/slog"
	"simulation/internal/config"
	"simulation/internal/model"
	"slices"
)

type MoveAction struct {
	CreatureSettings config.CreaturesSettings
}

func (action MoveAction) RunAction(gameMap *model.Map) error {
	slog.Debug("MoveAction called", "positionsBefore", gameMap.PosToOcc)
	// todo: пройтись по всем сущностям алгоритмом поиска
	// если это травоядное - найти траву
	// если это хищник - найти травоядное
	// ходят все в разнобой, всегда в разной последовательности, чтобы у отдельных особей не было преимущества
	//
	creatures := []model.Creature{}
	deletedCreatures := []model.Creature{} // fixme: можно перенести в model.Creature
	for _, v := range gameMap.PosToOcc {
		if creature, ok := v.(model.Creature); ok { // явно проверяем тип Creature
			creatures = append(creatures, creature)
		}
	}

	for _, creature := range creatures {
		if slices.Contains(deletedCreatures, creature) {
			slog.Info("Skip deleted creature on this cycle")
			continue
		}

		path, err := BreadthFirstSearchPath(gameMap, creature.GetPos(), creature.GetTargetSearchType())
		slog.Debug("Finded path for creature.", "creature", creature, "path", path)
		if err != nil {
			slog.Error("Can't find correct path.", "creature", creature, "error", err)
			continue // не осталось либо травоядных, либо травы, либо дойти до них невозможно
		}
		delete(gameMap.PosToOcc, creature.GetPos())
		slog.Debug("Creature deleted from previous Position", "position", creature.GetPos())
		pathIndex := min(len(path)-1, creature.GetSpeed())
		slog.Debug("PathIndex calculated", "index", pathIndex)
		if len(path)-1 == pathIndex {
			// достиг цели
			targetPos := path[pathIndex]
			occupier := gameMap.PosToOcc[targetPos]
			slog.Debug("Creature reach the target", "creature", creature, "target", occupier)
			if targetCreature, ok := occupier.(model.Creature); ok {
				if targetCreature.TakeDamage(creature.(*model.Predator).GetDamage()) <= 0 {
					// Вот и помер дед Максим
					delete(gameMap.PosToOcc, targetCreature.GetPos())
					deletedCreatures = append(deletedCreatures, targetCreature)
					slog.Debug("Target has been deleted", "target", targetCreature)
					creature.Heal(action.CreatureSettings.PredatorHeal)
				} else {
					// Существо не убили, клетку его не занимаем
					pathIndex -= 1
				}
			} else {
				// Если травка-муравка, то сразу удаляем
				delete(gameMap.PosToOcc, targetPos)
				slog.Debug("Target has been deleted", "target", targetCreature)
				creature.Heal(action.CreatureSettings.HerbivoreHeal)
			}
		}
		creature.SetPosition(path[pathIndex])
		gameMap.PosToOcc[creature.GetPos()] = creature
	}
	slog.Debug("MoveAction called", "positionsAfter", gameMap.PosToOcc)
	return nil
}

func BreadthFirstSearchPath(gameMap *model.Map, startPos model.Position, searchType model.OccupierType) ([]model.Position, error) {
	if !gameMap.IsInMap(startPos) {
		return nil, fmt.Errorf("start position is out of map")
	}

	queue := []model.Position{startPos}
	visited := make(map[model.Position]bool)
	cameFrom := make(map[model.Position]model.Position)

	// КРИТИЧНО: помечаем стартовую позицию как посещенную сразу,
	// чтобы она не добавилась в очередь сама от себя же.
	visited[startPos] = true

	// Оптимизация: используем индекс head вместо queue = queue[1:],
	// чтобы избежать O(N) сдвига элементов на каждом шаге.
	head := 0

	for head < len(queue) {
		currentPos := queue[head]
		head++

		// Если эта клетка занята нужным типом - мы нашли цель
		if occ, exists := gameMap.PosToOcc[currentPos]; exists && currentPos != startPos {
			if occ.GetType() == searchType {
				return reconstructPath(cameFrom, currentPos, startPos), nil
			}
		}

		// Добавляем следующие позиции для проверки
		// Передаем searchType, чтобы функция знала, что клетки с searchType - это цели, а не препятствия
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
		// 1. Пропускаем, если уже добавляли в очередь
		if visited[pos] {
			continue
		}

		// 2. Пропускаем, если вне карты
		if !gameMap.IsInMap(pos) {
			continue
		}

		// 3. Проверяем препятствия.
		// Если клетка занята, но это НЕ наша цель (searchType), то это препятствие.
		if occ, exists := gameMap.PosToOcc[pos]; exists {
			if occ.GetType() != searchType {
				continue // Непроходимое препятствие
			}
			// Если occ.GetType() == searchType, это наша цель, мы должны её добавить в очередь!
		}

		// Клетка валидна (пустая или наша цель)
		validPositions = append(validPositions, pos)
		visited[pos] = true // Помечаем СРАЗУ при добавлении, чтобы избежать дубликатов в очереди
	}

	return validPositions
}

func reconstructPath(cameFrom map[model.Position]model.Position, target, start model.Position) []model.Position {
	path := []model.Position{}

	for current := target; current != start; current = cameFrom[current] {
		path = append(path, current)
	}
	path = append(path, start)

	// Разворачиваем слайс
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	if len(path) == 0 {
		return []model.Position{start}
	}

	return path
}
