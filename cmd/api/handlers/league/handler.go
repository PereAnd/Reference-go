// Package league proporciona manejadores HTTP para endpoints relacionados con ligas.
// Este paquete implementa los adaptadores conductores (entrada) en la arquitectura hexagonal,
// traduciendo peticiones HTTP en llamadas a servicios y respuestas HTTP. Usa el framework
// Gin para el enrutamiento HTTP y manejo de peticiones.
package league

import "github.com/jairogloz/go-l/pkg/ports"

// Handler proporciona manejadores HTTP para operaciones de ligas.
// Actúa como un adaptador entre la capa HTTP (Gin) y la capa de servicio de aplicación,
// traduciendo peticiones/respuestas HTTP a/desde operaciones del dominio.
type Handler struct {
	// LeagueService es el servicio de aplicación usado para realizar operaciones de negocio de ligas.
	// Implementa la interfaz LeagueService, manteniendo el handler desacoplado de
	// implementaciones específicas de servicios.
	LeagueService ports.LeagueService
}
