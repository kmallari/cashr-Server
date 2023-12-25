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

var PasswordlessCodes = newPasswordlessCodesTable("public", "passwordless_codes", "")

type passwordlessCodesTable struct {
	postgres.Table

	// Columns
	AppID        postgres.ColumnString
	TenantID     postgres.ColumnString
	CodeID       postgres.ColumnString
	DeviceIDHash postgres.ColumnString
	LinkCodeHash postgres.ColumnString
	CreatedAt    postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PasswordlessCodesTable struct {
	passwordlessCodesTable

	EXCLUDED passwordlessCodesTable
}

// AS creates new PasswordlessCodesTable with assigned alias
func (a PasswordlessCodesTable) AS(alias string) *PasswordlessCodesTable {
	return newPasswordlessCodesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PasswordlessCodesTable with assigned schema name
func (a PasswordlessCodesTable) FromSchema(schemaName string) *PasswordlessCodesTable {
	return newPasswordlessCodesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PasswordlessCodesTable with assigned table prefix
func (a PasswordlessCodesTable) WithPrefix(prefix string) *PasswordlessCodesTable {
	return newPasswordlessCodesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PasswordlessCodesTable with assigned table suffix
func (a PasswordlessCodesTable) WithSuffix(suffix string) *PasswordlessCodesTable {
	return newPasswordlessCodesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPasswordlessCodesTable(schemaName, tableName, alias string) *PasswordlessCodesTable {
	return &PasswordlessCodesTable{
		passwordlessCodesTable: newPasswordlessCodesTableImpl(schemaName, tableName, alias),
		EXCLUDED:               newPasswordlessCodesTableImpl("", "excluded", ""),
	}
}

func newPasswordlessCodesTableImpl(schemaName, tableName, alias string) passwordlessCodesTable {
	var (
		AppIDColumn        = postgres.StringColumn("app_id")
		TenantIDColumn     = postgres.StringColumn("tenant_id")
		CodeIDColumn       = postgres.StringColumn("code_id")
		DeviceIDHashColumn = postgres.StringColumn("device_id_hash")
		LinkCodeHashColumn = postgres.StringColumn("link_code_hash")
		CreatedAtColumn    = postgres.IntegerColumn("created_at")
		allColumns         = postgres.ColumnList{AppIDColumn, TenantIDColumn, CodeIDColumn, DeviceIDHashColumn, LinkCodeHashColumn, CreatedAtColumn}
		mutableColumns     = postgres.ColumnList{DeviceIDHashColumn, LinkCodeHashColumn, CreatedAtColumn}
	)

	return passwordlessCodesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		AppID:        AppIDColumn,
		TenantID:     TenantIDColumn,
		CodeID:       CodeIDColumn,
		DeviceIDHash: DeviceIDHashColumn,
		LinkCodeHash: LinkCodeHashColumn,
		CreatedAt:    CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}