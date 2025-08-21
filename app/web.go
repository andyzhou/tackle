package app

import (
	"fmt"
	"github.com/andyzhou/tackle/app/base"
	"github.com/andyzhou/tackle/app/web"
	"github.com/andyzhou/tackle/define"
	"github.com/gin-gonic/gin"
)

/*
 * web page entry face
 */

//face info
type WebPageEntry struct {
	gin      *gin.Engine //gin reference obj
	cookie   *base.Cookie
	webEntry *web.MainEntry
	base.BaseEntry
}

//construct
func NewWebPageEntry(gin *gin.Engine) *WebPageEntry {
	this := &WebPageEntry{
		gin:      gin,
		cookie:   base.NewCookie(),
		webEntry: web.NewMainEntry(),
	}
	this.interInit()
	return this
}

//main entry
func (f *WebPageEntry) Entry(ctx *gin.Context) {
	//check or init user
	userId, cookieInfo, _ := f.CheckOrInitUser(f.cookie, ctx)

	//call sub entry
	f.webEntry.Entry(cookieInfo, userId, ctx)
}

//inter init
func (f *WebPageEntry) interInit() {
	//setup web app uri
	//like:  /[subApp]/[module]/[dataId]
	webSubGroupUrl := fmt.Sprintf("/:%v",
		define.ParaOfSubApp,
	)
	webSubModuleUrl := fmt.Sprintf("/:%v/:%v",
		define.ParaOfSubApp,
		define.ParaOfSubModule,
	)
	webSubDataIdUrl := fmt.Sprintf("/:%v/:%v/:%v",
		define.ParaOfSubApp,
		define.ParaOfSubModule,
		define.ParaOfSubDataId,
	)

	//register request entry
	f.gin.Any(define.UriOfRoot, f.Entry)
	f.gin.Any(webSubGroupUrl, f.Entry)
	f.gin.Any(webSubModuleUrl, f.Entry)
	f.gin.Any(webSubDataIdUrl, f.Entry)
}
