// Package team proporciona la implementación del repositorio MongoDB para persistencia de equipos.
// Este paquete implementa la interfaz TeamRepository, sirviendo como un adaptador conducido
// (salida) en la arquitectura hexagonal. Maneja todas las operaciones específicas de MongoDB
// para entidades de equipos, traduciendo entre modelos de dominio y documentos de base de datos.
package team

import (
	"github.com/jairogloz/go-l/pkg/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

// Asegura que Repository implementa ports.TeamRepository en tiempo de compilación.
var _ ports.TeamRepository = &Repository{}

// Repository implementa la interfaz TeamRepository usando MongoDB como mecanismo de persistencia.
// Proporciona métodos para insertar, recuperar y eliminar documentos de equipos en MongoDB.
type Repository struct {
	// Client es el cliente de MongoDB usado para operaciones de base de datos.
	Client *mongo.Client

	// Collection es la colección de MongoDB donde se almacenan los documentos de equipos.
	Collection *mongo.Collection
}
