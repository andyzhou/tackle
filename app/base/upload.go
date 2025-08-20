package base

import (
	"errors"
	"github.com/andyzhou/tackle/json"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"strings"
)

/*
 * upload file face
 */

//face info
type Upload struct {
}

//upload one file
//return fileKind, contentType, size, *file, error
func (f *Upload) UploadOneFileInfo(
	fileIdPara string,
	ctx *gin.Context) (*json.FileJson, multipart.File, error) {
	//check
	if fileIdPara == "" {
		return nil, nil, errors.New("invalid parameter")
	}

	//get file data
	fileInfo, err := ctx.FormFile(fileIdPara)
	if err != nil || fileInfo == nil {
		return nil, nil, err
	}

	//try open file
	file, subErr := fileInfo.Open()
	if subErr != nil || file == nil {
		return nil, nil, subErr
	}

	//format file info obj
	fileInfoObj := json.NewFileJson()
	fileInfoObj.Name = fileInfo.Filename

	//get file type
	fileTypeSlice, _ := fileInfo.Header["Content-Type"]
	if len(fileTypeSlice) > 0 {
		fileInfoObj.ContentType = fileTypeSlice[0]
	}
	//check file type
	kindSlice := strings.Split(fileInfoObj.ContentType, "/")
	fileKind := kindSlice[0]
	fileInfoObj.FileKind = fileKind
	fileInfoObj.Size = fileInfo.Size
	return fileInfoObj, file, nil
}