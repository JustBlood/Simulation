package service

import (
	"fmt"
	"math/rand/v2"
	"simulation/internal/config"
	"simulation/internal/model"
)

type SpawnOccupiers struct {
	Count              int
	SimulationSettings config.SimulationSettings
	CreaturesSettings  config.CreaturesSettings
	model.FactoryFunc
}

func (s SpawnOccupiers) RunAction(gameMap *model.Map) error {
	for range s.Count {
		for {
			pos := model.NewPosition(rand.IntN(s.SimulationSettings.MapHeight), rand.IntN(s.SimulationSettings.MapWidth))
			if gameMap.IsEmpty(pos) {
				occ, err := s.FactoryFunc(model.GenerateRandomParams(s.CreaturesSettings), pos)
				if err != nil {
					return fmt.Errorf("Can't create Occupier: %w", err)
				}
				err = gameMap.Set(pos, occ)
				if err != nil {
					return fmt.Errorf("Can't set occupier to pos %w", err)
				}
				break
			}
		}
	}
	return nil
}
