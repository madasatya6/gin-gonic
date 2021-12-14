// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/madasatya6/gin-gonic/internal/entity"
	"github.com/madasatya6/gin-gonic/internal/usecase/repo"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation - repository.
	Translation interface {
		Translate(context.Context, entity.Translation) (entity.Translation, error)
		History(context.Context) ([]entity.Translation, error)
		Setting() SettingRepo
	}

	// TranslationRepo - model.
	TranslationRepo interface {
		Store(context.Context, entity.Translation) error
		GetHistory(context.Context) ([]entity.Translation, error)
		SettingModel() repo.Setting
	}

	// TranslationWebAPI -.
	TranslationWebAPI interface {
		Translate(entity.Translation) (entity.Translation, error)
	}
)
