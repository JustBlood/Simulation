package service

import (
	"math/rand/v2"
	"simulation/internal/config"
	"simulation/internal/model"
)

type FactoryFunc func(OccupierParams model.OccupierParams, pos model.Position) (model.Occupier, error)

type Action interface {
	RunAction(gameMap *model.Map) error
}

func GenerateRandomParams(settings config.CreaturesSettings) model.OccupierParams {
	return model.OccupierParams{
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
