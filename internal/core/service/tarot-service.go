package service

import (
	"github.com/matienzar/tarot-api/internal/core/domain"
	"github.com/matienzar/tarot-api/internal/core/tarot"
)

type TarotService struct{}

func (s *TarotService) TirarCartas() [3]domain.CartaTarot {
	return tarot.TirarCartas()
}

func (s *TarotService) GenerarRespuesta(texto string, cartas [3]domain.CartaTarot) (domain.Respuesta, error) {
	// Implementación de generación de respuesta
	return domain.Respuesta{Texto: "Respuesta generada"}, nil
}
