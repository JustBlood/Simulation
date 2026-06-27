package model

import (
	"errors"
	"simulation/internal/config"
)

type Position struct {
	Row    int
	Column int
}

func NewPosition(row int, column int) Position {
	return Position{row, column}
}

type Map struct {
	posToOcc map[Position]Occupier
	height   int
	width    int
}

func NewMap(settings config.SimulationSettings) (*Map, error) {
	return &Map{make(map[Position]Occupier), settings.MapHeight, settings.MapWidth}, nil
}

func (m *Map) GetOccupierByPosition(pos Position) (Occupier, error) {
	occ, exists := m.posToOcc[pos]
	if !exists {
		return nil, errors.New("no occupier for position")
	}
	return occ, nil
}

func (m *Map) ReplaceOccupierPosition(occ Occupier, newPos Position) error {
	if !m.IsInMapBorders(newPos) {
		return errors.New("newPos isn't in map borders")
	}
	delete(m.posToOcc, occ.GetPos())
	occ.SetPos(newPos)
	m.posToOcc[newPos] = occ
	return nil
}

func (m *Map) DeleteOccupierByPosition(pos Position) {
	delete(m.posToOcc, pos)
}

func (m *Map) IsEmpty(pos Position) bool {
	_, occupied := m.posToOcc[pos]
	return !occupied
}

func (m *Map) SetOccupierToPosition(pos Position, occ Occupier) error {
	if !m.IsInMapBorders(pos) {
		return errors.New("Выход за границу")
	}
	m.posToOcc[pos] = occ
	return nil
}

func (m *Map) GetAllCreations() []Creature {
	creatures := []Creature{}
	for i := range m.posToOcc {
		if creature, ok := m.posToOcc[i].(Creature); ok {
			creatures = append(creatures, creature)
		}
	}
	return creatures
}

func (m *Map) CountOccupiersOfType(occupierType OccupierType) int {
	var count int
	for _, v := range m.posToOcc {
		if v.GetType() == occupierType {
			count++
		}
	}
	return count
}

func (m *Map) IsInMapBorders(pos Position) bool {
	return pos.Row >= 0 && pos.Row < m.height && pos.Column >= 0 && pos.Column < m.width
}

func (m *Map) GetHeight() int {
	return m.height
}

func (m *Map) GetWidth() int {
	return m.width
}
