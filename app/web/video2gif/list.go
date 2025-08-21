package video2gif

import (
	"github.com/andyzhou/tackle/app/base"
	wDefine "github.com/andyzhou/tackle/app/define"
	"github.com/andyzhou/tackle/app/view"
	"github.com/andyzhou/tackle/db"
	"github.com/andyzhou/tackle/define"
	"github.com/andyzhou/tackle/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

/*
 * list page entry
 */

//face info
type ListEntry struct {
	tpl *base.TplFace
	base.BaseEntry
}

//construct
func NewListEntry() *ListEntry {
	this := &ListEntry{
		tpl: base.NewTplFace(),
	}
	this.interInit()
	return this
}

//entry
func (f *ListEntry) Entry(
	cookieInfo string,
	userId int64,
	ctx *gin.Context) (string, error) {
	var (
		pageView view.Video2GifListPageView
	)

	//get key para
	page, _ := f.GetPara(define.ParaOfPageNo, ctx)
	pageInt, _ := strconv.Atoi(page)

	//setup page view
	pageView.CookiePlayerInfo = cookieInfo
	pageView.CookiePlayerId = userId

	//get batch file info
	video2gifDB := db.GetInterDB().GetVideo2GifDB()
	filesJson, _ := video2gifDB.GetFiles(userId, pageInt, define.RecSmallPage)

	pageView.FilesInfo = make([]*json.Video2GifFileJson, 0)
	if filesJson != nil {
		pageView.FilesInfo = filesJson
		if len(filesJson) > 0 {
			pageView.ListMoreSwitcher = true
		}
	}

	//convert view obj into map
	tplDataMap, _ := f.EncodeJsonObj2Map(pageView)

	//parse tpl
	mainTpl, subErr := f.tpl.ParseTpl(wDefine.TplOfVideo2GifList)
	if subErr != nil {
		log.Printf("page.list, err:%v\n", subErr.Error())
		return "", subErr
	}

	//fill and gen tpl content
	return f.tpl.GetTplContent(mainTpl, tplDataMap)
}

//inter init
func (f *ListEntry) interInit() {
	//init tpl obj
	f.InitTplObj(f.tpl, define.WebReqAppOfVideo2gif)
}
