package game

import (
	"log/slog"
	"simulation/internal/config"
	"simulation/internal/model"
	"simulation/internal/service"
	"time"
)

type Simulation struct {
	gameMap        *model.Map
	renderer       service.Renderer
	renderInterval int
	step           int
	initActions    []service.Action
	turnActions    []service.Action
}

func NewSimulation(settings config.GlobalSettings, renderer service.Renderer, initActions []service.Action, turnActions []service.Action) (*Simulation, error) {
	gameMap, err := model.NewMap(settings.SimulationSettings)
	if err != nil {
		return nil, err
	}
	return &Simulation{gameMap, renderer, settings.RenderIntervalMs, 0, initActions, turnActions}, nil
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
