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

type baseCreature struct {
	baseOccupier
	speed int
	hp    int
}

func (bc *baseCreature) GetSpeed() int {
	return bc.speed
}

func (bc *baseCreature) GetHp() int {
	return bc.hp
}

func (bc *baseCreature) SetPosition(newPos Position) {
	bc.pos = newPos
}

func (bc *baseCreature) TakeDamage(damage int) {
	bc.hp -= damage
}

func (bc *baseCreature) Heal(healed int) int {
	bc.hp += healed
	return bc.hp
}

func (bc *baseCreature) IsDead() bool {
	return bc.hp <= 0
}
