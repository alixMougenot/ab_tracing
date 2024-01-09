package db

import (
	"context"
	"fmt"

	"github.com/alixMougenot/ab_tracing/graph/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePlantReproductionMaterial(info model.PlantReproductionMaterialInput, ctx context.Context, pool *pgxpool.Pool) (string, error) {
	if info.AquisitionType == nil {
		return "", fmt.Errorf("aquisitionType cannot be nil")
	}
	if info.Visibility == nil {
		return "", fmt.Errorf("visibility cannot be nil")
	}
	if info.GerminationSource == nil {
		return "", fmt.Errorf("germinationSource cannot be nil")
	}
	if info.TreatmentSteps == nil {
		return "", fmt.Errorf("treatmentSteps cannot be nil")
	}
	if info.HarvestSource == nil {
		return "", fmt.Errorf("harvestSource cannot be nil")
	}
	if info.ProductionDate == nil {
		return "", fmt.Errorf("productionDate cannot be nil")
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
	if info.Name == nil {
		return "", fmt.Errorf("name cannot be nil")
	}
	if info.Type == nil {
		return "", fmt.Errorf("type cannot be nil")
	}
	if info.IsOrganic == nil {
		return "", fmt.Errorf("isOrganic cannot be nil")
	}

	row := pool.QueryRow(ctx, `INSERT INTO public.plant_reproduction_material 
	(aquisition_type, visibility, germination_source, treatment_steps, harvest_source,
	 production_date, name_latin, quantity, notes, aquisition_places, aquisition_bought,
	 \"name\", is_organic, unit, \"type\")
	 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	 RETURNING id`,
		info.AquisitionType, info.Visibility, info.GerminationSource, info.TreatmentSteps, info.HarvestSource,
		info.ProductionDate, info.LatinName, info.Quantity, info.Notes, info.AquisitionPlaces, info.AquisitionPurshaseInfo,
		info.Name, info.IsOrganic, info.Unit, info.Type)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdatePlantReproductionMaterial(id string, info model.PlantReproductionMaterialInput, ctx context.Context, pool *pgxpool.Pool) error {
	query := "UPDATE public.plant_reproduction_material SET"
	args := []interface{}{}
	i := 2 // $1 is the id

	if info.AquisitionType != nil {
		query += fmt.Sprintf(" aquisition_type = $%d,", i)
		args = append(args, info.AquisitionType)
		i++
	}
	if info.Visibility != nil {
		query += fmt.Sprintf(" visibility = $%d,", i)
		args = append(args, info.Visibility)
		i++
	}
	if info.GerminationSource != nil {
		query += fmt.Sprintf(" germination_source = $%d,", i)
		args = append(args, info.GerminationSource)
		i++
	}
	if info.TreatmentSteps != nil {
		query += fmt.Sprintf(" treatment_steps = $%d,", i)
		args = append(args, info.TreatmentSteps)
		i++
	}
	if info.HarvestSource != nil {
		query += fmt.Sprintf(" harvest_source = $%d,", i)
		args = append(args, info.HarvestSource)
		i++
	}
	if info.ProductionDate != nil {
		query += fmt.Sprintf(" production_date = $%d,", i)
		args = append(args, info.ProductionDate)
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
	if info.IsOrganic != nil {
		query += fmt.Sprintf(" is_organic = $%d,", i)
		args = append(args, info.IsOrganic)
		i++
	}
	if info.Type != nil {
		query += fmt.Sprintf(" \"type\" = $%d,", i)
		args = append(args, info.Type)
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

func DeletePlantReproductionMaterial(id string, ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "DELETE FROM public.plant_reproduction_material WHERE id = $1", id)
	return err
}
