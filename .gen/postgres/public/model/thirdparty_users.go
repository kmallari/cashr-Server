//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type ThirdpartyUsers struct {
	AppID            string `sql:"primary_key"`
	ThirdPartyID     string
	ThirdPartyUserID string
	UserID           string `sql:"primary_key"`
	Email            string
	TimeJoined       int64
}