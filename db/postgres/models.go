// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package db

import ()

type Category struct {
	ID        string  `json:"id"`
	UserID    string  `json:"user_id"`
	Label     string  `json:"label"`
	Icon      *string `json:"icon"`
	IsExpense bool    `json:"is_expense"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}

type Transaction struct {
	ID          string  `json:"id"`
	UserID      string  `json:"user_id"`
	Amount      string  `json:"amount"`
	Date        int64   `json:"date"`
	CategoryID  string  `json:"category_id"`
	Description *string `json:"description"`
	IsExpense   bool    `json:"is_expense"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}
