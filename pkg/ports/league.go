// Package ports define las interfaces (puertos) para la arquitectura hexagonal.
// Estas interfaces representan los contratos que el núcleo de la aplicación usa para interactuar
// con adaptadores externos. En la arquitectura hexagonal, los puertos definen qué necesita
// la aplicación, no cómo está implementado.
package ports

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// LeagueService define la interfaz del servicio de aplicación para operaciones de ligas.
// Este es un puerto conductor (entrada) que define las operaciones de lógica de negocio
// disponibles para ligas. Las implementaciones de esta interfaz deben contener
// reglas de negocio y orquestar operaciones del dominio.
type LeagueService interface {
	// Create crea una nueva liga en el sistema.
	// Valida los datos de la liga y aplica reglas de negocio antes de persistir.
	// Retorna un error si la liga no puede ser creada (ej., clave duplicada, fallo de validación).
	Create(ctx context.Context, league *domain.League) (err error)

	// Get obtiene una liga por su ID.
	// Retorna la liga si se encuentra, o un error si no se encuentra o si la operación falla.
	Get(ctx context.Context, id string) (league *domain.League, err error)

	// Delete elimina una liga del sistema por su ID.
	// Retorna un error si la liga no se encuentra o si la eliminación falla.
	Delete(ctx context.Context, id string) (err error)
}

// LeagueRepository define la interfaz del repositorio para operaciones de persistencia de ligas.
// Este es un puerto conducido (salida) que abstrae las operaciones de acceso a datos.
// Las implementaciones de esta interfaz manejan el mecanismo de persistencia real
// (ej., MongoDB, PostgreSQL, en memoria).
type LeagueRepository interface {
	// Insert persiste una nueva liga en el almacén de datos.
	// El ID de la liga debe ser generado si no se proporciona.
	// Retorna un error si la inserción falla (ej., clave duplicada, error de conexión).
	Insert(ctx context.Context, league *domain.League) (err error)

	// Get obtiene una liga del almacén de datos por su ID.
	// Retorna la liga si se encuentra, o un error si no se encuentra o si la operación falla.
	Get(ctx context.Context, id string) (league *domain.League, err error)

	// Delete elimina una liga del almacén de datos por su ID.
	// Retorna un error si la liga no se encuentra o si la eliminación falla.
	Delete(ctx context.Context, id string) (err error)
}
