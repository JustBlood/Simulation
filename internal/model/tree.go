package model

type Tree struct {
	baseOccupier
}

func (Tree) GetType() OccupierType {
	return TREE
}

func NewTree(pos Position) *Tree {
	return &Tree{baseOccupier{pos}}
}
