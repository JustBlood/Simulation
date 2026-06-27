package service

import (
	"simulation/internal/model"
)

type Action interface {
	RunAction(gameMap *model.Map) error
}
