package video2gif

import (
	"github.com/andyzhou/tackle/app/base"
	wDefine "github.com/andyzhou/tackle/app/define"
	"github.com/andyzhou/tackle/app/view"
	"github.com/andyzhou/tackle/lib/db"
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
	cookieInfo string,
	userId int64,
	ctx *gin.Context) (string, error) {
	var (
		pageView view.Video2GifHomePageView
	)

	//setup page view
	pageView.CookiePlayerInfo = cookieInfo
	pageView.CookiePlayerId = userId
	pageView.NoData = true

	if userId > 0 {
		//check db record
		video2gifDB := db.GetInterDB().GetVideo2GifDB()
		records, _ := video2gifDB.GetFiles(userId, 1, 1)
		if len(records) > 0 {
			pageView.NoData = false
		}
	}

	//convert view obj into map
	tplDataMap, _ := f.EncodeJsonObj2Map(pageView)

	//parse tpl
	mainTpl, subErr := f.tpl.ParseTpl(wDefine.TplOfVideo2GifHome)
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
	f.InitTplObj(f.tpl, define.WebReqAppOfVideo2gif)
}
