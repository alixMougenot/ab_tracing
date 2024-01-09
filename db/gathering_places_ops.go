package db

import (
	"context"
	"fmt"

	"github.com/alixMougenot/ab_tracing/graph/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateGatheringPlace(info model.GatheringPlaceInput, ctx context.Context, pool *pgxpool.Pool) (string, error) {
	if info.Name == nil {
		return "", fmt.Errorf("name cannot be nil")
	}
	if info.Notes == nil {
		return "", fmt.Errorf("notes cannot be nil")
	}
	if info.Address == nil {
		return "", fmt.Errorf("address cannot be nil")
	}
	if info.Country == nil {
		return "", fmt.Errorf("country cannot be nil")
	}
	if info.IsOrganicCompliant == nil {
		return "", fmt.Errorf("isOrganicCompliant cannot be nil")
	}

	row := pool.QueryRow(ctx, `INSERT INTO public.gathering_place 
	(\"name\", notes, address, country, is_organic_compliant)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		info.Name, info.Notes, info.Address, info.Country, info.IsOrganicCompliant)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdateGatheringPlace(id string, info model.GatheringPlaceInput, ctx context.Context, pool *pgxpool.Pool) error {
	query := "UPDATE public.gathering_place SET"
	args := []interface{}{}
	i := 2 // $1 is the id

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
	if info.Address != nil {
		query += fmt.Sprintf(" address = $%d,", i)
		args = append(args, info.Address)
		i++
	}
	if info.Country != nil {
		query += fmt.Sprintf(" country = $%d,", i)
		args = append(args, info.Country)
		i++
	}
	if info.IsOrganicCompliant != nil {
		query += fmt.Sprintf(" is_organic_compliant = $%d,", i)
		args = append(args, info.IsOrganicCompliant)
		i++
	}

	// If no fields to update, return early
	if len(args) == 1 {
		return fmt.Errorf("no fields to update")
	}

	// Remove the last comma and add the WHERE clause
	query = query[:len(query)-1]
	query = query + " WHERE id = $1"
	args = append(args, id)

	_, err := pool.Exec(ctx, query, args...)
	return err
}

func DeleteGatheringPlace(id string, ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "DELETE FROM public.gathering_place WHERE id = $1", id)
	return err
}
