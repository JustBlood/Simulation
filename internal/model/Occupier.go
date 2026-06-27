package model

import (
	"fmt"
)

type OccupierType string

const (
	ROCK      OccupierType = "ROCK"
	TREE      OccupierType = "TREE"
	GRASS     OccupierType = "GRASS"
	HERBIVORE OccupierType = "HERBIVORE"
	PREDATOR  OccupierType = "PREDATOR"
)

// 3. Единый источник истины — слайс всех возможных значений
var AllEntityTypes = []OccupierType{
	ROCK,
	TREE,
	GRASS,
	HERBIVORE,
	PREDATOR,
}

var CreatureOccupiersTypes = []OccupierType{
	HERBIVORE,
	PREDATOR,
}

type OccupierParams struct {
	Speed  int
	HP     int
	Damage int
}

type Occupier interface {
	GetType() OccupierType
	GetPos() Position
}

func NewOccupier(entityType OccupierType, params OccupierParams, pos Position) (Occupier, error) {
	switch entityType {
	case ROCK:
		return NewRock(pos), nil
	case TREE:
		return NewTree(pos), nil
	case GRASS:
		return NewGrass(pos), nil
	case HERBIVORE:
		return NewHerbivore(pos, params.Speed, params.HP), nil
	case PREDATOR:
		return NewPredator(pos, params.Speed, params.HP, params.Damage), nil
	default:
		return nil, fmt.Errorf("unknown entityType: %s", entityType)
	}
}

type BaseOccupier struct {
	pos Position
}

func (bo *BaseOccupier) GetPos() Position {
	return bo.pos
}

type Grass struct {
	BaseOccupier
}

func (Grass) GetType() OccupierType {
	return GRASS
}

func NewGrass(pos Position) *Grass {
	return &Grass{BaseOccupier{pos}}
}

type Tree struct {
	BaseOccupier
}

func (Tree) GetType() OccupierType {
	return TREE
}

func NewTree(pos Position) *Tree {
	return &Tree{BaseOccupier{pos}}
}

type Rock struct {
	BaseOccupier
}

func (Rock) GetType() OccupierType {
	return ROCK
}

func NewRock(pos Position) *Rock {
	return &Rock{BaseOccupier{pos}}
}
