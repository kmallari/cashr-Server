//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Category = newCategoryTable("public", "category", "")

type categoryTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnString
	UserID    postgres.ColumnString
	Label     postgres.ColumnString
	Icon      postgres.ColumnString
	IsExpense postgres.ColumnBool
	CreatedAt postgres.ColumnInteger
	UpdatedAt postgres.ColumnInteger
	LastUsed  postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CategoryTable struct {
	categoryTable

	EXCLUDED categoryTable
}

// AS creates new CategoryTable with assigned alias
func (a CategoryTable) AS(alias string) *CategoryTable {
	return newCategoryTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CategoryTable with assigned schema name
func (a CategoryTable) FromSchema(schemaName string) *CategoryTable {
	return newCategoryTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CategoryTable with assigned table prefix
func (a CategoryTable) WithPrefix(prefix string) *CategoryTable {
	return newCategoryTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CategoryTable with assigned table suffix
func (a CategoryTable) WithSuffix(suffix string) *CategoryTable {
	return newCategoryTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCategoryTable(schemaName, tableName, alias string) *CategoryTable {
	return &CategoryTable{
		categoryTable: newCategoryTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newCategoryTableImpl("", "excluded", ""),
	}
}

func newCategoryTableImpl(schemaName, tableName, alias string) categoryTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		UserIDColumn    = postgres.StringColumn("user_id")
		LabelColumn     = postgres.StringColumn("label")
		IconColumn      = postgres.StringColumn("icon")
		IsExpenseColumn = postgres.BoolColumn("is_expense")
		CreatedAtColumn = postgres.IntegerColumn("created_at")
		UpdatedAtColumn = postgres.IntegerColumn("updated_at")
		LastUsedColumn  = postgres.IntegerColumn("last_used")
		allColumns      = postgres.ColumnList{IDColumn, UserIDColumn, LabelColumn, IconColumn, IsExpenseColumn, CreatedAtColumn, UpdatedAtColumn, LastUsedColumn}
		mutableColumns  = postgres.ColumnList{UserIDColumn, LabelColumn, IconColumn, IsExpenseColumn, CreatedAtColumn, UpdatedAtColumn, LastUsedColumn}
	)

	return categoryTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		UserID:    UserIDColumn,
		Label:     LabelColumn,
		Icon:      IconColumn,
		IsExpense: IsExpenseColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		LastUsed:  LastUsedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
