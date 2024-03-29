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
	if info.Visibility == nil {
		return "", fmt.Errorf("visibility cannot be nil")
	}

	visibility, err := model.Visibility.ToPG(*info.Visibility)
	if err != nil {
		return "", err
	}

	row := pool.QueryRow(ctx, `INSERT INTO public.gathering_place 
	(name, notes, address, country, is_organic_compliant, visibility)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		info.Name, info.Notes, info.Address, info.Country, info.IsOrganicCompliant, visibility)

	var id string
	err = row.Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdateGatheringPlace(id string, info model.GatheringPlaceInput, ctx context.Context, pool *pgxpool.Pool) error {
	query := "UPDATE public.gathering_place SET"
	args := []interface{}{}
	args = append(args, id)
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
	if info.Visibility != nil {
		query += fmt.Sprintf(" visibility = $%d,", i)
		visibility, err := model.Visibility.ToPG(*info.Visibility)
		if err != nil {
			return err
		}
		args = append(args, visibility)
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

func DeleteGatheringPlace(id string, ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "DELETE FROM public.gather_places WHERE id = $1", id)
	return err
}

func GetGatheringPlace(id string, ctx context.Context, pool *pgxpool.Pool) (*model.GatheringPlace, error) {
	row := pool.QueryRow(ctx, `SELECT 
		"name", notes, address, country, is_organic_compatible, visibility
		FROM public.gather_places WHERE id = $1`, id)

	var ret model.GatheringPlace
	var visibility string
	err := row.Scan(&ret.Name, &ret.Notes, &ret.Address, &ret.Country, &ret.IsOrganicCompliant, &visibility)
	if err != nil {
		return nil, err
	}

	err = ret.Visibility.FromPG(visibility)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func ListGatheringPlaces(ctx context.Context, pool *pgxpool.Pool) ([]*model.GatheringPlace, error) {
	rows, err := pool.Query(ctx, `SELECT 
		"name", notes, address, country, is_organic_compatible, visibility
		FROM public.gather_places
		ORDER BY id DESC;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret := make([]*model.GatheringPlace, 0, 10)
	for rows.Next() {
		var tmp model.GatheringPlace
		var visibility string
		err := rows.Scan(&tmp.Name, &tmp.Notes, &tmp.Address, &tmp.Country, &tmp.IsOrganicCompliant, &visibility)
		if err != nil {
			return nil, err
		}

		err = tmp.Visibility.FromPG(visibility)
		if err != nil {
			return nil, err
		}

		ret = append(ret, &tmp)
	}

	return ret, nil
}
