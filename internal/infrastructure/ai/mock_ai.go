// /internal/infrastructure/ai/mock_ai.go
package ai

import (
	"fmt"

	"github.com/matienzar/tarot-api/internal/application"
	"github.com/matienzar/tarot-api/internal/core/domain"
	"github.com/matienzar/tarot-api/internal/core/tarot"
)

type MockAI struct{}

func NewMockAI() application.TarotAI {
	return &MockAI{}
}

func (ai *MockAI) GenerarRespuesta(pregunta string, cartas [3]domain.CartaTarot) (domain.Respuesta, error) {
	respuestaTexto := fmt.Sprintf("Basado en tu pregunta '%s', las cartas indican %s, %s y %s.",
		pregunta, cartas[0].Significado, cartas[1].Significado, cartas[2].Significado)
	return domain.Respuesta{Texto: respuestaTexto}, nil
}

func (ai *MockAI) TirarCartas() [3]domain.CartaTarot {
	return tarot.TirarCartas()
}
