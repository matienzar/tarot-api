package application

import "github.com/matienzar/tarot-api/internal/core/domain"

// PreguntaRepository define las operaciones que se pueden realizar sobre las preguntas.
type PreguntaRepository interface {
	Save(pregunta *domain.Pregunta) (int, error)
	FindAll() ([]*domain.Pregunta, error)
	FindByID(id int) (*domain.Pregunta, error) // Asegúrate de que este método esté definido si es necesario
}

// TarotAI define la interfaz para interactuar con un sistema de IA que genera respuestas basadas en las cartas del tarot.
type TarotAI interface {
	TirarCartas() [3]domain.CartaTarot
	GenerarRespuesta(pregunta string, cartas [3]domain.CartaTarot) (domain.Respuesta, error)
}
