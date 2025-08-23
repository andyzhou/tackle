package file

import (
	"fmt"
	"github.com/andyzhou/tackle/define"
	"github.com/andyzhou/tackle/lib/file"
	"github.com/andyzhou/tinylib/web"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
 * video2gif file read entry
 */


//face info
type Video2GifEntry struct {
	web.Base
}

//construct
func NewVideo2GifEntry() *Video2GifEntry {
	this := &Video2GifEntry{
	}
	return this
}

//entry
func (f *Video2GifEntry) Entry(
	cookieUserId int64,
	ctx *gin.Context) {
	//check
	if cookieUserId <= 0 {
		//not login
		return
	}

	//get path parameter
	dataId := ctx.Param(define.ParaOfSubDataId)

	//get file short uri
	shortUri := f.GetPara(define.ParaOfShortUri, ctx)
	download := f.GetPara(define.ParaOfDownload, ctx)

	//check key para
	if dataId == "" || shortUri == "" {
		ctx.Writer.Write([]byte("invalid parameter"))
		return
	}

	//verify file data for download
	isDownload, _ := strconv.ParseBool(download)

	//get response writer
	w := ctx.Writer

	//read file data by short uri
	video2gifFile := file.GetInterFile().GetVideo2GifFile()
	fileData, err := video2gifFile.ReadFile(shortUri)
	if err != nil || fileData == nil {
		ctx.Writer.Write([]byte(err.Error()))
		return
	}

	//if download opt
	if isDownload {
		//setup file name
		filetype := http.DetectContentType(fileData)
		fileName := fmt.Sprintf("%v", time.Now().Unix())

		//format file extend name
		if filetype != "" {
			fileTypeArr := strings.Split(filetype, "/")
			if len(fileTypeArr) > 1 {
				fileName = fmt.Sprintf("%v.%v", fileName, fileTypeArr[1])
			}
		}

		//setup download header
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
		w.Header().Set("Content-Length", string(len(fileData)))
	}

	//output file data
	w.Write(fileData)
}