package player

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jairogloz/go-l/pkg/domain"
)

// Create es un método de la estructura Service que crea un nuevo jugador en la base de datos.
//
// Parámetros:
//   - player: Un puntero a una estructura domain.Player. Esto permite al método modificar la estructura Player original.
//
// El método establece el campo CreatedAt de la estructura Player a la hora actual y luego llama al método Insert
// del campo Repo de la estructura Service, pasando la estructura Player como argumento.
//
// Si el método Insert retorna un error, el método Create verifica si el error es un error de clave duplicada.
// Si lo es, registra el error y retorna un nuevo AppError con el código y mensaje de error de clave duplicada.
// Si el error no es un error de clave duplicada, registra el error y retorna un nuevo error envolviendo el error
// original con un mensaje indicando que hubo un error al crear el jugador.
//
// Si el método Insert no retorna un error, el método Create establece el campo ID de la estructura Player
// al ID retornado por el método Insert. Esto permite al llamador del método Create acceder al ID del jugador recién creado.
//
// Valores de retorno:
//   - err: Un error que será nil si el jugador fue creado exitosamente. Si hubo un error, será un objeto de error
//     describiendo el fallo.
func (s *Service) Create(ctx context.Context, player *domain.Player) (err error) {
	now := time.Now().UTC()
	player.CreatedAt = &now

	err = s.Repo.Insert(ctx, player)
	if err != nil {
		if errors.Is(err, domain.ErrDuplicateKey) {
			log.Println("Duplicate key error")
			appErr := domain.AppError{
				Code: domain.ErrCodeDuplicateKey,
				Msg:  "error creating player: duplicate key error",
			}
			return appErr
		}
		log.Println(err.Error())
		return fmt.Errorf("error creating player: %w", err)
	}

	return nil
}
