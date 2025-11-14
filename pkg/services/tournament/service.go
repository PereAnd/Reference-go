// Package tournament proporciona la implementación del servicio de aplicación para operaciones de torneos.
// Este paquete implementa la interfaz TournamentService, conteniendo la lógica de negocio
// y orquestación para operaciones relacionadas con torneos. Actúa como la capa de aplicación
// en la arquitectura hexagonal, coordinando entre las capas de dominio e infraestructura.
package tournament

import "github.com/jairogloz/go-l/pkg/ports"

// Asegura que Service implementa la interfaz TournamentService en tiempo de compilación.
var _ ports.TournamentService = &Service{}

// Service implementa la interfaz TournamentService, proporcionando lógica de negocio
// para operaciones de torneos. Orquesta operaciones del dominio y coordina
// con la capa de repositorio para la persistencia.
type Service struct {
	// Repo es el repositorio usado para persistir y recuperar datos de torneos.
	// Implementa la interfaz TournamentRepository, permitiendo que el servicio
	// permanezca desacoplado de implementaciones específicas de persistencia.
	Repo ports.TournamentRepository
}
