package usecase

import (
	"context"
	"fmt"

	"github.com/madasatya6/gin-gonic/internal/entity"
)


//repository setting
type SettingRepo struct{
	*TranslationUseCase
}

func (uc *TranslationUseCase) Setting() SettingRepo {
	setting := SettingRepo{uc}
	return setting
}

//Show all setting
func (uc *SettingRepo) All(ctx context.Context) ([]entity.Setting, error) {
	model := uc.repo.SettingModel()
	translations, err := model.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - Setting - All(): %w", err)
	}
	
	return translations, nil
}

func (uc *SettingRepo) AllWithPagination(ctx context.Context) ([]interface{}, error) {
	model := uc.repo.SettingModel()
	translations, err := model.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - Setting pagination - AllWithPagination(): %w", err)
	}

	return translations, nil
}

