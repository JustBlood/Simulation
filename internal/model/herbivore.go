package model

type Herbivore struct {
	baseCreature
}

func NewHerbivore(pos Position, speed int, hp int) *Herbivore {
	return &Herbivore{
		baseCreature: baseCreature{
			baseOccupier: baseOccupier{pos: pos},
			speed:        speed,
			hp:           hp,
		},
	}
}

func (h *Herbivore) GetType() OccupierType {
	return HERBIVORE
}

func (h *Herbivore) GetTargetSearchType() OccupierType {
	return GRASS
}
