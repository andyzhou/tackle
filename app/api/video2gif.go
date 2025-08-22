package api

import (
	"errors"
	"fmt"
	"github.com/andyzhou/tackle/app/base"
	aDefine "github.com/andyzhou/tackle/app/define"
	"github.com/andyzhou/tackle/app/form"
	"github.com/andyzhou/tackle/define"
	"github.com/andyzhou/tackle/file"
	"github.com/gin-gonic/gin"
)

/*
 * video2gif api face
 */

//face info
type Video2Gif struct {
	base.Upload
}

//construct
func NewVideo2Gif() *Video2Gif {
	this := &Video2Gif{}
	return this
}

//entry
//return jsonObj, errCode, error
func (f *Video2Gif) Entry(
	cookieUserId int64,
	ctx *gin.Context) (interface{}, int, error) {
	//get sub action
	subAct := ctx.Param(define.ParaOfSubAct)
	switch subAct {
	case aDefine.SubActOfUpload:
		{
			//upload video
			return f.uploadVideo(cookieUserId, ctx)
			break
		}
	default:
		{
			//default
			return nil, define.CodeInvalidOpt, fmt.Errorf("video2gif, invalid act `%v`", subAct)
		}
	}
	return nil, 0, nil
}

//upload video
//return jsonObj, errCode, error
func (f *Video2Gif) uploadVideo(
	cookieUserId int64,
	ctx *gin.Context) (interface{}, int, error) {
	var (
		reqForm form.Video2GifUploadForm
		ajaxResp interface{}
		errCode int
		err error
	)
	//check
	if cookieUserId <= 0 {
		return nil, define.CodeNoAccess, errors.New("user not login")
	}

	//get form para
	err = ctx.ShouldBind(&reqForm)
	if err != nil {
		//failed
		code := define.CodeInterError
		return nil, code, err
	}

	//get form key para
	fileId := reqForm.FileId
	startTime := reqForm.StartTime
	if fileId == "" || startTime < 0 {
		//failed
		code := define.CodeInvalidParam
		return nil, code, errors.New("invalid parameter")
	}

	//read uploaded file
	fileJson, fileReader, subErr := f.UploadOneFileInfo(fileId, ctx)
	if subErr != nil {
		//failed
		code := define.CodeInterError
		return nil, code, subErr
	}

	//process video2gif
	ajaxResp, errCode, err = file.GetInterFile().
								GetVideo2GifFile().
								UploadOriginFile(
									cookieUserId,
									fileJson,
									fileReader,
									reqForm.StartTime)
	return ajaxResp, errCode, err
}