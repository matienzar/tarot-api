package ai

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hupe1980/go-huggingface"
	"github.com/matienzar/tarot-api/internal/application"
	"github.com/matienzar/tarot-api/internal/core/domain"
	"github.com/matienzar/tarot-api/internal/core/tarot"
)

type HuggingFaceAI struct {
	apiKey string
}

func NewHuggingFaceAI() application.TarotAI {
	apiKey := os.Getenv("HUGGINGFACE_API_KEY")

	if apiKey == "" {
		log.Fatal("HUGGINGFACE_API_KEY environment variable is not set")
	}

	return &HuggingFaceAI{apiKey: apiKey}
}

func (ai *HuggingFaceAI) TirarCartas() [3]domain.CartaTarot {
	return tarot.TirarCartas()
}

func (ai *HuggingFaceAI) GenerarRespuesta(pregunta string, cartas [3]domain.CartaTarot) (domain.Respuesta, error) {

	model := "mistralai/Mistral-7B-Instruct-v0.2"
	numSequences := 30
	maxLength := 1000 // Ajusta este valor según lo que necesites
	// fmt.Println("Modelo: ", model)
	ic := huggingface.NewInferenceClient(ai.apiKey)

	// Formatear el mensaje/prompt
	mensaje := fmt.Sprintf(
		"Como experto tarotista español, te preguntan: '%s'. Barajas las cartas y sacas: '%s', '%s' y '%s'. Las cartas me dicen que ",
		pregunta,
		cartas[0].Nombre,
		cartas[1].Nombre,
		cartas[2].Nombre,
	)

	// Preparar los datos para la solicitud de conversación
	request := huggingface.TextGenerationRequest{
		Inputs: mensaje,
		Options: huggingface.Options{
			WaitForModel: boolPointer(true),
			UseCache:     boolPointer(false),
		},
		Parameters: huggingface.TextGenerationParameters{
			ReturnFullText:     boolPointer(true),
			MaxNewTokens:       &maxLength,
			NumReturnSequences: &numSequences,
		},
		Model: model, // Reemplaza con el modelo adecuado
	}

	// Enviar la solicitud al endpoint de conversación
	res, err := ic.TextGeneration(context.Background(), &request)
	if err != nil {
		log.Fatal(err)
	}

	// Acceder al texto generado
	if len(res) > 0 {
		return domain.Respuesta{Texto: res[0].GeneratedText}, nil
	}

	return domain.Respuesta{Texto: ""}, nil // Maneja el caso en el que no haya texto generado

}

func boolPointer(b bool) *bool {
	return &b
}
