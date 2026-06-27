package config

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

type GlobalSettings struct {
	SimulationSettings
	CreaturesSettings
	OccupiersSettings
}

type SimulationSettings struct {
	MapWidth         int `json:"map_width"`
	MapHeight        int `json:"map_height"`
	RenderIntervalMs int `json:"render_interval_ms"`
}

type CreaturesSettings struct {
	MinOccupiers  int `json:"min_creatures"`
	MaxOccupiers  int `json:"max_creatures"`
	MinSpeed      int `json:"min_speed"`
	MaxSpeed      int `json:"max_speed"`
	MinDamage     int `json:"min_damage"`
	MaxDamage     int `json:"max_damage"`
	MinHp         int `json:"min_hp"`
	MaxHp         int `json:"max_hp"`
	HerbivoreHeal int `json:"herbivore_heal"`
	PredatorHeal  int `json:"predator_heal"`
}

type OccupiersSettings struct {
	MinGrassCount     int `json:"min_grass_count"`
	MinHerbivoreCount int `json:"min_herbivore_count"`
}

func LoadGlobalSettings(filename string) (*GlobalSettings, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error in reading file: %w", err)
	}

	var settings GlobalSettings
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}
	return &settings, nil
}

func InitLogger() (*os.File, error) {
	logFile, err := os.OpenFile("./log/simulation.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	return logFile, nil
}
