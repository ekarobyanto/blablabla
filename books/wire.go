//go:build wireinject
// +build wireinject

package books

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

func InitializeBookHandler(db *sqlx.DB) *BookHandler {
	wire.Build(
		NewBookHandler,
		NewBookService,
		NewBookRepo,
	)

	return &BookHandler{}
}
