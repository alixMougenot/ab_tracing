package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/alixMougenot/ab_tracing/graph/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePlantTreatment(info model.PlantTreatmentInput, ctx context.Context, pool *pgxpool.Pool) (string, error) {
	if info.AquisitionType == nil {
		return "", fmt.Errorf("aquisitionType cannot be nil")
	}
	if info.Visibility == nil {
		return "", fmt.Errorf("visibility cannot be nil")
	}
	if info.HomeProductionIngredients == nil {
		return "", fmt.Errorf("home production ingredients cannot be nil")
	}
	if info.CreationDate == nil {
		return "", fmt.Errorf("creation date cannot be nil")
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
		return "", fmt.Errorf("purchase cannot be nil")
	}
	if info.Name == nil {
		return "", fmt.Errorf("name cannot be nil")
	}
	if info.IsOrganicCompliant == nil {
		return "", fmt.Errorf("isOrganicCompatible cannot be nil")
	}

	row := pool.QueryRow(ctx, `INSERT INTO public.plant_treatments
	(aquisition_type, visibility, production_ingredients, creation_date,
	quantity, unit, notes, aquisition_places, aquisition_bought,
	\"name\", is_organic_compatible)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`,
		strings.ToLower(info.AquisitionType.String()), strings.ToLower(info.Visibility.String()),
		info.HomeProductionIngredients, info.CreationDate,
		info.Quantity, info.Unit, info.Notes, info.AquisitionPlaces, info.AquisitionPurshaseInfo,
		info.Name, info.IsOrganicCompliant)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdatePlantTreatment(id string, info model.PlantTreatmentInput, ctx context.Context, pool *pgxpool.Pool) error {
	query := "UPDATE public.plant_treatments SET"
	args := []interface{}{}
	args = append(args, id)
	i := 2 // $1 is the id

	if info.AquisitionType != nil {
		query += fmt.Sprintf(" aquisition_type = $%d,", i)
		args = append(args, strings.ToLower(info.AquisitionType.String()))
		i++
	}
	if info.Visibility != nil {
		query += fmt.Sprintf(" visibility = $%d,", i)
		args = append(args, strings.ToLower(info.Visibility.String()))
		i++
	}
	if info.HomeProductionIngredients != nil {
		query += fmt.Sprintf(" production_ingredients = $%d,", i)
		args = append(args, info.HomeProductionIngredients)
		i++
	}
	if info.CreationDate != nil {
		query += fmt.Sprintf(" creation_date = $%d,", i)
		args = append(args, info.CreationDate)
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
	if info.Name != nil {
		query += fmt.Sprintf(" \"name\" = $%d,", i)
		args = append(args, info.Name)
		i++
	}
	if info.IsOrganicCompliant != nil {
		query += fmt.Sprintf(" is_organic_compatible = $%d,", i)
		args = append(args, info.IsOrganicCompliant)
		i++
	}

	// If no fields to update, return early
	if len(args) == 1 {
		return fmt.Errorf("no fields to update")
	}

	// Remove the last comma and add the WHERE clause
	query = query[:len(query)-1]
	query = query + " WHERE id = $%d"

	_, err := pool.Exec(ctx, query, args...)
	return err
}

func DeletePlantTreatment(id string, ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "DELETE FROM public.plant_treatments WHERE id = $1", id)
	return err
}

func GetPlantTreatment(id string, ctx context.Context, pool *pgxpool.Pool) (*model.PlantTreatment, error) {
	row := pool.QueryRow(ctx, `SELECT 
	  (id, aquisition_type, visibility, production_ingredients, creation_date,
		quantity, unit, notes, aquisition_places, aquisition_bought,
		"name", is_organic_compatible)
		FROM public.plant_treatments
		WHERE id = $1`, id)
	var plantTreatment model.PlantTreatment
	var aquisitionType string
	var visibility string
	err := row.Scan(&plantTreatment.ID, &aquisitionType, &visibility,
		&plantTreatment.HomeProductionIngredients, &plantTreatment.CreationDate, &plantTreatment.Quantity,
		&plantTreatment.Unit, &plantTreatment.Notes, &plantTreatment.AquisitionPlaces,
		&plantTreatment.AquisitionPurshaseInfo, &plantTreatment.Name, &plantTreatment.IsOrganicCompliant)
	if err != nil {
		return nil, err
	}

	err = plantTreatment.Visibility.UnmarshalGQL(strings.ToUpper(visibility))
	if err != nil {
		return nil, err
	}

	err = plantTreatment.AquisitionType.UnmarshalGQL(strings.ToUpper(aquisitionType))
	if err != nil {
		return nil, err
	}

	return &plantTreatment, nil
}

func ListPlantTreatments(ctx context.Context, pool *pgxpool.Pool) ([]*model.PlantTreatment, error) {
	rows, err := pool.Query(ctx, `SELECT 
	  id, aquisition_type, visibility, production_ingredients, creation_date,
		quantity, unit, notes, aquisition_places, aquisition_bought,
		"name", is_organic_compatible
		FROM public.plant_treatments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	plantTreatments := make([]*model.PlantTreatment, 0, 10)
	for rows.Next() {
		var plantTreatment model.PlantTreatment
		var aquisitionType string
		var visibility string
		err = rows.Scan(&plantTreatment.ID, &aquisitionType, &visibility,
			&plantTreatment.HomeProductionIngredients, &plantTreatment.CreationDate, &plantTreatment.Quantity,
			&plantTreatment.Unit, &plantTreatment.Notes, &plantTreatment.AquisitionPlaces,
			&plantTreatment.AquisitionPurshaseInfo, &plantTreatment.Name, &plantTreatment.IsOrganicCompliant)
		if err != nil {
			return nil, err
		}
		err = plantTreatment.Visibility.UnmarshalGQL(strings.ToUpper(visibility))
		if err != nil {
			return nil, err
		}

		err = plantTreatment.AquisitionType.UnmarshalGQL(strings.ToUpper(aquisitionType))
		if err != nil {
			return nil, err
		}

		plantTreatments = append(plantTreatments, &plantTreatment)
	}
	return plantTreatments, nil
}
