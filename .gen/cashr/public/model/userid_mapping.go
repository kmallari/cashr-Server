//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type UseridMapping struct {
	AppID              string  `sql:"primary_key" json:"app_id" xml:"app_id"`
	SupertokensUserID  string  `sql:"primary_key" json:"supertokens_user_id" xml:"supertokens_user_id"`
	ExternalUserID     string  `sql:"primary_key" json:"external_user_id" xml:"external_user_id"`
	ExternalUserIDInfo *string `json:"external_user_id_info" xml:"external_user_id_info"`
}