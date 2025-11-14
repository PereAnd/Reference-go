// Package tournament proporciona manejadores HTTP para endpoints relacionados con torneos.
// Este paquete implementa los adaptadores conductores (entrada) en la arquitectura hexagonal,
// traduciendo peticiones HTTP en llamadas a servicios y respuestas HTTP. Usa el framework
// Gin para el enrutamiento HTTP y manejo de peticiones.
package tournament

import "github.com/jairogloz/go-l/pkg/ports"

// Handler proporciona manejadores HTTP para operaciones de torneos.
// Actúa como un adaptador entre la capa HTTP (Gin) y la capa de servicio de aplicación,
// traduciendo peticiones/respuestas HTTP a/desde operaciones del dominio.
type Handler struct {
	// TournamentService es el servicio de aplicación usado para realizar operaciones de negocio de torneos.
	// Implementa la interfaz TournamentService, manteniendo el handler desacoplado de
	// implementaciones específicas de servicios.
	TournamentService ports.TournamentService
}
