package db

import (
	"context"
	"fmt"

	"github.com/alixMougenot/ab_tracing/graph/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateGrowingMaterial(info model.GrowingMaterialInput, ctx context.Context, pool *pgxpool.Pool) (string, error) {
	var visibility string
	var aquisitionType string
	var err error

	if info.Name == nil {
		return "", fmt.Errorf("name cannot be nil")
	}
	if info.Notes == nil {
		return "", fmt.Errorf("notes cannot be nil")
	}
	if info.Visibility == nil {
		return "", fmt.Errorf("visibility cannot be nil")
	} else {
		visibility, err = model.Visibility.ToPG(*info.Visibility)
		if err != nil {
			return "", err
		}
	}
	if info.IsOrganicCompliant == nil {
		return "", fmt.Errorf("isOrganicCompliant cannot be nil")
	}
	if info.Quantity == nil {
		return "", fmt.Errorf("quantity cannot be nil")
	}
	if info.Unit == nil {
		return "", fmt.Errorf("unit cannot be nil")
	}
	if info.AquisitionType == nil {
		return "", fmt.Errorf("aquisitionType cannot be nil")
	} else {
		aquisitionType, err = model.AquisitionType.ToPG(*info.AquisitionType)
		if err != nil {
			return "", err
		}
	}
	if info.AquisitionPlaces == nil {
		return "", fmt.Errorf("aquisitionPlaces cannot be nil")
	}
	if info.AquisitionPurshaseInfo == nil {
		return "", fmt.Errorf("aquisitionPurshaseInfo cannot be nil")
	}
	if info.HomeProductionIngredients == nil {
		return "", fmt.Errorf("homeProductionIngredients cannot be nil")
	}
	if info.CreationDate == nil {
		return "", fmt.Errorf("creationDate cannot be nil")
	}

	row := pool.QueryRow(ctx, `INSERT INTO public.growing_material
	(creation_date, "name", notes, visibility, is_organic_compliant, quantity, unit, 
	aquisition_type, aquisition_places, aquisition_bought, production_steps) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
	RETURNING id`,
		info.CreationDate, info.Name, info.Notes, visibility,
		info.IsOrganicCompliant, info.Quantity, info.Unit, aquisitionType,
		info.AquisitionPlaces, info.AquisitionPurshaseInfo, info.HomeProductionIngredients)

	var id string
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdateGrowingMaterial(id string, info model.GrowingMaterialInput, ctx context.Context, pool *pgxpool.Pool) error {
	query := "UPDATE public.growing_material SET"
	args := []interface{}{}
	args = append(args, id)
	i := 2

	if info.CreationDate != nil {
		query += fmt.Sprintf(" creation_date = $%d,", i)
		args = append(args, info.CreationDate)
		i++
	}
	if info.Name != nil {
		query += fmt.Sprintf(" \"name\" = $%d,", i)
		args = append(args, info.Name)
		i++
	}
	if info.Notes != nil {
		query += fmt.Sprintf(" notes = $%d,", i)
		args = append(args, info.Notes)
		i++
	}
	if info.Visibility != nil {
		query += fmt.Sprintf(" visibility = $%d,", i)
		visibility, err := model.Visibility.ToPG(*info.Visibility)
		if err != nil {
			return err
		}
		args = append(args, visibility)
		i++
	}
	if info.IsOrganicCompliant != nil {
		query += fmt.Sprintf(" is_organic_compliant = $%d,", i)
		args = append(args, info.IsOrganicCompliant)
		i++
	}
	if info.Quantity != nil {
		query += fmt.Sprintf(" quantity = $%d,", i)
		args = append(args, info.Quantity)
		i++
	}
	if info.Unit != nil {
		query += fmt.Sprintf(" unit = $%d,", i)
		args = append(args, info.Unit)
		i++
	}
	if info.AquisitionType != nil {
		query += fmt.Sprintf(" aquisition_type = $%d,", i)
		aquisition, err := model.AquisitionType.ToPG(*info.AquisitionType)
		if err != nil {
			return err
		}
		args = append(args, aquisition)
		i++
	}
	if info.AquisitionPlaces != nil {
		query += fmt.Sprintf(" aquisition_places = $%d,", i)
		args = append(args, info.AquisitionPlaces)
		i++
	}
	if info.AquisitionPurshaseInfo != nil {
		query += fmt.Sprintf(" aquisition_bought = $%d,", i)
		args = append(args, info.AquisitionPurshaseInfo)
		i++
	}
	if info.HomeProductionIngredients != nil {
		query += fmt.Sprintf(" homeProductionIngredients = $%d,", i)
		args = append(args, info.HomeProductionIngredients)
		i++
	}

	// If no fields to update, return early
	if len(args) == 1 {
		return fmt.Errorf("no fields to update")
	}

	// Remove the last comma and add the WHERE clause
	query = query[:len(query)-1]
	query = query + " WHERE id = $1"

	_, err := pool.Exec(ctx, query, args...)
	return err
}

func DeleteGrowingMaterial(id string, ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "DELETE FROM public.growing_material WHERE id = $1", id)
	return err
}

