//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type JwtSigningKeys struct {
	AppID     string `sql:"primary_key"`
	KeyID     string `sql:"primary_key"`
	KeyString string
	Algorithm string
	CreatedAt *int64
}
