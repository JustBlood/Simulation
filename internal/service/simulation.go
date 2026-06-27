package service

import (
	"log/slog"
	"math/rand/v2"
	"simulation/internal/config"
	"simulation/internal/model"
	"time"
)

type Simulation struct {
	gameMap        *model.Map
	renderer       Renderer
	renderInterval int
	initActions    []Action
	turnActions    []Action
}

func NewSimulation(settings *config.GlobalSettings, renderer Renderer) (*Simulation, error) {
	needToCreateOccupiers := settings.MinOccupiers + rand.IntN(settings.MaxOccupiers-settings.MinOccupiers)

	counts, err := distributeCreatures(needToCreateOccupiers, len(model.AllEntityTypes))

	if err != nil {
		slog.Error("Error in initialization", "error", err)
		return nil, err
	}

	initActions := getInitActions(counts, settings)
	turnActions := getTurnActions(settings)

	gameMap, err := model.NewMap(settings.SimulationSettings)
	if err != nil {
		return nil, err
	}
	return &Simulation{gameMap, renderer, settings.RenderIntervalMs, initActions, turnActions}, nil
}

func (s *Simulation) NextTurn() {
	slog.Debug("Simulation.NextTurn called")
	for _, action := range s.turnActions {
		action.RunAction(s.gameMap)
	}
	s.renderer.Render(s.gameMap)
	time.Sleep(time.Duration(s.renderInterval) * time.Millisecond)
}

func (s *Simulation) Start() {
	for _, action := range s.initActions {
		action.RunAction(s.gameMap)
	}
	s.renderer.Render(s.gameMap)
}

func (s *Simulation) Stop() {
	s.renderer.Close()
}

func getInitActions(counts []int, settings *config.GlobalSettings) []Action {
	return []Action{
		SpawnOccupiers{
			Count:              counts[0],
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewHerbivore(pos, OccupierParams.Speed, OccupierParams.HP), nil
			},
		},
		SpawnOccupiers{
			Count:              counts[1],
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewPredator(pos, OccupierParams.Speed, OccupierParams.HP, OccupierParams.Damage), nil
			},
		},
		SpawnOccupiers{
			Count:              counts[2],
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewTree(pos), nil
			},
		},
		SpawnOccupiers{
			Count:              counts[3],
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewGrass(pos), nil
			},
		},
		SpawnOccupiers{
			Count:              counts[4],
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewRock(pos), nil
			},
		},
	}
}

func getTurnActions(settings *config.GlobalSettings) []Action {
	return []Action{
		RespawnOccupiers{
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			MinimalCount:       settings.MinGrassCount,
			OccupierType:       model.GRASS,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewGrass(pos), nil
			},
		},
		RespawnOccupiers{
			SimulationSettings: settings.SimulationSettings,
			CreaturesSettings:  settings.CreaturesSettings,
			MinimalCount:       settings.MinHerbivoreCount,
			OccupierType:       model.HERBIVORE,
			FactoryFunc: func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error) {
				return model.NewHerbivore(pos, OccupierParams.Speed, OccupierParams.HP), nil
			},
		},
		MoveAction{
			CreatureSettings: settings.CreaturesSettings,
			PathService:      &BreadthSearchPathService{},
		},
	}
}

func distributeCreatures(total, zones int) ([]int, error) {
	result := make([]int, zones)

	base := total / zones
	remainder := total % zones

	for i := 0; i < zones; i++ {
		result[i] = base
		if i < remainder {
			result[i]++
		}
	}

	return result, nil
}
