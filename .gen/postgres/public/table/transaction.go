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

var Transaction = newTransactionTable("public", "transaction", "")

type transactionTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnString
	UserID      postgres.ColumnString
	Amount      postgres.ColumnFloat
	Date        postgres.ColumnInteger
	CategoryID  postgres.ColumnString
	Description postgres.ColumnString
	IsExpense   postgres.ColumnBool
	CreatedAt   postgres.ColumnInteger
	UpdatedAt   postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TransactionTable struct {
	transactionTable

	EXCLUDED transactionTable
}

// AS creates new TransactionTable with assigned alias
func (a TransactionTable) AS(alias string) *TransactionTable {
	return newTransactionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TransactionTable with assigned schema name
func (a TransactionTable) FromSchema(schemaName string) *TransactionTable {
	return newTransactionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TransactionTable with assigned table prefix
func (a TransactionTable) WithPrefix(prefix string) *TransactionTable {
	return newTransactionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TransactionTable with assigned table suffix
func (a TransactionTable) WithSuffix(suffix string) *TransactionTable {
	return newTransactionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTransactionTable(schemaName, tableName, alias string) *TransactionTable {
	return &TransactionTable{
		transactionTable: newTransactionTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newTransactionTableImpl("", "excluded", ""),
	}
}

func newTransactionTableImpl(schemaName, tableName, alias string) transactionTable {
	var (
		IDColumn          = postgres.StringColumn("id")
		UserIDColumn      = postgres.StringColumn("user_id")
		AmountColumn      = postgres.FloatColumn("amount")
		DateColumn        = postgres.IntegerColumn("date")
		CategoryIDColumn  = postgres.StringColumn("category_id")
		DescriptionColumn = postgres.StringColumn("description")
		IsExpenseColumn   = postgres.BoolColumn("is_expense")
		CreatedAtColumn   = postgres.IntegerColumn("created_at")
		UpdatedAtColumn   = postgres.IntegerColumn("updated_at")
		allColumns        = postgres.ColumnList{IDColumn, UserIDColumn, AmountColumn, DateColumn, CategoryIDColumn, DescriptionColumn, IsExpenseColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns    = postgres.ColumnList{UserIDColumn, AmountColumn, DateColumn, CategoryIDColumn, DescriptionColumn, IsExpenseColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return transactionTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		UserID:      UserIDColumn,
		Amount:      AmountColumn,
		Date:        DateColumn,
		CategoryID:  CategoryIDColumn,
		Description: DescriptionColumn,
		IsExpense:   IsExpenseColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
