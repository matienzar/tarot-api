// /internal/core/domain/question.go
package domain

import (
	"time"
)

type Pregunta struct {
	ID        int
	Texto     string
	FechaHora time.Time
	Cartas    [3]CartaTarot
	Respuesta Respuesta
}

type CartaTarot struct {
	Nombre      string
	Significado string
}

type Respuesta struct {
	Texto string
}
