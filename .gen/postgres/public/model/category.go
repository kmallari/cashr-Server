//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Category struct {
	ID        string `sql:"primary_key"`
	UserID    string
	Label     string
	Icon      *string
	IsExpense bool
	CreatedAt int64
	UpdatedAt int64
}
