package game

import (
	"context"
	"database/sql"

	"github.com/thecodingcouple/gren-con-api/pkg/game"
)

// Implements the pkg Repository interface using
type pgRepository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) pgRepository {
	return pgRepository{db}
}

func (r pgRepository) Get(ctx context.Context, id string) (game.Game, error) {
	//TODO: implement
	return game.Game{}, nil
}

func (r pgRepository) Create(ctx context.Context, game game.Game) error {
	//TODO: implement
	return nil
}

func (r pgRepository) Update(ctx context.Context, game game.Game) error {
	//TODO: implement
	return nil
}

func (r pgRepository) Delete(ctx context.Context, id string) error {
	//TODO: implement
	return nil
}

func (r pgRepository) Count(ctx context.Context) (int, error) {
	//TODO: implement
	return nil
}

func (r pgRepository) Query(ctx context.Context, offset, limit int) ([]game.Game, error) {
	//TODO: implement
	return nil
}
