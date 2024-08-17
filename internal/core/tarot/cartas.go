package tarot

import (
	"math/rand"

	"github.com/matienzar/tarot-api/internal/core/domain"
)

// Lista completa de cartas del tarot
var cartasTarot = []domain.CartaTarot{
	// Arcanos Mayores
	{Nombre: "El Loco", Significado: "Nuevos comienzos, aventura, libertad"},
	{Nombre: "El Mago", Significado: "Habilidad, poder, manifestación"},
	{Nombre: "La Sacerdotisa", Significado: "Intuición, misterio, sabiduría interna"},
	{Nombre: "La Emperatriz", Significado: "Abundancia, creatividad, fertilidad"},
	{Nombre: "El Emperador", Significado: "Autoridad, estructura, control"},
	{Nombre: "El Hierofante", Significado: "Tradición, espiritualidad, guía"},
	{Nombre: "Los Enamorados", Significado: "Amor, elección, armonía"},
	{Nombre: "El Carro", Significado: "Victoria, determinación, avance"},
	{Nombre: "La Fuerza", Significado: "Coraje, fortaleza, paciencia"},
	{Nombre: "El Ermitaño", Significado: "Búsqueda interior, soledad, reflexión"},
	{Nombre: "La Rueda de la Fortuna", Significado: "Ciclos, cambio, destino"},
	{Nombre: "La Justicia", Significado: "Equilibrio, justicia, verdad"},
	{Nombre: "El Colgado", Significado: "Sacrificio, cambio de perspectiva, paciencia"},
	{Nombre: "La Muerte", Significado: "Transformación, fin de un ciclo, renacimiento"},
	{Nombre: "La Templanza", Significado: "Equilibrio, moderación, armonía"},
	{Nombre: "El Diablo", Significado: "Tentación, ataduras, materialismo"},
	{Nombre: "La Torre", Significado: "Cambio repentino, ruptura, revelación"},
	{Nombre: "La Estrella", Significado: "Esperanza, inspiración, serenidad"},
	{Nombre: "La Luna", Significado: "Confusión, intuición, ilusiones"},
	{Nombre: "El Sol", Significado: "Éxito, alegría, vitalidad"},
	{Nombre: "El Juicio", Significado: "Renacimiento, evaluación, perdón"},
	{Nombre: "El Mundo", Significado: "Logro, integración, realización"},

	// Arcanos Menores
	// Copas
	{Nombre: "As de Copas", Significado: "Amor, nuevas emociones, intuición"},
	{Nombre: "2 de Copas", Significado: "Unión, asociación, armonía"},
	{Nombre: "3 de Copas", Significado: "Celebración, amistad, alegría"},
	{Nombre: "4 de Copas", Significado: "Apatía, introspección, reevaluación"},
	{Nombre: "5 de Copas", Significado: "Pérdida, duelo, decepción"},
	{Nombre: "6 de Copas", Significado: "Nostalgia, recuerdos, infancia"},
	{Nombre: "7 de Copas", Significado: "Ilusiones, opciones, fantasía"},
	{Nombre: "8 de Copas", Significado: "Abandono, búsqueda espiritual, dejar atrás"},
	{Nombre: "9 de Copas", Significado: "Satisfacción, gratitud, deseos cumplidos"},
	{Nombre: "10 de Copas", Significado: "Felicidad, familia, armonía"},
	{Nombre: "Sota de Copas", Significado: "Mensajes emocionales, creatividad"},
	{Nombre: "Caballo de Copas", Significado: "Romance, idealismo, búsqueda de la verdad"},
	{Nombre: "Reina de Copas", Significado: "Empatía, comprensión, emocionalidad"},
	{Nombre: "Rey de Copas", Significado: "Control emocional, sabiduría, diplomacia"},

	// Oros
	{Nombre: "As de Oros", Significado: "Nuevas oportunidades, prosperidad, seguridad"},
	{Nombre: "2 de Oros", Significado: "Equilibrio, adaptación, gestión"},
	{Nombre: "3 de Oros", Significado: "Trabajo en equipo, habilidad, planificación"},
	{Nombre: "4 de Oros", Significado: "Seguridad, control, posesividad"},
	{Nombre: "5 de Oros", Significado: "Pérdida financiera, inseguridad, desafío"},
	{Nombre: "6 de Oros", Significado: "Generosidad, equilibrio financiero, donación"},
	{Nombre: "7 de Oros", Significado: "Paciencia, reflexión, inversión"},
	{Nombre: "8 de Oros", Significado: "Trabajo duro, maestría, dedicación"},
	{Nombre: "9 de Oros", Significado: "Éxito, autosuficiencia, logros"},
	{Nombre: "10 de Oros", Significado: "Riqueza, legado, estabilidad"},
	{Nombre: "Sota de Oros", Significado: "Nuevas oportunidades financieras, estudio"},
	{Nombre: "Caballo de Oros", Significado: "Responsabilidad, trabajo diligente, método"},
	{Nombre: "Reina de Oros", Significado: "Abundancia, confort, cuidado"},
	{Nombre: "Rey de Oros", Significado: "Éxito material, estabilidad, autoridad"},
	// Espadas
	{Nombre: "As de Espadas", Significado: "Claridad, verdad, revelación"},
	{Nombre: "2 de Espadas", Significado: "Conflicto, indecisión, equilibrio"},
	{Nombre: "3 de Espadas", Significado: "Dolor, tristeza, separación"},
	{Nombre: "4 de Espadas", Significado: "Reposo, recuperación, meditación"},
	{Nombre: "5 de Espadas", Significado: "Conflicto, derrota, estrategia"},
	{Nombre: "6 de Espadas", Significado: "Transición, cambio, viaje"},
	{Nombre: "7 de Espadas", Significado: "Estrategia, evasión, deshonestidad"},
	{Nombre: "8 de Espadas", Significado: "Restricción, miedo, confusión"},
	{Nombre: "9 de Espadas", Significado: "Ansiedad, miedo, pesadillas"},
	{Nombre: "10 de Espadas", Significado: "Ruina, traición, fin de una fase"},
	{Nombre: "Sota de Espadas", Significado: "Curiosidad, vigilancia, ideas nuevas"},
	{Nombre: "Caballo de Espadas", Significado: "Acción rápida, ambición, conflicto"},
	{Nombre: "Reina de Espadas", Significado: "Independencia, claridad, justicia"},
	{Nombre: "Rey de Espadas", Significado: "Autoridad, intelecto, verdad"},

	// Bastos
	{Nombre: "As de Bastos", Significado: "Inspiración, nuevos comienzos, energía"},
	{Nombre: "2 de Bastos", Significado: "Planificación, visión, decisiones"},
	{Nombre: "3 de Bastos", Significado: "Expansión, oportunidades, progreso"},
	{Nombre: "4 de Bastos", Significado: "Celebración, estabilidad, logros"},
	{Nombre: "5 de Bastos", Significado: "Competencia, conflicto, desafío"},
	{Nombre: "6 de Bastos", Significado: "Victoria, reconocimiento, éxito"},
	{Nombre: "7 de Bastos", Significado: "Defensa, determinación, coraje"},
	{Nombre: "8 de Bastos", Significado: "Velocidad, movimiento, comunicación"},
	{Nombre: "9 de Bastos", Significado: "Resiliencia, perseverancia, pruebas"},
	{Nombre: "10 de Bastos", Significado: "Carga, responsabilidad, agotamiento"},
	{Nombre: "Sota de Bastos", Significado: "Entusiasmo, exploración, mensajes"},
	{Nombre: "Caballo de Bastos", Significado: "Acción, aventura, impulsividad"},
	{Nombre: "Reina de Bastos", Significado: "Confianza, liderazgo, creatividad"},
	{Nombre: "Rey de Bastos", Significado: "Autoridad, visión, impulso"},
}

// TirarCartas baraja y selecciona 3 cartas del tarot
func TirarCartas() [3]domain.CartaTarot {
	// Barajar las cartas
	rand.Shuffle(len(cartasTarot), func(i, j int) {
		cartasTarot[i], cartasTarot[j] = cartasTarot[j], cartasTarot[i]
	})

	// Seleccionar las primeras 3 cartas del mazo barajado
	var resultado [3]domain.CartaTarot
	copy(resultado[:], cartasTarot[:3])

	return resultado
}
