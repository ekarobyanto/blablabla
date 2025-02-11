//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/username/mentoring_study_case/users"
)

func InitializeAuthHandler(db *sqlx.DB) *AuthHandler {
	wire.Build(
		NewAuthHandler,
		NewAuthService,
		users.NewUserService,
		users.NewUserRepo,
	)

	return &AuthHandler{}
}
