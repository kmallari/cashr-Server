//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type UserLastActive struct {
	AppID          string `sql:"primary_key" json:"app_id" xml:"app_id"`
	UserID         string `sql:"primary_key" json:"user_id" xml:"user_id"`
	LastActiveTime *int64 `json:"last_active_time" xml:"last_active_time"`
}
