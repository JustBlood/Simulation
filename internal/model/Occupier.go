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

type Grass struct {
	pos Position
}

func (Grass) GetType() OccupierType {
	return GRASS
}

func (g Grass) GetPos() Position {
	return g.pos
}

func NewGrass(pos Position) *Grass {
	return &Grass{pos}
}

type Tree struct {
	pos Position
}

func (Tree) GetType() OccupierType {
	return TREE
}

func (t Tree) GetPos() Position {
	return t.pos
}

func NewTree(pos Position) *Tree {
	return &Tree{pos}
}

type Rock struct {
	pos Position
}

func (Rock) GetType() OccupierType {
	return ROCK
}

func (r Rock) GetPos() Position {
	return r.pos
}

func NewRock(pos Position) *Rock {
	return &Rock{pos}
}
