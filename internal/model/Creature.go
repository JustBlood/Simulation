package model

type Creature interface {
	Occupier

	GetSpeed() int
	SetPosition(pos Position)
	GetHp() int
	TakeDamage(int) int
	Heal(int) int
	GetTargetSearchType() OccupierType
}

type Herbivore struct {
	pos   Position
	speed int
	hp    int
}

type Predator struct {
	pos    Position
	speed  int
	hp     int
	damage int
}

func NewHerbivore(pos Position, speed int, hp int) *Herbivore {
	return &Herbivore{pos, speed, hp}
}

func (h *Herbivore) GetType() OccupierType {
	return HERBIVORE
}

func (h *Herbivore) GetPos() Position {
	return h.pos
}

func (h *Herbivore) GetSpeed() int {
	return h.speed
}

func (h *Herbivore) GetHp() int {
	return h.hp
}

func (h *Herbivore) SetPosition(newPos Position) {
	h.pos = newPos
}

func (h *Herbivore) TakeDamage(damage int) int {
	h.hp -= damage
	return h.hp
}

func (h *Herbivore) Heal(healed int) int {
	h.hp += healed
	return h.hp
}

func (h *Herbivore) GetTargetSearchType() OccupierType {
	return GRASS
}

func NewPredator(pos Position, speed int, hp int, damage int) *Predator {
	return &Predator{pos, speed, hp, damage}
}

func (p *Predator) GetType() OccupierType {
	return PREDATOR
}

func (p *Predator) GetPos() Position {
	return p.pos
}

func (p *Predator) GetSpeed() int {
	return p.speed
}

func (p *Predator) GetHp() int {
	return p.hp
}

func (p *Predator) SetPosition(newPos Position) {
	p.pos = newPos
}

func (p *Predator) GetDamage() int {
	return p.damage
}

func (p *Predator) TakeDamage(damage int) int {
	p.hp -= damage
	return p.hp
}

func (p *Predator) Heal(healed int) int {
	p.hp += healed
	return p.hp
}

func (p *Predator) GetTargetSearchType() OccupierType {
	return HERBIVORE
}
