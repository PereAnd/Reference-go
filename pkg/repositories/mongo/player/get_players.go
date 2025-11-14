package player

import (
	"context"
	"fmt"

	"github.com/jairogloz/go-l/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
)

// GetPlayersByTeamID obtiene todos los jugadores de la colección de MongoDB que pertenecen a un equipo con el ID de equipo proporcionado.
func (r *Repository) GetPlayersByTeamID(ctx context.Context, teamID string) ([]*domain.Player, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{"team_info.team_id": teamID})
	if err != nil {
		return nil, fmt.Errorf("error getting players: %w", err)
	}
	defer cursor.Close(ctx)

	var players []*domain.Player
	for cursor.Next(ctx) {
		var player domain.Player
		err := cursor.Decode(&player)
		if err != nil {
			return nil, fmt.Errorf("error decoding player: %w", err)
		}
		players = append(players, &player)
	}
	if len(players) == 0 {
		return nil, domain.ErrNotFound
	}

	return players, nil
}
