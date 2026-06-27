package config

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

var globalSettings *GlobalSettings

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
	if globalSettings == nil {
		data, err := os.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("error in reading file: %w", err)
		}

		var simSettings SimulationSettings
		var creatureSettings CreaturesSettings
		var occSettings OccupiersSettings
		err = json.Unmarshal(data, &simSettings)
		if err != nil {
			return nil, fmt.Errorf("error in JSON parsing: %w", err)
		}
		err = json.Unmarshal(data, &creatureSettings)
		if err != nil {
			return nil, fmt.Errorf("error in JSON parsing: %w", err)
		}
		err = json.Unmarshal(data, &occSettings)
		if err != nil {
			return nil, fmt.Errorf("error in JSON parsing: %w", err)
		}
		globalSettings = &GlobalSettings{simSettings, creatureSettings, occSettings}
	}

	return globalSettings, nil
}

func InitLogger() *os.File {
	logFile, err := os.OpenFile("./log/simulation.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	// Структурированный JSON-логгер (удобно grep'ать)
	logger := slog.New(slog.NewTextHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	return logFile
}
