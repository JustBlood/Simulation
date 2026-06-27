package model

type Creature interface {
	Occupier

	Speed() int
	HP() int
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

func (bc *baseCreature) Speed() int {
	return bc.speed
}

func (bc *baseCreature) HP() int {
	return bc.hp
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
