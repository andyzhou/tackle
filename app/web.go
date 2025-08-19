package app

import (
	"github.com/andyzhou/tackle/app/base"
	"github.com/andyzhou/tackle/app/subApp"
	"github.com/gin-gonic/gin"
)

/*
 * web page entry face
 */

//face info
type WebPageEntry struct {
	gin      *gin.Engine //gin reference obj
	cookie   *base.Cookie
	subEntry *subApp.MainEntry
	base.BaseEntry
}

//construct
func NewWebPageEntry(gin *gin.Engine) *WebPageEntry {
	this := &WebPageEntry{
		gin:      gin,
		cookie:   base.NewCookie(),
		subEntry: subApp.NewMainEntry(),
	}
	return this
}

//main entry
func (f *WebPageEntry) Entry(ctx *gin.Context) {
	//check or init player
	playerId, cookieInfo, _ := f.CheckOrInitPlayer(f.cookie, ctx)

	//call sub entry
	f.subEntry.Entry(cookieInfo, playerId, ctx)
}
