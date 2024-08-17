package application

import (
	"encoding/json"
	"net/http"

	"github.com/matienzar/tarot-api/internal/core/domain"
)

type APIHandler struct {
	preguntaService *PreguntaService
}

func NewAPIHandler(preguntaService *PreguntaService) *APIHandler {
	return &APIHandler{preguntaService: preguntaService}
}

func (h *APIHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var requestPayload struct {
			Texto string `json:"texto"`
		}

		if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
			http.Error(w, "PAYLOAD incorrecto", http.StatusBadRequest)
			return
		}

		pregunta, err := h.preguntaService.RealizarPregunta(requestPayload.Texto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responsePayload := struct {
			ID        int    `json:"id"`
			Respuesta string `json:"respuesta"`
		}{
			ID:        pregunta.ID,
			Respuesta: pregunta.Respuesta.Texto,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responsePayload)
	case http.MethodGet:
		preguntas, err := h.preguntaService.repo.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responsePayload := struct {
			Preguntas []*domain.Pregunta `json:"preguntas"`
		}{
			Preguntas: preguntas,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responsePayload)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
