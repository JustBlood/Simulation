package model

type Creature interface {
	Occupier

	GetSpeed() int
	SetPosition(pos Position)
	GetHp() int
	TakeDamage(int)
	Heal(int) int
	GetTargetSearchType() OccupierType
	IsDead() bool
}

type BaseCreature struct {
	BaseOccupier
	speed int
	hp    int
}

type Herbivore struct {
	BaseCreature
}

type Predator struct {
	BaseCreature
	damage int
}

func (bc *BaseCreature) GetSpeed() int {
	return bc.speed
}

func (bc *BaseCreature) GetHp() int {
	return bc.hp
}

func (bc *BaseCreature) SetPosition(newPos Position) {
	bc.pos = newPos
}

func (bc *BaseCreature) TakeDamage(damage int) {
	bc.hp -= damage
}

func (bc *BaseCreature) Heal(healed int) int {
	bc.hp += healed
	return bc.hp
}

func (bc *BaseCreature) IsDead() bool {
	return bc.hp <= 0
}

func NewHerbivore(pos Position, speed int, hp int) *Herbivore {
	return &Herbivore{
		BaseCreature: BaseCreature{
			BaseOccupier: BaseOccupier{pos: pos},
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

func NewPredator(pos Position, speed int, hp int, damage int) *Predator {
	return &Predator{
		BaseCreature: BaseCreature{
			BaseOccupier: BaseOccupier{pos: pos},
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
