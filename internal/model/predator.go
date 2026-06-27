package model

type Predator struct {
	baseCreature
	damage int
}

func NewPredator(pos Position, speed int, hp int, damage int) *Predator {
	return &Predator{
		baseCreature: baseCreature{
			baseOccupier: baseOccupier{pos: pos},
			speed:        speed,
			hp:           hp,
		},
		damage: damage,
	}
}

func (p *Predator) GetType() OccupierType {
	return PREDATOR
}

func (p *Predator) GetTargetSearchType() OccupierType {
	return HERBIVORE
}

func (p *Predator) GetDamage() int {
	return p.damage
}
