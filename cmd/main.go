package main

import (
	"fmt"
	"simulation/internal/application"
	"simulation/internal/config"
)

func main() {
	logFile := config.InitLogger()
	defer logFile.Close()

	settings, err := config.LoadGlobalSettings("./config/config.json")
	if err != nil {
		fmt.Print(err)
		return
	}

	var app application.SimulationApp
	app = application.NewConsoleEndlessSimulationApp(settings)
	app.Start()
}
