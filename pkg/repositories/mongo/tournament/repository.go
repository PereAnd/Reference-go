// Package tournament proporciona la implementación del repositorio MongoDB para persistencia de torneos.
// Este paquete implementa la interfaz TournamentRepository, sirviendo como un adaptador conducido
// (salida) en la arquitectura hexagonal. Maneja todas las operaciones específicas de MongoDB
// para entidades de torneos, traduciendo entre modelos de dominio y documentos de base de datos.
package tournament

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/jairogloz/go-l/pkg/ports"
)

// Asegura que Repository implementa ports.TournamentRepository en tiempo de compilación.
var _ ports.TournamentRepository = &Repository{}

// Repository implementa la interfaz TournamentRepository usando MongoDB como mecanismo de persistencia.
// Proporciona métodos para insertar, recuperar y eliminar documentos de torneos en MongoDB.
type Repository struct {
	// Client es el cliente de MongoDB usado para operaciones de base de datos.
	Client *mongo.Client

	// Collection es la colección de MongoDB donde se almacenan los documentos de torneos.
	Collection *mongo.Collection
}
