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

var TenantThirdpartyProviderClients = newTenantThirdpartyProviderClientsTable("public", "tenant_thirdparty_provider_clients", "")

type tenantThirdpartyProviderClientsTable struct {
	postgres.Table

	// Columns
	ConnectionURIDomain postgres.ColumnString
	AppID               postgres.ColumnString
	TenantID            postgres.ColumnString
	ThirdPartyID        postgres.ColumnString
	ClientType          postgres.ColumnString
	ClientID            postgres.ColumnString
	ClientSecret        postgres.ColumnString
	Scope               postgres.ColumnString
	ForcePkce           postgres.ColumnBool
	AdditionalConfig    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TenantThirdpartyProviderClientsTable struct {
	tenantThirdpartyProviderClientsTable

	EXCLUDED tenantThirdpartyProviderClientsTable
}

// AS creates new TenantThirdpartyProviderClientsTable with assigned alias
func (a TenantThirdpartyProviderClientsTable) AS(alias string) *TenantThirdpartyProviderClientsTable {
	return newTenantThirdpartyProviderClientsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TenantThirdpartyProviderClientsTable with assigned schema name
func (a TenantThirdpartyProviderClientsTable) FromSchema(schemaName string) *TenantThirdpartyProviderClientsTable {
	return newTenantThirdpartyProviderClientsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TenantThirdpartyProviderClientsTable with assigned table prefix
func (a TenantThirdpartyProviderClientsTable) WithPrefix(prefix string) *TenantThirdpartyProviderClientsTable {
	return newTenantThirdpartyProviderClientsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TenantThirdpartyProviderClientsTable with assigned table suffix
func (a TenantThirdpartyProviderClientsTable) WithSuffix(suffix string) *TenantThirdpartyProviderClientsTable {
	return newTenantThirdpartyProviderClientsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTenantThirdpartyProviderClientsTable(schemaName, tableName, alias string) *TenantThirdpartyProviderClientsTable {
	return &TenantThirdpartyProviderClientsTable{
		tenantThirdpartyProviderClientsTable: newTenantThirdpartyProviderClientsTableImpl(schemaName, tableName, alias),
		EXCLUDED:                             newTenantThirdpartyProviderClientsTableImpl("", "excluded", ""),
	}
}

func newTenantThirdpartyProviderClientsTableImpl(schemaName, tableName, alias string) tenantThirdpartyProviderClientsTable {
	var (
		ConnectionURIDomainColumn = postgres.StringColumn("connection_uri_domain")
		AppIDColumn               = postgres.StringColumn("app_id")
		TenantIDColumn            = postgres.StringColumn("tenant_id")
		ThirdPartyIDColumn        = postgres.StringColumn("third_party_id")
		ClientTypeColumn          = postgres.StringColumn("client_type")
		ClientIDColumn            = postgres.StringColumn("client_id")
		ClientSecretColumn        = postgres.StringColumn("client_secret")
		ScopeColumn               = postgres.StringColumn("scope")
		ForcePkceColumn           = postgres.BoolColumn("force_pkce")
		AdditionalConfigColumn    = postgres.StringColumn("additional_config")
		allColumns                = postgres.ColumnList{ConnectionURIDomainColumn, AppIDColumn, TenantIDColumn, ThirdPartyIDColumn, ClientTypeColumn, ClientIDColumn, ClientSecretColumn, ScopeColumn, ForcePkceColumn, AdditionalConfigColumn}
		mutableColumns            = postgres.ColumnList{ClientIDColumn, ClientSecretColumn, ScopeColumn, ForcePkceColumn, AdditionalConfigColumn}
	)

	return tenantThirdpartyProviderClientsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ConnectionURIDomain: ConnectionURIDomainColumn,
		AppID:               AppIDColumn,
		TenantID:            TenantIDColumn,
		ThirdPartyID:        ThirdPartyIDColumn,
		ClientType:          ClientTypeColumn,
		ClientID:            ClientIDColumn,
		ClientSecret:        ClientSecretColumn,
		Scope:               ScopeColumn,
		ForcePkce:           ForcePkceColumn,
		AdditionalConfig:    AdditionalConfigColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}