package game

import (
	"context"
)

// Repository encapsulates the logic to access games from the db.
type Repository interface {
	// Get game by ID.
	Get(ctx context.Context, id string) (Game, error)
	// Returns the number of games.
	Count(ctx context.Context) (int, error)
	// Returns the list of games with the given offset and limit.
	Query(ctx context.Context, offset, limit int) ([]Game, error)
	// Adds a new game to the DB.
	Create(ctx context.Context, game Game) error
	// Updates the game with given ID in the db.
	Update(ctx context.Context, game Game) error
	// Removes the game with given ID from the db.
	Delete(ctx context.Context, id string) error
}
