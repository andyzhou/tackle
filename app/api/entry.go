package api

import (
	"fmt"
	"github.com/andyzhou/tackle/app/base"
	aDefine "github.com/andyzhou/tackle/app/define"
	"github.com/andyzhou/tackle/define"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * api entry
 * - ajax api for web, client, etc.
 * - all return json format data
 */
type InterApiEntry struct {
	//sub api
	video2Gif *Video2Gif

	//base
	base.BaseEntry
}

//construct
func NewInterApiEntry() *InterApiEntry {
	//self init
	this := &InterApiEntry{
		video2Gif: NewVideo2Gif(),
	}
	return this
}

//entry
func (f *InterApiEntry) Entry(ctx *gin.Context) {
	var (
		ajaxResp interface{}
		apiResp interface{}
		errCode int
		err error
	)
	//get core path parameter
	subApp := ctx.Param(define.ParaOfSubApp) //p1

	//refer domain check
	bRet := f.CheckReferDomain(ctx)
	if !bRet {
		//can't access domain
		resp := f.AjaxResp(nil, define.CodeNoAccess, "domain not allow access")
		ctx.JSON(http.StatusOK, resp)
		return
	}

	////auth check
	//_, err = f.CheckHeaderAuth(ctx)
	//if err != nil {
	//	//auth failed
	//	msg := fmt.Sprintf("web auth failed, err:%v", err.Error())
	//	resp := f.AjaxResp(nil, define.CodeInvalidAppAndToken, msg)
	//	ctx.JSON(http.StatusOK, resp)
	//	return
	//}

	//call sub face by sub app
	switch subApp {
	case aDefine.SubApiOfVideo2Gif:
		{
			//for video2gif
			apiResp, errCode, err = f.video2Gif.Entry(ctx)
			break
		}
	default:
		{
			//default sub app
			errCode = define.CodeInvalidApi
			err = fmt.Errorf("invalid sub app `%v`", subApp)
			break
		}
	}

	//output ajax result
	if err != nil {
		//failed
		ajaxResp = f.AjaxResp(nil, errCode, err.Error())
	}else{
		//succeed
		ajaxResp = f.AjaxResp(apiResp, errCode)
	}
	ctx.JSON(http.StatusOK, ajaxResp)
}
