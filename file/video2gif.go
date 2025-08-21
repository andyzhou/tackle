package file

import (
	"errors"
	"fmt"
	"github.com/andyzhou/pond"
	"github.com/andyzhou/tackle/conf"
	"github.com/andyzhou/tackle/db"
	"github.com/andyzhou/tackle/define"
	"github.com/andyzhou/tackle/json"
	"github.com/andyzhou/tinylib/util"
	"io"
	"mime/multipart"
	"os"
	"time"
)

/*
 * video2gif file storage
 * - include snap, gif images
 */

//face info
type Video2Gif struct {
	pond *pond.Pond
	video *Video
	shortUrl *util.ShortUrl
	Base
}

//construct
func NewVideo2Gif() *Video2Gif {
	this := &Video2Gif{
		video: NewVideo(),
		shortUrl: util.NewShortUrl(),
	}
	this.interInit()
	return this
}

//quit
func (f *Video2Gif) Quit() {
	if f.pond != nil {
		f.pond.Quit()
		f.pond = nil
	}
}

//upload origin file
//ajaxResp, errCode, error
func (f *Video2Gif) UploadOriginFile(
	userId int64,
	info *json.FileJson,
	file multipart.File,
	startTime int) (interface{}, int, error) {
	var (
		tempVideoFile string
		ajaxResp interface{}
		errCode int
		err error
	)
	//check
	if info == nil || file == nil {
		return nil, define.CodeInvalidParam, errors.New("invalid parameter of `UploadOriginFile`")
	}

	//defer remove origin video file
	defer func() {
		if tempVideoFile != "" {
			os.RemoveAll(tempVideoFile)
		}
	}()

	//save origin video file
	tempVideoFile, err = f.saveOriginVideoFile(info, file)
	if err != nil {
		errCode = define.CodeInterError
		return ajaxResp, errCode, err
	}

	//get video meta info
	videoMetaInfo := f.video.GetMetaInfo(tempVideoFile)
	if videoMetaInfo == nil {
		errCode = define.CodeInterError
		return ajaxResp, errCode, errors.New("can't get video meta info")
	}

	//take snap image
	snapBytes, subErr := f.video.TakeSnapImage(tempVideoFile, videoMetaInfo, startTime)
	if subErr != nil {
		errCode = define.CodeInterError
		return ajaxResp, errCode, subErr
	}

	//task gif image
	endTime := startTime + define.AnimateGifMaxSeconds
	gifBytes, subErrOne := f.video.GenAnimateGif(tempVideoFile, videoMetaInfo, startTime, endTime)
	if subErrOne != nil {
		errCode = define.CodeInterError
		return ajaxResp, errCode, subErrOne
	}

	//save snap and gif into pond storage
	snapShortUrl, err := f.pond.WriteData(snapBytes)
	if err != nil {
		errCode = define.CodeInterError
		return ajaxResp, errCode, err
	}
	gifShortUrl, err := f.pond.WriteData(gifBytes)
	if err != nil {
		errCode = define.CodeInterError
		return ajaxResp, errCode, err
	}

	//gen snap bytes md5
	md5 := f.GenMd5(snapBytes)

	//check db
	video2GifDB := db.GetInterDB().GetVideo2GifDB()
	fileInfo, _ := video2GifDB.GetFileByMd5(userId, md5)
	if fileInfo != nil {
		//already exists, return old data
		errCode = define.CodeSuccess
		return fileInfo.ShortUrl, errCode, nil
	}

	//gen unique gif file obj short url
	uniqueShortUrl, _ := f.shortUrl.Generator(md5)

	//format video2gif json obj
	gifObj := json.NewVideo2GifFileJson()
	gifObj.ShortUrl = uniqueShortUrl
	gifObj.UserId = userId
	gifObj.Md5 = md5
	gifObj.Snap = snapShortUrl
	gifObj.Gif = gifShortUrl
	gifObj.CreateAt = time.Now().Unix()

	//save into local db
	err = video2GifDB.AddFile(gifObj)
	if err != nil {
		errCode = define.CodeDataSaveFailed
		return ajaxResp, errCode, err
	}

	//success
	errCode = define.CodeSuccess
	return uniqueShortUrl, errCode, nil
}

//del file
func (f *Video2Gif) DelFile(
	shortUrl string) error {
	//check
	if shortUrl == "" {
		return errors.New("invalid parameter")
	}
	if f.pond == nil {
		return errors.New("inter pond not init")
	}

	//call base api
	err := f.pond.DelData(shortUrl)
	return err
}

//read file
func (f *Video2Gif) ReadFile(
	shortUrl string) ([]byte, error) {
	//check
	if shortUrl == "" {
		return nil, errors.New("invalid parameter")
	}
	if f.pond == nil {
		return nil, errors.New("inter pond not init")
	}

	//call base api
	byteData, err := f.pond.ReadData(shortUrl)
	return byteData, err
}

//write file
func (f *Video2Gif) WriteFile(
	shortUrl string,
	data []byte) error {
	//check
	if shortUrl == "" || data == nil {
		return errors.New("invalid parameter")
	}
	if f.pond == nil {
		return errors.New("inter pond not init")
	}

	//call base api
	_, err := f.pond.WriteData(data, shortUrl)
	return err
}

//save origin video file
//return tempVideoFile, error
func (f *Video2Gif) saveOriginVideoFile(
	info *json.FileJson,
	file multipart.File) (string, error) {
	//check
	if file == nil {
		return "", errors.New("invalid parameter")
	}

	//get main conf
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	privatePath := mainConf.PrivatePath
	tempPath := mainConf.TempPath

	//format temp file path
	now := time.Now().UnixNano()
	tempVideoFile := fmt.Sprintf("%v/%v/%v:%v", privatePath, tempPath, now, info.Name)

	//save into local file
	dstFile, err := os.Create(tempVideoFile)
	if err != nil {
		return "", err
	}

	//copy file content
	_, err = io.Copy(dstFile, file)
	if err != nil {
		return "", err
	}
	return tempVideoFile, nil
}

//inter init
func (f *Video2Gif) interInit() {
	//get main conf
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	privatePath := mainConf.PrivatePath
	filePath := fmt.Sprintf("%v/%v", privatePath, define.FilePathOfVideo2Gif)

	//init pond storage
	pond, err := f.InitPond(filePath)
	if err != nil {
		panic(any(err))
	}
	f.pond = pond
}