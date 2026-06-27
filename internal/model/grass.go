package model

type Grass struct {
	baseOccupier
}

func (Grass) GetType() OccupierType {
	return GRASS
}

func NewGrass(pos Position) *Grass {
	return &Grass{baseOccupier{pos}}
}
