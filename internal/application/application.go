package application

import (
	"log/slog"
	"simulation/internal/config"
	"simulation/internal/service"
)

type SimulationApp interface {
	Start()
}

type consoleEndlessSimulationApp struct {
	settings *config.GlobalSettings
	renderer service.Renderer
}

func NewConsoleEndlessSimulationApp(settings *config.GlobalSettings) *consoleEndlessSimulationApp {
	return &consoleEndlessSimulationApp{settings: settings, renderer: service.NewConsoleRenderer()}
}

func (app *consoleEndlessSimulationApp) Start() {
	sim, err := service.NewSimulation(app.settings, app.renderer)
	if err != nil {
		slog.Error("Can't start simulation", "error", err)
		return
	}
	slog.Debug("Simulation starting.", "simulation", sim)
	sim.Start()
	slog.Debug("Simulation started.")
	for {
		sim.NextTurn()
	}
}
