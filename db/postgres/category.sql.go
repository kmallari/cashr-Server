// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
insert into category (user_id, is_expense, label, icon)
values ($1, $2, $3, $4)
returning id, user_id, label, icon, is_expense, created_at, updated_at
`

type CreateCategoryParams struct {
	UserID    string  `json:"user_id"`
	IsExpense bool    `json:"is_expense"`
	Label     string  `json:"label"`
	Icon      *string `json:"icon"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory,
		arg.UserID,
		arg.IsExpense,
		arg.Label,
		arg.Icon,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Label,
		&i.Icon,
		&i.IsExpense,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
delete
from category
where id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getUserCategories = `-- name: GetUserCategories :many
select id, user_id, label, icon, is_expense, created_at, updated_at
from category
where user_id = $1
  and is_expense = $2
order by created_at desc
`

type GetUserCategoriesParams struct {
	UserID    string `json:"user_id"`
	IsExpense bool   `json:"is_expense"`
}

func (q *Queries) GetUserCategories(ctx context.Context, arg GetUserCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, getUserCategories, arg.UserID, arg.IsExpense)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Label,
			&i.Icon,
			&i.IsExpense,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :one
update category
set icon = coalesce($2, icon)
    and label = coalesce($3, label)
    and is_expense = coalesce($4, is_expense)
    and updated_at = $5
where id = $1
returning id, user_id, label, icon, is_expense, created_at, updated_at
`

type UpdateCategoryParams struct {
	ID        string  `json:"id"`
	Icon      *string `json:"icon"`
	Label     string  `json:"label"`
	IsExpense bool    `json:"is_expense"`
	UpdatedAt int64   `json:"updated_at"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, updateCategory,
		arg.ID,
		arg.Icon,
		arg.Label,
		arg.IsExpense,
		arg.UpdatedAt,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Label,
		&i.Icon,
		&i.IsExpense,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
