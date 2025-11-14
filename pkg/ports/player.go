package ports

import (
	"context"

	"github.com/jairogloz/go-l/pkg/domain"
)

// PlayerService define la interfaz del servicio de aplicación para operaciones de jugadores.
// Este es un puerto conductor (entrada) que define las operaciones de lógica de negocio
// disponibles para jugadores. Las implementaciones deben contener reglas de negocio y
// orquestar operaciones del dominio.
type PlayerService interface {
	// Create crea un nuevo jugador en el sistema.
	// Valida los datos del jugador y aplica reglas de negocio antes de persistir.
	// Retorna un error si el jugador no puede ser creado (ej., clave duplicada, fallo de validación).
	Create(ctx context.Context, player *domain.Player) (err error)

	// Get obtiene un jugador por su ID.
	// Retorna el jugador si se encuentra, o un error si no se encuentra o si la operación falla.
	Get(ctx context.Context, id string) (player *domain.Player, err error)

	// Delete elimina un jugador del sistema por su ID.
	// Retorna un error si el jugador no se encuentra o si la eliminación falla.
	Delete(ctx context.Context, id string) (err error)
}

// PlayerRepository define la interfaz del repositorio para operaciones de persistencia de jugadores.
// Este es un puerto conducido (salida) que abstrae las operaciones de acceso a datos.
// Las implementaciones manejan el mecanismo de persistencia real (ej., MongoDB, PostgreSQL).
type PlayerRepository interface {
	// Insert persiste un nuevo jugador en el almacén de datos.
	// El ID del jugador debe ser generado si no se proporciona.
	// Retorna un error si la inserción falla (ej., clave duplicada, error de conexión).
	Insert(ctx context.Context, player *domain.Player) (err error)

	// Get obtiene un jugador del almacén de datos por su ID.
	// Retorna el jugador si se encuentra, o un error si no se encuentra o si la operación falla.
	Get(ctx context.Context, id string) (player *domain.Player, err error)

	// GetPlayersByTeamID obtiene todos los jugadores que pertenecen a un equipo específico.
	// Retorna una lista de jugadores para el ID del equipo dado, o un error si la operación falla.
	// Retorna un error si no se encuentran jugadores para el equipo.
	GetPlayersByTeamID(ctx context.Context, id string) (players []*domain.Player, err error)

	// Delete elimina un jugador del almacén de datos por su ID.
	// Retorna un error si el jugador no se encuentra o si la eliminación falla.
	Delete(ctx context.Context, id string) (err error)
}
