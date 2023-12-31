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

var Tenants = newTenantsTable("public", "tenants", "")

type tenantsTable struct {
	postgres.Table

	// Columns
	AppID         postgres.ColumnString
	TenantID      postgres.ColumnString
	CreatedAtTime postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TenantsTable struct {
	tenantsTable

	EXCLUDED tenantsTable
}

// AS creates new TenantsTable with assigned alias
func (a TenantsTable) AS(alias string) *TenantsTable {
	return newTenantsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TenantsTable with assigned schema name
func (a TenantsTable) FromSchema(schemaName string) *TenantsTable {
	return newTenantsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TenantsTable with assigned table prefix
func (a TenantsTable) WithPrefix(prefix string) *TenantsTable {
	return newTenantsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TenantsTable with assigned table suffix
func (a TenantsTable) WithSuffix(suffix string) *TenantsTable {
	return newTenantsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTenantsTable(schemaName, tableName, alias string) *TenantsTable {
	return &TenantsTable{
		tenantsTable: newTenantsTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newTenantsTableImpl("", "excluded", ""),
	}
}

func newTenantsTableImpl(schemaName, tableName, alias string) tenantsTable {
	var (
		AppIDColumn         = postgres.StringColumn("app_id")
		TenantIDColumn      = postgres.StringColumn("tenant_id")
		CreatedAtTimeColumn = postgres.IntegerColumn("created_at_time")
		allColumns          = postgres.ColumnList{AppIDColumn, TenantIDColumn, CreatedAtTimeColumn}
		mutableColumns      = postgres.ColumnList{CreatedAtTimeColumn}
	)

	return tenantsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		AppID:         AppIDColumn,
		TenantID:      TenantIDColumn,
		CreatedAtTime: CreatedAtTimeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
