//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type EmailpasswordUserToTenant struct {
	AppID    string `sql:"primary_key" json:"app_id" xml:"app_id"`
	TenantID string `sql:"primary_key" json:"tenant_id" xml:"tenant_id"`
	UserID   string `sql:"primary_key" json:"user_id" xml:"user_id"`
	Email    string `json:"email" xml:"email"`
}
