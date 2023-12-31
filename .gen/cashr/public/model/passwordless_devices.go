//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PasswordlessDevices struct {
	AppID          string  `sql:"primary_key" json:"app_id" xml:"app_id"`
	TenantID       string  `sql:"primary_key" json:"tenant_id" xml:"tenant_id"`
	DeviceIDHash   string  `sql:"primary_key" json:"device_id_hash" xml:"device_id_hash"`
	Email          *string `json:"email" xml:"email"`
	PhoneNumber    *string `json:"phone_number" xml:"phone_number"`
	LinkCodeSalt   string  `json:"link_code_salt" xml:"link_code_salt"`
	FailedAttempts int32   `json:"failed_attempts" xml:"failed_attempts"`
}
