package repo

import (
	"context"
	"fmt"

	"github.com/madasatya6/gin-gonic/internal/entity"
)

// GetHistory -.
type Setting struct{
	*TranslationRepo
}

func (r *TranslationRepo) SettingModel() Setting {
	repo := Setting{r}
	return repo
}

func (r *Setting) GetAll(ctx context.Context) ([]entity.Setting, error) {
	sql, _, err := r.Builder.
		Select("id, name, value, field_type, tab").
		From("setting").
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetSetting - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetSetting - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Setting, 0)

	for rows.Next() {
		e := entity.Setting{}

		err = rows.Scan(&e.ID, &e.Name, &e.Value, &e.FieldType, &e.Tab)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetSetting - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

func (r *Setting) All(ctx context.Context) ([]interface{}, error) {
	//support pagination
	sql, _, err := r.Builder.
		Select("*").
		From("setting").
		ToSql()
	
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - All setting - r.Builder: %w", err)
	}
	
	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - All setting - r.PoolQuery: %w", err)
	}
	defer rows.Close()

	entities := make([]interface{}, 0)
	for rows.Next() {
		e := entity.Setting{}

		if err := rows.Scan(&e.ID, &e.Name, &e.Value, &e.FieldType, &e.Tab); err != nil {
			return nil, fmt.Errorf("TranslationRepo - All setting - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

