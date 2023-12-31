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

var UserMetadata = newUserMetadataTable("public", "user_metadata", "")

type userMetadataTable struct {
	postgres.Table

	// Columns
	AppID        postgres.ColumnString
	UserID       postgres.ColumnString
	UserMetadata postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type UserMetadataTable struct {
	userMetadataTable

	EXCLUDED userMetadataTable
}

// AS creates new UserMetadataTable with assigned alias
func (a UserMetadataTable) AS(alias string) *UserMetadataTable {
	return newUserMetadataTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserMetadataTable with assigned schema name
func (a UserMetadataTable) FromSchema(schemaName string) *UserMetadataTable {
	return newUserMetadataTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserMetadataTable with assigned table prefix
func (a UserMetadataTable) WithPrefix(prefix string) *UserMetadataTable {
	return newUserMetadataTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserMetadataTable with assigned table suffix
func (a UserMetadataTable) WithSuffix(suffix string) *UserMetadataTable {
	return newUserMetadataTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserMetadataTable(schemaName, tableName, alias string) *UserMetadataTable {
	return &UserMetadataTable{
		userMetadataTable: newUserMetadataTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newUserMetadataTableImpl("", "excluded", ""),
	}
}

func newUserMetadataTableImpl(schemaName, tableName, alias string) userMetadataTable {
	var (
		AppIDColumn        = postgres.StringColumn("app_id")
		UserIDColumn       = postgres.StringColumn("user_id")
		UserMetadataColumn = postgres.StringColumn("user_metadata")
		allColumns         = postgres.ColumnList{AppIDColumn, UserIDColumn, UserMetadataColumn}
		mutableColumns     = postgres.ColumnList{UserMetadataColumn}
	)

	return userMetadataTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		AppID:        AppIDColumn,
		UserID:       UserIDColumn,
		UserMetadata: UserMetadataColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
