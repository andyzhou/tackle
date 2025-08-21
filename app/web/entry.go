package web

import (
	"fmt"
	"github.com/andyzhou/tackle/app/base"
	wDefine "github.com/andyzhou/tackle/app/define"
	"github.com/andyzhou/tackle/app/view"
	"github.com/andyzhou/tackle/app/web/file"
	"github.com/andyzhou/tackle/app/web/home"
	"github.com/andyzhou/tackle/app/web/video2gif"
	"github.com/andyzhou/tackle/conf"
	"github.com/andyzhou/tackle/define"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

/*
 * web app main entry face
 */

//face info
type MainEntry struct {
	home *home.PageEntry
	video2gif *video2gif.PageEntry
	file *file.InterEntry
	base.BaseEntry
}

//construct
func NewMainEntry() *MainEntry {
	this := &MainEntry{
		home: home.NewPageEntry(),
		video2gif: video2gif.NewPageEntry(),
		file: file.NewInterEntry(),
	}
	return this
}

//entry
//dynamic load sub page
//return pageContent, error
func (f *MainEntry) Entry(
	cookieInfo string,
	userId int64,
	ctx *gin.Context) {
	var (
		dynamicPageContent string
	)
	//get key path para
	subApp := ctx.Param(define.ParaOfSubApp)
	if subApp == "" {
		//get default app home
		mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
		defaultApp := mainConf.DefaultApp
		subApp = defaultApp
	}

	//get dynamic page para
	//if has value means load sub page pass ajax mode
	dynamicPage, _ := f.GetPara(define.ParaOfPage, ctx)
	isDynamicPage, _ := strconv.ParseBool(dynamicPage)

	//call sub face by group
	switch subApp {
	case define.WebReqAppOfVideo2gif:
		{
			//for video2gif
			dynamicPageContent, _ = f.video2gif.Entry(cookieInfo, userId, ctx)
			break
		}
	case define.WebReqAppOfFile:
		{
			//for read file
			//process and return
			f.file.Entry(userId, ctx)
			return
		}
	default:
		{
			//for general page app
			dynamicPageContent, _ = f.home.Entry(cookieInfo, userId, ctx)
		}
	}

	if isDynamicPage {
		//return page content for dynamic call
		ctx.String(http.StatusOK, dynamicPageContent)
	}else{
		//for full page mode
		//get browser request uri
		browserReqUri := ctx.Request.RequestURI
		if strings.Contains(browserReqUri, "?") == false {
			browserReqUri = fmt.Sprintf("%v?", browserReqUri)
		}

		//init page view
		pageView := view.MainPageView{}
		pageView.CookiePlayerInfo = cookieInfo
		pageView.CookiePlayerId = userId
		pageView.BrowserOrgUri = browserReqUri

		//get and set app header tpl
		mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
		pageView.DefaultApp = mainConf.DefaultApp

		//output global main page tpl
		ctx.HTML(http.StatusOK, wDefine.TplOfGlobalMain, pageView)
	}
}