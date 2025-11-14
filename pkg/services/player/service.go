// Package player proporciona la implementación del servicio de aplicación para operaciones de jugadores.
// Este paquete implementa la interfaz PlayerService, conteniendo la lógica de negocio
// y orquestación para operaciones relacionadas con jugadores. Actúa como la capa de aplicación
// en la arquitectura hexagonal, coordinando entre las capas de dominio e infraestructura.
package player

import (
	"github.com/jairogloz/go-l/pkg/ports"
)

// Asegura que Service implementa la interfaz PlayerService en tiempo de compilación.
var _ ports.PlayerService = &Service{}

// Service implementa la interfaz PlayerService, proporcionando lógica de negocio
// para operaciones de jugadores. Orquesta operaciones del dominio y coordina
// con la capa de repositorio para la persistencia.
type Service struct {
	// Repo es el repositorio usado para persistir y recuperar datos de jugadores.
	// Implementa la interfaz PlayerRepository, permitiendo que el servicio
	// permanezca desacoplado de implementaciones específicas de persistencia.
	Repo ports.PlayerRepository
}
