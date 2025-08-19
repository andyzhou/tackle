package view

import "github.com/andyzhou/tinylib/util"

/*
 * base page view
 */

type BaseView struct {
	CookiePlayerInfo string `json:"CookiePlayerInfo"`
	CookiePlayerId   int64  `json:"CookiePlayerId"`
	BrowserOrgUri    string `json:"BrowserOrgUri"`
	NotifyAddr       string `json:"NotifyAddr"`
	util.BaseJson
}
