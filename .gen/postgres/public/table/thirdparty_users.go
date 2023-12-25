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

var ThirdpartyUsers = newThirdpartyUsersTable("public", "thirdparty_users", "")

type thirdpartyUsersTable struct {
	postgres.Table

	// Columns
	AppID            postgres.ColumnString
	ThirdPartyID     postgres.ColumnString
	ThirdPartyUserID postgres.ColumnString
	UserID           postgres.ColumnString
	Email            postgres.ColumnString
	TimeJoined       postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ThirdpartyUsersTable struct {
	thirdpartyUsersTable

	EXCLUDED thirdpartyUsersTable
}

// AS creates new ThirdpartyUsersTable with assigned alias
func (a ThirdpartyUsersTable) AS(alias string) *ThirdpartyUsersTable {
	return newThirdpartyUsersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ThirdpartyUsersTable with assigned schema name
func (a ThirdpartyUsersTable) FromSchema(schemaName string) *ThirdpartyUsersTable {
	return newThirdpartyUsersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ThirdpartyUsersTable with assigned table prefix
func (a ThirdpartyUsersTable) WithPrefix(prefix string) *ThirdpartyUsersTable {
	return newThirdpartyUsersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ThirdpartyUsersTable with assigned table suffix
func (a ThirdpartyUsersTable) WithSuffix(suffix string) *ThirdpartyUsersTable {
	return newThirdpartyUsersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newThirdpartyUsersTable(schemaName, tableName, alias string) *ThirdpartyUsersTable {
	return &ThirdpartyUsersTable{
		thirdpartyUsersTable: newThirdpartyUsersTableImpl(schemaName, tableName, alias),
		EXCLUDED:             newThirdpartyUsersTableImpl("", "excluded", ""),
	}
}

func newThirdpartyUsersTableImpl(schemaName, tableName, alias string) thirdpartyUsersTable {
	var (
		AppIDColumn            = postgres.StringColumn("app_id")
		ThirdPartyIDColumn     = postgres.StringColumn("third_party_id")
		ThirdPartyUserIDColumn = postgres.StringColumn("third_party_user_id")
		UserIDColumn           = postgres.StringColumn("user_id")
		EmailColumn            = postgres.StringColumn("email")
		TimeJoinedColumn       = postgres.IntegerColumn("time_joined")
		allColumns             = postgres.ColumnList{AppIDColumn, ThirdPartyIDColumn, ThirdPartyUserIDColumn, UserIDColumn, EmailColumn, TimeJoinedColumn}
		mutableColumns         = postgres.ColumnList{ThirdPartyIDColumn, ThirdPartyUserIDColumn, EmailColumn, TimeJoinedColumn}
	)

	return thirdpartyUsersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		AppID:            AppIDColumn,
		ThirdPartyID:     ThirdPartyIDColumn,
		ThirdPartyUserID: ThirdPartyUserIDColumn,
		UserID:           UserIDColumn,
		Email:            EmailColumn,
		TimeJoined:       TimeJoinedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}