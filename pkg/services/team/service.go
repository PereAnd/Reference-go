// Package team proporciona la implementación del servicio de aplicación para operaciones de equipos.
// Este paquete implementa la interfaz TeamService, conteniendo la lógica de negocio
// y orquestación para operaciones relacionadas con equipos. Actúa como la capa de aplicación
// en la arquitectura hexagonal, coordinando entre las capas de dominio e infraestructura.
package team

import (
	"github.com/jairogloz/go-l/pkg/ports"
)

// Asegura que Service implementa la interfaz TeamService en tiempo de compilación.
var _ ports.TeamService = &Service{}

// Service implementa la interfaz TeamService, proporcionando lógica de negocio
// para operaciones de equipos. Orquesta operaciones del dominio y coordina
// con la capa de repositorio para la persistencia.
type Service struct {
	// Repo es el repositorio usado para persistir y recuperar datos de equipos.
	// Implementa la interfaz TeamRepository, permitiendo que el servicio
	// permanezca desacoplado de implementaciones específicas de persistencia.
	Repo ports.TeamRepository
}
