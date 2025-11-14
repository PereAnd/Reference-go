// Package player proporciona manejadores HTTP para endpoints relacionados con jugadores.
// Este paquete implementa los adaptadores conductores (entrada) en la arquitectura hexagonal,
// traduciendo peticiones HTTP en llamadas a servicios y respuestas HTTP. Usa el framework
// Gin para el enrutamiento HTTP y manejo de peticiones.
package player

import "github.com/jairogloz/go-l/pkg/ports"

// Handler proporciona manejadores HTTP para operaciones de jugadores.
// Actúa como un adaptador entre la capa HTTP (Gin) y la capa de servicio de aplicación,
// traduciendo peticiones/respuestas HTTP a/desde operaciones del dominio.
type Handler struct {
	// PlayerService es el servicio de aplicación usado para realizar operaciones de negocio de jugadores.
	// Implementa la interfaz PlayerService, manteniendo el handler desacoplado de
	// implementaciones específicas de servicios.
	PlayerService ports.PlayerService
}
