package main

import (
	"log"
	"net/http"
	"os"

	"github.com/matienzar/tarot-api/internal/application"
	"github.com/matienzar/tarot-api/internal/infrastructure/ai"
	"github.com/matienzar/tarot-api/internal/infrastructure/repository"
)

func main() {
	// Alternar entre MockAI y OpenAITarotAI basado en una variable de entorno
	var tarotAI application.TarotAI
	if os.Getenv("USE_HUGGINGFACE") == "true" {
		tarotAI = ai.NewHuggingFaceAI()
	} else {
		tarotAI = ai.NewMockAI()
	}

	// Initialize the SQLite repository
	repo, err := repository.NewSQLitePreguntaRepository("file:/tmp/tarot.db")
	if err != nil {
		log.Fatalf("Error inicializando el repository: %v", err)
	}

	preguntaService := application.NewPreguntaService(repo, tarotAI)
	apiHandler := application.NewAPIHandler(preguntaService)

	http.HandleFunc("/pregunta", apiHandler.HandleRequest)

	port := "8080"
	log.Printf("HTTP Server corriendo en el puerto %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Falló el arranque del servidor: %v", err)
	}
}

/*
func main() {

	// Alternar entre MockAI y OpenAITarotAI basado en una variable de entorno
	var tarotAI application.TarotAI
	if os.Getenv("USE_HUGGINGFACE") == "true" {
		tarotAI = ai.NewHuggingFaceAI()
	} else {
		tarotAI = ai.NewMockAI()
	}

	tarotAI = ai.NewHuggingFaceAI()
	// Initialize the SQLite repository
	repo, err := repository.NewSQLitePreguntaRepository("file:/tmp/tarot.db")
	if err != nil {
		log.Fatalf("Error initializing repository: %v", err)
	}

	preguntaService := application.NewPreguntaService(repo, tarotAI)

	// Pedimos al usuario que nos pregunte
	fmt.Println("Por favor, escribe tu pregunta para el tarot y presiona Enter:")
	reader := bufio.NewReader(os.Stdin)
	preguntaARealizar, _ := reader.ReadString('\n')

	// Eliminar el salto de línea al final de la entrada si no está vacío
	if len(preguntaARealizar) > 1 {
		preguntaARealizar = preguntaARealizar[:len(preguntaARealizar)-1]
	} else {
		preguntaARealizar = "¿El Barça ganará la liga de fútbol?"
	}

	// Ejemplo de cómo podrías usar el servicio para realizar una pregunta
	pregunta, err := preguntaService.RealizarPregunta(preguntaARealizar)
	if err != nil {
		panic(err)
	}

	// Imprimir la respuesta obtenida
	println(pregunta.Respuesta.Texto)

}
*/
