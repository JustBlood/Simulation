package main

import (
	"errors"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"simulation/internal/config"
	"simulation/internal/game"
	"simulation/internal/model"
	"simulation/internal/service"
)

func main() {
	logFile := config.InitLogger()
	defer logFile.Close()

	settings, err := config.LoadGlobalSettings("../config.json")
	if err != nil {
		fmt.Print(err)
		return
	}
	needToCreateOccupiers := settings.MinOccupiers + rand.IntN(settings.MaxOccupiers-settings.MinOccupiers)

	avg := needToCreateOccupiers / len(model.AllEntityTypes)
	counts, err := DistributeCreatures(needToCreateOccupiers, len(model.AllEntityTypes), avg*3/4, avg*5/4, 2)

	if err != nil {
		slog.Error("Error in initialization", "error", err)
		return
	}

	initActions := []service.Action{
		service.SpawnOccupiers{
			Count:              counts[0],
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewHerbivore(pos, OccupierParams.Speed, OccupierParams.HP), nil
			},
		},
		service.SpawnOccupiers{
			Count:              counts[1],
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewPredator(pos, OccupierParams.Speed, OccupierParams.HP, OccupierParams.Damage), nil
			},
		},
		service.SpawnOccupiers{
			Count:              counts[2],
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewTree(pos), nil
			},
		},
		service.SpawnOccupiers{
			Count:              counts[3],
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewGrass(pos), nil
			},
		},
		service.SpawnOccupiers{
			Count:              counts[4],
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewRock(pos), nil
			},
		},
	}
	turnActions := []service.Action{
		service.RespawnOccupiers{
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			MinimalCount:       settings.MinGrassCount,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewGrass(pos), nil
			},
		},
		service.RespawnOccupiers{
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			MinimalCount:       settings.MinHerbivoreCount,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewHerbivore(pos, OccupierParams.Speed, OccupierParams.HP), nil
			},
		},
		service.MoveAction{
			CreatureSettings: settings.CreaturesSettings,
			PathService:      &service.BreadthSearchPathService{},
		},
	}

	sim, err := game.NewSimulation(*settings, service.NewConsoleRenderer(), initActions, turnActions)
	if err != nil {
		fmt.Print(err)
		return
	}
	slog.Debug("Simulation starting.", "simulation", sim)
	sim.Start()
	slog.Debug("Simulation started.")
	for {
		sim.NextTurn()
	}
}

func DistributeCreatures(total, zones, min, max int, varianceFactor float64) ([]int, error) {
	if total < zones*min || total > zones*max {
		return nil, errors.New("невозможно распределить: total выходит за границы min/max")
	}
	if min > max {
		return nil, errors.New("min не может быть больше max")
	}

	result := make([]int, zones)

	base := total / zones
	remainder := total % zones

	for i := 0; i < zones; i++ {
		result[i] = base
		if i < remainder {
			result[i]++
		}
	}

	swaps := int(float64(zones) * 5 * varianceFactor)

	for i := 0; i < swaps; i++ {
		idx1 := rand.IntN(zones)
		idx2 := rand.IntN(zones)
		if idx1 == idx2 {
			continue
		}

		if result[idx1] > min && result[idx2] < max {
			result[idx1]--
			result[idx2]++
		}
	}

	return result, nil
}
