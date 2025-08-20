package home

import (
	"github.com/andyzhou/tackle/app/base"
	wDefine "github.com/andyzhou/tackle/app/define"
	"github.com/andyzhou/tackle/define"
	"github.com/andyzhou/tinylib/util"
	"github.com/gin-gonic/gin"
	"log"
)

/*
 * home entry
 */

//face info
type PageEntry struct {
	//sub face
	home *HomeEntry

	//base
	base.BaseEntry
	tpl *base.TplFace
	faces *util.FaceMap
}

//construct
func NewPageEntry() *PageEntry {
	this := &PageEntry{
		home: NewHomeEntry(),

		tpl: base.NewTplFace(),
		faces: util.NewFaceMap(),
	}
	this.interInit()
	return this
}

//entry
//dynamic load sub page
//return pageContent, error
func (f *PageEntry) Entry(
	cookieInfo string,
	playerId int64,
	ctx *gin.Context) (string, error) {
	//get path parameter
	subModule := ctx.Param(define.ParaOfSubModule)
	if subModule == "" {
		subModule = "home"
	}

	//get sub face by name
	subFace := f.faces.GetFace(subModule)
	if subFace == nil {
		//un-support sub face
		//show 404 page
		return f.GetNotFoundAjaxPage(ctx)
	}

	//call dynamic sub face
	//return sub page content
	respValues, _ := f.faces.Call(subModule, "Entry", cookieInfo, playerId, ctx)
	respPage, err := f.AnalyzeCallResp(respValues...)
	if err != nil {
		log.Printf("web.home.entry, err:%v\n", err.Error())
	}
	return respPage, err
}

//inter init
func (f *PageEntry) interInit() {
	//bind sub faces
	f.faces.Bind(wDefine.SubFaceOfHome, f.home)
}