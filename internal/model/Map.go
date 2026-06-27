package model

import (
	"errors"
	"simulation/internal/config"
)

type Position struct {
	Height int
	Width  int
}

func NewPosition(height int, width int) Position {
	return Position{height, width}
}

type Map struct {
	PosToOcc map[Position]Occupier
	height   int
	width    int
}

func NewMap(settings config.SimulationSettings) (*Map, error) {
	return &Map{make(map[Position]Occupier), settings.MapHeight, settings.MapWidth}, nil
}

func (m *Map) IsEmpty(pos Position) bool {
	_, occupied := m.PosToOcc[pos]
	return !occupied
}

func (m *Map) Set(pos Position, occ Occupier) error {
	if !m.IsInMapBorders(pos) {
		return errors.New("Выход за границу")
	}
	m.PosToOcc[pos] = occ
	return nil
}

func (m *Map) GetAllCreationsFromMap() []Creature {
	creatures := []Creature{}
	for i := range m.PosToOcc {
		if creature, ok := m.PosToOcc[i].(Creature); ok {
			creatures = append(creatures, creature)
		}
	}
	return creatures
}

func (m *Map) CountOccupiersOfType(occupierType OccupierType) int {
	var count int
	for _, v := range m.PosToOcc {
		if v.GetType() == occupierType {
			count++
		}
	}
	return count
}

func (m *Map) IsInMapBorders(pos Position) bool {
	return pos.Height >= 0 && pos.Height < m.height && pos.Width >= 0 && pos.Width < m.width
}

func (m *Map) GetHeight() int {
	return m.height
}

func (m *Map) GetWidth() int {
	return m.width
}
