//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type KeyValue struct {
	AppID         string `sql:"primary_key"`
	TenantID      string `sql:"primary_key"`
	Name          string `sql:"primary_key"`
	Value         *string
	CreatedAtTime *int64
}
