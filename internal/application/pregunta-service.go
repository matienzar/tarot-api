package application

import (
	"fmt"
	"time"

	"github.com/matienzar/tarot-api/internal/core/domain"
)

type PreguntaService struct {
	repo  PreguntaRepository
	tarot TarotAI
}

func NewPreguntaService(repo PreguntaRepository, tarot TarotAI) *PreguntaService {
	return &PreguntaService{
		repo:  repo,
		tarot: tarot,
	}
}

func (s *PreguntaService) RealizarPregunta(texto string) (*domain.Pregunta, error) {
	if s.tarot == nil {
		return nil, fmt.Errorf("tarot service es nil")
	}

	cartas := s.tarot.TirarCartas() // Aquí podría estar fallando si s.tarot es nil
	respuesta, err := s.tarot.GenerarRespuesta(texto, cartas)
	if err != nil {
		return nil, err
	}

	pregunta := &domain.Pregunta{
		Texto:     texto,
		FechaHora: time.Now(),
		Cartas:    cartas,
		Respuesta: respuesta,
	}

	if s.repo == nil {
		return nil, fmt.Errorf("repository es nil")
	}

	id, err := s.repo.Save(pregunta)
	if err != nil {
		return nil, err
	}

	// Recuperar el ID del registro guardado
	preguntaGuardada, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return preguntaGuardada, nil
}
