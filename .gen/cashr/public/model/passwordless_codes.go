//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PasswordlessCodes struct {
	AppID        string `sql:"primary_key" json:"app_id" xml:"app_id"`
	TenantID     string `sql:"primary_key" json:"tenant_id" xml:"tenant_id"`
	CodeID       string `sql:"primary_key" json:"code_id" xml:"code_id"`
	DeviceIDHash string `json:"device_id_hash" xml:"device_id_hash"`
	LinkCodeHash string `json:"link_code_hash" xml:"link_code_hash"`
	CreatedAt    int64  `json:"created_at" xml:"created_at"`
}
