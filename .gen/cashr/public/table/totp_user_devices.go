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

var TotpUserDevices = newTotpUserDevicesTable("public", "totp_user_devices", "")

type totpUserDevicesTable struct {
	postgres.Table

	// Columns
	AppID      postgres.ColumnString
	UserID     postgres.ColumnString
	DeviceName postgres.ColumnString
	SecretKey  postgres.ColumnString
	Period     postgres.ColumnInteger
	Skew       postgres.ColumnInteger
	Verified   postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type TotpUserDevicesTable struct {
	totpUserDevicesTable

	EXCLUDED totpUserDevicesTable
}

// AS creates new TotpUserDevicesTable with assigned alias
func (a TotpUserDevicesTable) AS(alias string) *TotpUserDevicesTable {
	return newTotpUserDevicesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new TotpUserDevicesTable with assigned schema name
func (a TotpUserDevicesTable) FromSchema(schemaName string) *TotpUserDevicesTable {
	return newTotpUserDevicesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new TotpUserDevicesTable with assigned table prefix
func (a TotpUserDevicesTable) WithPrefix(prefix string) *TotpUserDevicesTable {
	return newTotpUserDevicesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new TotpUserDevicesTable with assigned table suffix
func (a TotpUserDevicesTable) WithSuffix(suffix string) *TotpUserDevicesTable {
	return newTotpUserDevicesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newTotpUserDevicesTable(schemaName, tableName, alias string) *TotpUserDevicesTable {
	return &TotpUserDevicesTable{
		totpUserDevicesTable: newTotpUserDevicesTableImpl(schemaName, tableName, alias),
		EXCLUDED:             newTotpUserDevicesTableImpl("", "excluded", ""),
	}
}

func newTotpUserDevicesTableImpl(schemaName, tableName, alias string) totpUserDevicesTable {
	var (
		AppIDColumn      = postgres.StringColumn("app_id")
		UserIDColumn     = postgres.StringColumn("user_id")
		DeviceNameColumn = postgres.StringColumn("device_name")
		SecretKeyColumn  = postgres.StringColumn("secret_key")
		PeriodColumn     = postgres.IntegerColumn("period")
		SkewColumn       = postgres.IntegerColumn("skew")
		VerifiedColumn   = postgres.BoolColumn("verified")
		allColumns       = postgres.ColumnList{AppIDColumn, UserIDColumn, DeviceNameColumn, SecretKeyColumn, PeriodColumn, SkewColumn, VerifiedColumn}
		mutableColumns   = postgres.ColumnList{SecretKeyColumn, PeriodColumn, SkewColumn, VerifiedColumn}
	)

	return totpUserDevicesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		AppID:      AppIDColumn,
		UserID:     UserIDColumn,
		DeviceName: DeviceNameColumn,
		SecretKey:  SecretKeyColumn,
		Period:     PeriodColumn,
		Skew:       SkewColumn,
		Verified:   VerifiedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
