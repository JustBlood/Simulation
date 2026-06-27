package game

import (
	"log/slog"
	"simulation/internal/config"
	"simulation/internal/model"
	"simulation/internal/service"
)

type Simulation struct {
	gameMap     *model.Map
	renderer    service.Render
	step        int
	initActions []service.Action
	turnActions []service.Action
}

func NewSimulation(settings config.GlobalSettings, renderer service.Render, initActions []service.Action, turnActions []service.Action) (*Simulation, error) {
	gameMap, err := model.NewMap(settings.SimulationSettings)
	if err != nil {
		return nil, err
	}
	return &Simulation{gameMap, renderer, 0, initActions, turnActions}, nil
}

func (s *Simulation) NextTurn() {
	slog.Debug("Simulation.NextTurn called")
	for _, action := range s.turnActions {
		action.RunAction(s.gameMap)
	}
	s.renderer(s.gameMap)
}

func (s *Simulation) Start() {
	for _, action := range s.initActions {
		action.RunAction(s.gameMap)
	}
	s.renderer(s.gameMap)
}

func (s *Simulation) Pause() {

}
