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
	OccupierType       model.OccupierType
	model.FactoryFunc
}

func (r RespawnOccupiers) RunAction(gameMap *model.Map) error {
	countToRespawn := r.MinimalCount - gameMap.CountOccupiersOfType(r.OccupierType)
	if countToRespawn > 0 {
		return SpawnOccupiers{countToRespawn, r.SimulationSettings, r.CreaturesSettings, r.FactoryFunc}.RunAction(gameMap)
	}
	return nil
}
