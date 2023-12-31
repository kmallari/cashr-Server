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

var ThirdpartyUserToTenant = newThirdpartyUserToTenantTable("public", "thirdparty_user_to_tenant", "")

type thirdpartyUserToTenantTable struct {
	postgres.Table

	// Columns
	AppID            postgres.ColumnString
	TenantID         postgres.ColumnString
	UserID           postgres.ColumnString
	ThirdPartyID     postgres.ColumnString
	ThirdPartyUserID postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ThirdpartyUserToTenantTable struct {
	thirdpartyUserToTenantTable

	EXCLUDED thirdpartyUserToTenantTable
}

// AS creates new ThirdpartyUserToTenantTable with assigned alias
func (a ThirdpartyUserToTenantTable) AS(alias string) *ThirdpartyUserToTenantTable {
	return newThirdpartyUserToTenantTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ThirdpartyUserToTenantTable with assigned schema name
func (a ThirdpartyUserToTenantTable) FromSchema(schemaName string) *ThirdpartyUserToTenantTable {
	return newThirdpartyUserToTenantTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ThirdpartyUserToTenantTable with assigned table prefix
func (a ThirdpartyUserToTenantTable) WithPrefix(prefix string) *ThirdpartyUserToTenantTable {
	return newThirdpartyUserToTenantTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ThirdpartyUserToTenantTable with assigned table suffix
func (a ThirdpartyUserToTenantTable) WithSuffix(suffix string) *ThirdpartyUserToTenantTable {
	return newThirdpartyUserToTenantTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newThirdpartyUserToTenantTable(schemaName, tableName, alias string) *ThirdpartyUserToTenantTable {
	return &ThirdpartyUserToTenantTable{
		thirdpartyUserToTenantTable: newThirdpartyUserToTenantTableImpl(schemaName, tableName, alias),
		EXCLUDED:                    newThirdpartyUserToTenantTableImpl("", "excluded", ""),
	}
}

func newThirdpartyUserToTenantTableImpl(schemaName, tableName, alias string) thirdpartyUserToTenantTable {
	var (
		AppIDColumn            = postgres.StringColumn("app_id")
		TenantIDColumn         = postgres.StringColumn("tenant_id")
		UserIDColumn           = postgres.StringColumn("user_id")
		ThirdPartyIDColumn     = postgres.StringColumn("third_party_id")
		ThirdPartyUserIDColumn = postgres.StringColumn("third_party_user_id")
		allColumns             = postgres.ColumnList{AppIDColumn, TenantIDColumn, UserIDColumn, ThirdPartyIDColumn, ThirdPartyUserIDColumn}
		mutableColumns         = postgres.ColumnList{ThirdPartyIDColumn, ThirdPartyUserIDColumn}
	)

	return thirdpartyUserToTenantTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		AppID:            AppIDColumn,
		TenantID:         TenantIDColumn,
		UserID:           UserIDColumn,
		ThirdPartyID:     ThirdPartyIDColumn,
		ThirdPartyUserID: ThirdPartyUserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
