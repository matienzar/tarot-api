package repository

import (
	"github.com/matienzar/tarot-api/internal/core/domain"
)

type InMemoryPreguntaRepository struct {
	data map[int]*domain.Pregunta
}

func NewInMemoryPreguntaRepository() *InMemoryPreguntaRepository {
	return &InMemoryPreguntaRepository{data: make(map[int]*domain.Pregunta)}
}

func (repo *InMemoryPreguntaRepository) Save(pregunta *domain.Pregunta) error {
	repo.data[pregunta.ID] = pregunta
	return nil
}

func (repo *InMemoryPreguntaRepository) FindByID(id int) (*domain.Pregunta, error) {
	if pregunta, exists := repo.data[id]; exists {
		return pregunta, nil
	}
	return nil, nil // o podemos devolver un error
}
