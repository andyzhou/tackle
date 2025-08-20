package video2gif

import (
	"github.com/andyzhou/tackle/app/base"
	wDefine "github.com/andyzhou/tackle/app/define"
	"github.com/andyzhou/tackle/app/view"
	"github.com/andyzhou/tackle/define"
	"github.com/gin-gonic/gin"
	"log"
)

/*
 * post page entry
 */

//face info
type PostEntry struct {
	tpl *base.TplFace
	base.BaseEntry
}

//construct
func NewPostEntry() *PostEntry {
	this := &PostEntry{
		tpl: base.NewTplFace(),
	}
	this.interInit()
	return this
}

//entry
func (f *PostEntry) Entry(
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
	mainTpl, subErr := f.tpl.ParseTpl(wDefine.TplOfVideo2GifPost)
	if subErr != nil {
		log.Printf("page.home, err:%v\n", subErr.Error())
		return "", subErr
	}

	//fill and gen tpl content
	return f.tpl.GetTplContent(mainTpl, tplDataMap)
}

//inter init
func (f *PostEntry) interInit() {
	//init tpl obj
	f.InitTplObj(f.tpl, define.WebReqAppOfVideo2gif)
}
