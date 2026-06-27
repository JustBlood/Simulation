package model

import (
	"math/rand/v2"
	"simulation/internal/config"
)

type OccupierParams struct {
	Speed  int
	HP     int
	Damage int
}

type FactoryFunc func(OccupierParams OccupierParams, pos Position) (Occupier, error)

func GenerateRandomParams(settings config.CreaturesSettings) OccupierParams {
	return OccupierParams{
		Speed:  randomInt(settings.MinSpeed, settings.MaxSpeed),
		HP:     randomInt(settings.MinHp, settings.MaxHp),
		Damage: randomInt(settings.MinDamage, settings.MaxDamage),
	}
}

func randomInt(min, max int) int {
	if min >= max {
		return min
	}
	return min + rand.IntN(max-min)
}
