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

var AllAuthRecipeUsers = newAllAuthRecipeUsersTable("public", "all_auth_recipe_users", "")

type allAuthRecipeUsersTable struct {
	postgres.Table

	// Columns
	AppID                         postgres.ColumnString
	TenantID                      postgres.ColumnString
	UserID                        postgres.ColumnString
	PrimaryOrRecipeUserID         postgres.ColumnString
	IsLinkedOrIsAPrimaryUser      postgres.ColumnBool
	RecipeID                      postgres.ColumnString
	TimeJoined                    postgres.ColumnInteger
	PrimaryOrRecipeUserTimeJoined postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AllAuthRecipeUsersTable struct {
	allAuthRecipeUsersTable

	EXCLUDED allAuthRecipeUsersTable
}

// AS creates new AllAuthRecipeUsersTable with assigned alias
func (a AllAuthRecipeUsersTable) AS(alias string) *AllAuthRecipeUsersTable {
	return newAllAuthRecipeUsersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AllAuthRecipeUsersTable with assigned schema name
func (a AllAuthRecipeUsersTable) FromSchema(schemaName string) *AllAuthRecipeUsersTable {
	return newAllAuthRecipeUsersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AllAuthRecipeUsersTable with assigned table prefix
func (a AllAuthRecipeUsersTable) WithPrefix(prefix string) *AllAuthRecipeUsersTable {
	return newAllAuthRecipeUsersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AllAuthRecipeUsersTable with assigned table suffix
func (a AllAuthRecipeUsersTable) WithSuffix(suffix string) *AllAuthRecipeUsersTable {
	return newAllAuthRecipeUsersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAllAuthRecipeUsersTable(schemaName, tableName, alias string) *AllAuthRecipeUsersTable {
	return &AllAuthRecipeUsersTable{
		allAuthRecipeUsersTable: newAllAuthRecipeUsersTableImpl(schemaName, tableName, alias),
		EXCLUDED:                newAllAuthRecipeUsersTableImpl("", "excluded", ""),
	}
}

func newAllAuthRecipeUsersTableImpl(schemaName, tableName, alias string) allAuthRecipeUsersTable {
	var (
		AppIDColumn                         = postgres.StringColumn("app_id")
		TenantIDColumn                      = postgres.StringColumn("tenant_id")
		UserIDColumn                        = postgres.StringColumn("user_id")
		PrimaryOrRecipeUserIDColumn         = postgres.StringColumn("primary_or_recipe_user_id")
		IsLinkedOrIsAPrimaryUserColumn      = postgres.BoolColumn("is_linked_or_is_a_primary_user")
		RecipeIDColumn                      = postgres.StringColumn("recipe_id")
		TimeJoinedColumn                    = postgres.IntegerColumn("time_joined")
		PrimaryOrRecipeUserTimeJoinedColumn = postgres.IntegerColumn("primary_or_recipe_user_time_joined")
		allColumns                          = postgres.ColumnList{AppIDColumn, TenantIDColumn, UserIDColumn, PrimaryOrRecipeUserIDColumn, IsLinkedOrIsAPrimaryUserColumn, RecipeIDColumn, TimeJoinedColumn, PrimaryOrRecipeUserTimeJoinedColumn}
		mutableColumns                      = postgres.ColumnList{PrimaryOrRecipeUserIDColumn, IsLinkedOrIsAPrimaryUserColumn, RecipeIDColumn, TimeJoinedColumn, PrimaryOrRecipeUserTimeJoinedColumn}
	)

	return allAuthRecipeUsersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		AppID:                         AppIDColumn,
		TenantID:                      TenantIDColumn,
		UserID:                        UserIDColumn,
		PrimaryOrRecipeUserID:         PrimaryOrRecipeUserIDColumn,
		IsLinkedOrIsAPrimaryUser:      IsLinkedOrIsAPrimaryUserColumn,
		RecipeID:                      RecipeIDColumn,
		TimeJoined:                    TimeJoinedColumn,
		PrimaryOrRecipeUserTimeJoined: PrimaryOrRecipeUserTimeJoinedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}