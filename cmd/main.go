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

	counts, _ := DistributeCreatures(needToCreateOccupiers, len(model.AllEntityTypes), 1, 20, 0.5)

	// herbivoreCount := settings.MinHerbivoreCount + rand.IntN(needToCreateOccupiers-settings.MinHerbivoreCount)
	// needToCreateOccupiers -= herbivoreCount
	// grassCount := settings.MinGrassCount + rand.IntN(needToCreateOccupiers-settings.MinGrassCount)
	// needToCreateOccupiers -= grassCount
	// predatorCount := rand.IntN(needToCreateOccupiers)
	// needToCreateOccupiers -= predatorCount
	// treeCount := needToCreateOccupiers / 2
	// needToCreateOccupiers -= treeCount
	// rockCount := needToCreateOccupiers

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
			MinimalCount:       1,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewHerbivore(pos, OccupierParams.Speed, OccupierParams.HP), nil
			},
		},
		service.MoveAction{
			CreatureSettings: settings.CreaturesSettings,
		},
	}

	sim, err := game.NewSimulation(*settings, service.RenderInConsole, initActions, turnActions)
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
	// 1. Проверка на физическую возможность
	if total < zones*min || total > zones*max {
		return nil, errors.New("невозможно распределить: total выходит за границы min/max")
	}
	if min > max {
		return nil, errors.New("min не может быть больше max")
	}

	result := make([]int, zones)

	// 2. Шаг "Идеальная равномерность"
	// Распределяем базу
	base := total / zones
	remainder := total % zones // Остаток, который нужно распределить по 1 штуке

	for i := 0; i < zones; i++ {
		result[i] = base
		if i < remainder {
			result[i]++ // Первым зонам достается по +1, чтобы сумма сошлась ровно в total
		}
	}

	// 3. Шаг "Вносим случайность" (Метод переливаний)
	// Сколько раз мы будем "перекидывать" существ между зонами.
	// varianceFactor = 0.0 -> 0 перестановок (идеально ровно)
	// varianceFactor = 1.0 -> zones * 5 перестановок (умеренный хаос)
	// varianceFactor = 3.0 -> zones * 15 перестановок (сильный хаос, но в рамках min/max)
	swaps := int(float64(zones) * 5 * varianceFactor)

	for i := 0; i < swaps; i++ {
		// Выбираем две случайные разные зоны
		idx1 := rand.IntN(zones)
		idx2 := rand.IntN(zones)
		if idx1 == idx2 {
			continue
		}

		// Пытаемся забрать у idx1 и отдать idx2
		// Проверяем, что idx1 может отдать (не уйдет ниже min)
		// и idx2 может принять (не уйдет выше max)
		if result[idx1] > min && result[idx2] < max {
			result[idx1]--
			result[idx2]++
		}
	}

	return result, nil
}
