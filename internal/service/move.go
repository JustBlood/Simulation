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
	slog.Debug("MoveAction called", "gameMap", gameMap)
	creatures := gameMap.GetAllCreations()

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

		pathIndex := min(len(path)-1, creatures[i].Speed())
		slog.Debug("PathIndex calculated", "index", pathIndex)
		targetReached := len(path)-1 == pathIndex
		var killed bool
		if targetReached {
			killed, err = action.attackTheTarget(path, pathIndex, gameMap, creatures[i])
			if err != nil {
				return err
			}

			if !killed {
				pathIndex -= 1
			}
		}

		slog.Debug("Creature deleted from previous Position", "position", creatures[i].GetPos())
		err = gameMap.ReplaceOccupierPosition(creatures[i], path[pathIndex])
		if err != nil {
			return err
		}
	}
	slog.Debug("MoveAction called", "gameMap", gameMap)
	return nil
}

func (action MoveAction) attackTheTarget(path []model.Position, pathIndex int, gameMap *model.Map, creature model.Creature) (bool, error) {
	targetPos := path[pathIndex]
	occupier, err := gameMap.GetOccupierByPosition(targetPos)
	if err != nil {
		return false, err
	}
	slog.Debug("Creature reach the target", "creature", creature, "target", occupier)

	targetCreature, targetIsCreature := occupier.(model.Creature)
	targetKilled := true
	if targetIsCreature {
		slog.Debug("Target has been attacked", "target", targetCreature, "attacked", creature)
		targetCreature.TakeDamage(creature.(*model.Predator).GetDamage())
		targetKilled = targetCreature.IsDead()
	}

	if targetKilled {
		action.afterKill(targetPos, creature, gameMap)
	}

	return targetKilled, nil
}

func (action MoveAction) afterKill(targetPos model.Position, attacked model.Creature, gameMap *model.Map) {
	gameMap.DeleteOccupierByPosition(targetPos)
	attacked.Heal(action.CreatureSettings.PredatorHeal)
	slog.Debug("Target has been deleted (killed)", "target", targetPos, "attacked", attacked)
}
