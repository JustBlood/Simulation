package service

import (
	"simulation/internal/config"
	"simulation/internal/model"
)

type CountToRespawnFunc func(gameMap *model.Map) int

type RespawnOccupiers struct {
	SimulationSettings config.SimulationSettings
	CreaturesSettings  config.CreaturesSettings
	MinimalCount       int
	model.FactoryFunc
}

func (r RespawnOccupiers) RunAction(gameMap *model.Map) error {
	occ, _ := r.FactoryFunc(model.GenerateRandomParams(r.CreaturesSettings), model.NewPosition(-1, -1))
	countToRespawn := r.MinimalCount - gameMap.CountOccupiersOfType(occ.GetType())
	if countToRespawn > 0 {
		return SpawnOccupiers{countToRespawn, r.SimulationSettings, r.CreaturesSettings, r.FactoryFunc}.RunAction(gameMap)
	}
	return nil
}
