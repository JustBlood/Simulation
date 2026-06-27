package service

import (
	"log/slog"
	"simulation/internal/config"
	"simulation/internal/model"
)

type MoveAction struct {
	CreatureSettings config.CreaturesSettings
	PathService      PathService
}

func (action MoveAction) RunAction(gameMap *model.Map) error {
	slog.Debug("MoveAction called", "positionsBefore", gameMap.PosToOcc)
	creatures := []model.Creature{}
	for i := range gameMap.PosToOcc {
		if creature, ok := gameMap.PosToOcc[i].(model.Creature); ok {
			creatures = append(creatures, creature)
		}
	}

	for i := range creatures {
		if creatures[i].IsDead() {
			slog.Info("Skip deleted creature on this cycle")
			continue
		}

		path, err := action.PathService.FindPath(gameMap, creatures[i].GetPos(), creatures[i].GetTargetSearchType())
		slog.Debug("Finded path for creature.", "creature", creatures[i], "path", path)
		if err != nil {
			slog.Error("Can't find correct path.", "creature", creatures[i], "error", err)
			continue
		}
		delete(gameMap.PosToOcc, creatures[i].GetPos())
		slog.Debug("Creature deleted from previous Position", "position", creatures[i].GetPos())
		pathIndex := min(len(path)-1, creatures[i].GetSpeed())
		slog.Debug("PathIndex calculated", "index", pathIndex)
		reachTheTarget := len(path)-1 == pathIndex
		if reachTheTarget {
			targetPos := path[pathIndex]
			occupier := gameMap.PosToOcc[targetPos]
			slog.Debug("Creature reach the target", "creature", creatures[i], "target", occupier)
			if targetCreature, ok := occupier.(model.Creature); ok {
				targetCreature.TakeDamage(creatures[i].(*model.Predator).GetDamage())
				if targetCreature.IsDead() {
					delete(gameMap.PosToOcc, targetCreature.GetPos())
					slog.Debug("Target has been deleted", "target", targetCreature)
					creatures[i].Heal(action.CreatureSettings.PredatorHeal)
				} else {
					pathIndex -= 1
				}
			} else {
				// Если травка-муравка, то сразу удаляем
				delete(gameMap.PosToOcc, targetPos)
				slog.Debug("Target has been deleted", "target", targetCreature)
				creatures[i].Heal(action.CreatureSettings.HerbivoreHeal)
			}
		}
		creatures[i].SetPosition(path[pathIndex])
		gameMap.PosToOcc[creatures[i].GetPos()] = creatures[i]
	}
	slog.Debug("MoveAction called", "positionsAfter", gameMap.PosToOcc)
	return nil
}