func GetGrowingMaterial(id string, ctx context.Context, pool *pgxpool.Pool) (*model.GrowingMaterial, error) {
	row := pool.QueryRow(ctx, `SELECT
	id, creation_date, "name", notes, visibility, is_organic_compliant, 
	quantity, unit, aquisition_type, aquisition_places, aquisition_bought, 
  production_steps 
	FROM public.growing_material
	WHERE id = $1`, id)

	var growingMaterial model.GrowingMaterial
	var visibility string
	var aquisitionType string
	err := row.Scan(&growingMaterial.ID, &growingMaterial.CreationDate, &growingMaterial.Name,
		&growingMaterial.Notes, &visibility, &growingMaterial.IsOrganicCompliant,
		&growingMaterial.Quantity, &growingMaterial.Unit, &aquisitionType,
		&growingMaterial.AquisitionPlaces, &growingMaterial.AquisitionPurshaseInfo,
		&growingMaterial.HomeProductionIngredients)
	if err != nil {
		return nil, err
	}

	err = growingMaterial.Visibility.FromPG(visibility)
	if err != nil {
		return nil, err
	}

	err = growingMaterial.AquisitionType.FromPG(aquisitionType)
	if err != nil {
		return nil, err
	}

	return &growingMaterial, nil
}

func ListGrowingMaterials(ctx context.Context, pool *pgxpool.Pool) ([]*model.GrowingMaterial, error) {
	rows, err := pool.Query(ctx, `SELECT 
	id, creation_date, "name", notes, visibility, is_organic_compliant, quantity, unit,
	aquisition_type, aquisition_places, aquisition_bought, production_steps 
	FROM public.growing_material
	ORDER BY creation_date DESC;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	growingMaterials := make([]*model.GrowingMaterial, 0, 10)
	for rows.Next() {
		var growingMaterial model.GrowingMaterial
		var visibility string
		var aquisitionType string
		err := rows.Scan(&growingMaterial.ID, &growingMaterial.CreationDate, &growingMaterial.Name,
			&growingMaterial.Notes, &visibility, &growingMaterial.IsOrganicCompliant,
			&growingMaterial.Quantity, &growingMaterial.Unit, &aquisitionType,
			&growingMaterial.AquisitionPlaces, &growingMaterial.AquisitionPurshaseInfo,
			&growingMaterial.HomeProductionIngredients)
		if err != nil {
			return nil, err
		}

		err = growingMaterial.Visibility.FromPG(visibility)
		if err != nil {
			return nil, err
		}

		err = growingMaterial.AquisitionType.FromPG(aquisitionType)
		if err != nil {
			return nil, err
		}

		growingMaterials = append(growingMaterials, &growingMaterial)
	}
	return growingMaterials, nil
}
