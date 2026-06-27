package model

type Rock struct {
	baseOccupier
}

func (Rock) GetType() OccupierType {
	return ROCK
}

func NewRock(pos Position) *Rock {
	return &Rock{baseOccupier{pos}}
}
