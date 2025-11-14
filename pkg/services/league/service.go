// Package league proporciona la implementación del servicio de aplicación para operaciones de ligas.
// Este paquete implementa la interfaz LeagueService, conteniendo la lógica de negocio
// y orquestación para operaciones relacionadas con ligas. Actúa como la capa de aplicación
// en la arquitectura hexagonal, coordinando entre las capas de dominio e infraestructura.
package league

import "github.com/jairogloz/go-l/pkg/ports"

// Asegura que Service implementa la interfaz LeagueService en tiempo de compilación.
var _ ports.LeagueService = &Service{}

// Service implementa la interfaz LeagueService, proporcionando lógica de negocio
// para operaciones de ligas. Orquesta operaciones del dominio y coordina
// con la capa de repositorio para la persistencia.
type Service struct {
	// Repo es el repositorio usado para persistir y recuperar datos de ligas.
	// Implementa la interfaz LeagueRepository, permitiendo que el servicio
	// permanezca desacoplado de implementaciones específicas de persistencia.
	Repo ports.LeagueRepository
}
