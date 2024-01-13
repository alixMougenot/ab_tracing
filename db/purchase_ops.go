package db

import (
	"context"
	"fmt"

	"github.com/alixMougenot/ab_tracing/graph/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateSupplyInfo(info model.SupplyInfoInput, ctx context.Context, pool *pgxpool.Pool) (string, error) {
	var err error
	var visibility string

	if info.Visibility == nil {
		return "", fmt.Errorf("visibility cannot be nil")
	} else {
		visibility, err = model.Visibility.ToPG(*info.Visibility)
		if err != nil {
			return "", err
		}
	}
	if info.Name == nil {
		return "", fmt.Errorf("name cannot be nil")
	}
	if info.Country == nil {
		return "", fmt.Errorf("country cannot be nil")
	}
	if info.Supplier == nil {
		return "", fmt.Errorf("supplier cannot be nil")
	}
	if info.Bill == nil {
		return "", fmt.Errorf("bill cannot be nil")
	}
	if info.Notes == nil {
		return "", fmt.Errorf("notes cannot be nil")
	}

	row := pool.QueryRow(ctx, `INSERT INTO public.supply_info 
	(visibility, "name", country, supplier, bill_link, notes)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`, visibility,
		info.Name, info.Country, info.Supplier, info.Bill, info.Notes)

	var retid = ""
	err = row.Scan(&retid)

	if err != nil {
		return "", err
	}

	return retid, nil
}

func UpdateSupplyInfo(id string, info model.SupplyInfoInput, ctx context.Context, pool *pgxpool.Pool) error {
	query := "UPDATE public.supply_info SET"
	args := make([]interface{}, 0, 10)
	args = append(args, id)
	i := 2 // $1 is the id

	if info.Visibility != nil {
		query += fmt.Sprintf(" visibility = $%d,", i)
		visibility, err := model.Visibility.ToPG(*info.Visibility)
		if err != nil {
			return err
		}
		args = append(args, visibility)
		i++
	}
	if info.Name != nil {
		query += fmt.Sprintf(" \"name\" = $%d,", i)
		args = append(args, info.Name)
		i++
	}
	if info.Country != nil {
		query += fmt.Sprintf(" country = $%d,", i)
		args = append(args, info.Country)
		i++
	}
	if info.Supplier != nil {
		query += fmt.Sprintf(" supplier = $%d,", i)
		args = append(args, info.Supplier)
		i++
	}
	if info.Bill != nil {
		query += fmt.Sprintf(" bill_link = $%d,", i)
		args = append(args, info.Bill)
		i++
	}
	if info.Notes != nil {
		query += fmt.Sprintf(" notes = $%d,", i)
		args = append(args, info.Notes)
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

func DeleteSupplyInfo(id string, ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "DELETE FROM public.supply_info WHERE id = $1", id)
	return err
}

func GetSupplyInfo(id string, ctx context.Context, pool *pgxpool.Pool) (*model.SupplyInfo, error) {
	row := pool.QueryRow(ctx, "SELECT id, visibility, \"name\", country, supplier, bill_link, notes FROM public.supply_info WHERE id = $1", id)
	var ret model.SupplyInfo
	var visibility string
	err := row.Scan(&ret.ID, &visibility, &ret.Name, &ret.Country, &ret.Supplier, &ret.Bill, &ret.Notes)
	if err != nil {
		return nil, err
	}

	err = ret.Visibility.FromPG(visibility)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func ListSupplyInfos(ctx context.Context, pool *pgxpool.Pool) ([]*model.SupplyInfo, error) {
	rows, err := pool.Query(ctx, "SELECT id, visibility, \"name\", country, supplier, bill_link, notes FROM public.supply_info ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret := make([]*model.SupplyInfo, 0, 10)
	for rows.Next() {
		var tmp model.SupplyInfo
		var visibility string
		err := rows.Scan(&tmp.ID, &visibility, &tmp.Name, &tmp.Country, &tmp.Supplier, &tmp.Bill, &tmp.Notes)
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
