package db

import (
	"context"
	"fmt"

	"github.com/alixMougenot/ab_tracing/graph/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePlant(info model.PlantInput, ctx context.Context, pool *pgxpool.Pool) (string, error) {
	var visibility string
	var aquisitionType string
	var err error

	if info.AquisitionType == nil {
		return "", fmt.Errorf("aquisitionType cannot be nil")
	} else {
		aquisitionType, err = model.AquisitionType.ToPG(*info.AquisitionType)
		if err != nil {
			return "", err
		}
	}
	if info.Visibility == nil {
		return "", fmt.Errorf("visibility cannot be nil")
	} else {
		visibility, err = model.Visibility.ToPG(*info.Visibility)
		if err != nil {
			return "", err
		}
	}
	if info.TreatmentSteps == nil {
		return "", fmt.Errorf("treatmentSteps cannot be nil")
	}
	if info.Name == nil {
		return "", fmt.Errorf("name cannot be nil")
	}
	if info.LatinName == nil {
		return "", fmt.Errorf("nameLatin cannot be nil")
	}
	if info.Quantity == nil {
		return "", fmt.Errorf("quantity cannot be nil")
	}
	if info.Unit == nil {
		return "", fmt.Errorf("unit cannot be nil")
	}
	if info.Notes == nil {
		return "", fmt.Errorf("notes cannot be nil")
	}
	if info.AquisitionPlaces == nil {
		return "", fmt.Errorf("aquisitionPlaces cannot be nil")
	}
	if info.AquisitionPurshaseInfo == nil {
		return "", fmt.Errorf("aquisition purachase information cannot be nil")
	}
	if info.IsOrganic == nil {
		return "", fmt.Errorf("isOrganic cannot be nil")
	}
	if info.IsStockPlant == nil {
		return "", fmt.Errorf("isStockPlant cannot be nil")
	}

	row := pool.QueryRow(ctx, `INSERT INTO public.plant
	(aquisition_type, visibility, "source", grafting_sources,
	 maturation_sources, treatment_sources, planting_date, name_latin,
	quantity, notes, aquisition_places, aquisition_bought, is_stock_plant,
	"name", is_organic, unit)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) RETURNING id`,
		aquisitionType, visibility,
		info.PlantingSource, info.GraftingSteps,
		info.MaturationSteps, info.TreatmentSteps, info.PlantingDate, info.LatinName,
		info.Quantity, info.Notes, info.AquisitionPlaces, info.AquisitionPurshaseInfo, info.IsStockPlant,
		info.Name, info.IsOrganic, info.Unit)

	var id string
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdatePlant(id string, info model.PlantInput, ctx context.Context, pool *pgxpool.Pool) error {
	query := "UPDATE public.plant SET"
	args := []interface{}{}
	args = append(args, id)
	i := 2 // $1 is the id

	if info.AquisitionType != nil {
		query += fmt.Sprintf(" aquisition_type = $%d,", i)
		aquisitionType, err := model.AquisitionType.ToPG(*info.AquisitionType)
		if err != nil {
			return err
		}
		args = append(args, aquisitionType)
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
	if info.PlantingSource != nil {
		query += fmt.Sprintf(" \"source\" = $%d,", i)
		args = append(args, info.PlantingSource)
		i++
	}
	if info.GraftingSteps != nil {
		query += fmt.Sprintf(" grafting_sources = $%d,", i)
		args = append(args, info.GraftingSteps)
		i++
	}
	if info.MaturationSteps != nil {
		query += fmt.Sprintf(" maturation_sources = $%d,", i)
		args = append(args, info.MaturationSteps)
		i++
	}
	if info.TreatmentSteps != nil {
		query += fmt.Sprintf(" treatment_sources = $%d,", i)
		args = append(args, info.TreatmentSteps)
		i++
	}
	if info.PlantingDate != nil {
		query += fmt.Sprintf(" planting_date = $%d,", i)
		args = append(args, info.PlantingDate)
		i++
	}
	if info.Name != nil {
		query += fmt.Sprintf(" \"name\" = $%d,", i)
		args = append(args, info.Name)
		i++
	}
	if info.LatinName != nil {
		query += fmt.Sprintf(" name_latin = $%d,", i)
		args = append(args, info.LatinName)
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
	if info.Notes != nil {
		query += fmt.Sprintf(" notes = $%d,", i)
		args = append(args, info.Notes)
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
	if info.IsStockPlant != nil {
		query += fmt.Sprintf(" is_stock_plant = $%d,", i)
		args = append(args, info.IsStockPlant)
		i++
	}
	if info.IsOrganic != nil {
		query += fmt.Sprintf(" is_organic = $%d,", i)
		args = append(args, info.IsOrganic)
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

func DeletePlant(id string, ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "DELETE FROM public.plant WHERE id = $1", id)
	return err
}

func GetPlant(id string, ctx context.Context, pool *pgxpool.Pool) (*model.Plant, error) {
	row := pool.QueryRow(ctx, `SELECT 
	(id, aquisition_type, visibility, "source", grafting_sources,
	maturation_sources, treatment_sources, planting_date, name_latin,
  quantity, notes, aquisition_places, aquisition_bought, is_stock_plant,
  "name", is_organic, unit)
  FROM public.plant WHERE id = $1`, id)

	var plant model.Plant
	var visibility string
	var aquisitionType string
	err := row.Scan(&plant.ID, &aquisitionType, &visibility, &plant.PlantingSource,
		&plant.GraftingSteps, &plant.MaturationSteps, &plant.TreatmentSteps, &plant.PlantingDate,
		&plant.LatinName, &plant.Quantity, &plant.Notes, &plant.AquisitionPlaces,
		&plant.AquisitionPurshaseInfo, &plant.IsStockPlant, &plant.Name, &plant.IsOrganic, &plant.Unit)
	if err != nil {
		return nil, err
	}

	err = plant.Visibility.FromPG(visibility)
	if err != nil {
		return nil, err
	}

	err = plant.AquisitionType.FromPG(aquisitionType)
	if err != nil {
		return nil, err
	}

	return &plant, nil
}

func ListPlant(ctx context.Context, pool *pgxpool.Pool) ([]*model.Plant, error) {
	rows, err := pool.Query(ctx, `SELECT 
	id, aquisition_type, visibility, "source", grafting_sources,
	maturation_sources, treatment_sources, planting_date, name_latin,
	quantity, notes, aquisition_places, aquisition_bought, is_stock_plant,
	"name", is_organic, unit
	FROM public.plant 
	ORDER BY planting_date DESC;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plants := make([]*model.Plant, 0, 10)
	for rows.Next() {
		var plant model.Plant
		var visibility string
		var aquisitionType string
		err := rows.Scan(&plant.ID, &aquisitionType, &visibility, &plant.PlantingSource,
			&plant.GraftingSteps, &plant.MaturationSteps, &plant.TreatmentSteps, &plant.PlantingDate,
			&plant.LatinName, &plant.Quantity, &plant.Notes, &plant.AquisitionPlaces,
			&plant.AquisitionPurshaseInfo, &plant.IsStockPlant, &plant.Name, &plant.IsOrganic, &plant.Unit)
		if err != nil {
			return nil, err
		}

		err = plant.Visibility.FromPG(visibility)
		if err != nil {
			return nil, err
		}

		err = plant.AquisitionType.FromPG(aquisitionType)
		if err != nil {
			return nil, err
		}
		plants = append(plants, &plant)
	}

	return plants, nil
}
