package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var globalSettings *GlobalSettings

type GlobalSettings struct {
	SimulationSettings
	CreaturesSettings
	OccupiersSettings
}

type SimulationSettings struct {
	MapWidth  int `json:"map_width"`
	MapHeight int `json:"map_height"`
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
			return nil, fmt.Errorf("ошибка чтения файла: %w", err)
		}

		var simSettings SimulationSettings
		var creatureSettings CreaturesSettings
		var occSettings OccupiersSettings
		err = json.Unmarshal(data, &simSettings)
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
		}
		err = json.Unmarshal(data, &creatureSettings)
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
		}
		err = json.Unmarshal(data, &occSettings)
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
		}
		globalSettings = &GlobalSettings{simSettings, creatureSettings, occSettings}
	}

	return globalSettings, nil
}
