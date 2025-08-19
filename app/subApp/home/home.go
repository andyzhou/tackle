package home

import (
	"github.com/andyzhou/tackle/app/base"
	wDefine "github.com/andyzhou/tackle/app/define"
	"github.com/andyzhou/tackle/app/view"
	"github.com/andyzhou/tackle/define"
	"github.com/gin-gonic/gin"
	"log"
)

/*
 * home page entry
 */

//face info
type HomeEntry struct {
	tpl *base.TplFace
	base.BaseEntry
}

//construct
func NewHomeEntry() *HomeEntry {
	this := &HomeEntry{
		tpl: base.NewTplFace(),
	}
	this.interInit()
	return this
}

//entry
func (f *HomeEntry) Entry(
	subApp,
	cookieInfo string,
	playerId int64,
	ctx *gin.Context) (string, error) {
	var (
		pageView view.BaseView
	)

	//setup page view
	pageView.CookiePlayerInfo = cookieInfo
	pageView.CookiePlayerId = playerId

	//convert view obj into map
	tplDataMap, _ := f.EncodeJsonObj2Map(pageView)

	//parse tpl
	mainTpl, subErr := f.tpl.ParseTpl(wDefine.TplOfPageHome)
	if subErr != nil {
		log.Printf("page.home, err:%v\n", subErr.Error())
		return "", subErr
	}

	//fill and gen tpl content
	return f.tpl.GetTplContent(mainTpl, tplDataMap)
}

//inter init
func (f *HomeEntry) interInit() {
	//init tpl obj
	f.InitTplObj(f.tpl, define.WebReqAppOfHome)
}
