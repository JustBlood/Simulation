package model

type OccupierType string

const (
	ROCK      OccupierType = "ROCK"
	TREE      OccupierType = "TREE"
	GRASS     OccupierType = "GRASS"
	HERBIVORE OccupierType = "HERBIVORE"
	PREDATOR  OccupierType = "PREDATOR"
)

var AllEntityTypes = []OccupierType{
	ROCK,
	TREE,
	GRASS,
	HERBIVORE,
	PREDATOR,
}

type Occupier interface {
	GetType() OccupierType
	GetPos() Position
	SetPos(newPos Position)
}

type baseOccupier struct {
	pos Position
}

func (bo *baseOccupier) GetPos() Position {
	return bo.pos
}

func (bo *baseOccupier) SetPos(newPos Position) {
	bo.pos = newPos
}
