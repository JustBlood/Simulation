package model

import (
	"errors"
	"simulation/internal/config"
)

type Map struct {
	PosToOcc map[Position]Occupier
	height   int
	width    int
}

type Position struct {
	Height int
	Width  int
}

type Pair[T any] struct {
	first  T
	second T
}

func NewPosition(height int, width int) Position {
	return Position{height, width}
}

// fixme: перенести всё что ниже в actions, main.go

func NewMap(settings config.SimulationSettings) (*Map, error) {
	return &Map{make(map[Position]Occupier), settings.MapHeight, settings.MapWidth}, nil
}

func (m *Map) IsEmpty(pos Position) bool {
	_, occupied := m.PosToOcc[pos]
	return !occupied
}

func (m *Map) Set(pos Position, occ Occupier) error {
	if !m.IsInMap(pos) {
		return errors.New("Выход за границу")
	}
	m.PosToOcc[pos] = occ
	return nil
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

func (m *Map) IsInMap(pos Position) bool {
	return pos.Height >= 0 && pos.Height < m.height && pos.Width >= 0 && pos.Width < m.width
}

func (m *Map) GetHeight() int {
	return m.height
}

func (m *Map) GetWidth() int {
	return m.width
}
