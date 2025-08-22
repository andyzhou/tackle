package db

import (
	genJson "encoding/json"
	"errors"
	"fmt"
	"github.com/andyzhou/tackle/db/base"
	"github.com/andyzhou/tackle/define"
	"github.com/andyzhou/tackle/json"
	"time"
)

/*
 * video2gif db opt face
 */

//face info
type Video2Gif struct {
	db *base.SqlLite
	base.Base
}

//construct
func NewVideo2Gif() *Video2Gif {
	this := &Video2Gif{}
	this.interInit()
	return this
}

//quit
func (f *Video2Gif) Quit() {
	if f.db != nil {
		f.db.Close()
		f.db = nil
	}
}

////////////////////////
//////api for users/////
////////////////////////

//add new user
//return newUserId, error
func (f *Video2Gif) AddNewUser(ip string) (int64, error) {
	//format sql
	sql := fmt.Sprintf("INSERT INTO %v(ip, createAt) " +
		"VALUES(?, ?)", define.TabNameOfVideo2GifUsers)
	now := time.Now().Unix()
	values := []interface{}{
		ip,
		now,
	}

	//save into db
	newUserId, _, err := f.db.Execute(sql, values)
	return newUserId, err
}

////////////////////////
//////api for files/////
////////////////////////
//get batch files
func (f *Video2Gif) GetFiles(
	userId int64,
	page, pageSize int,
	sortFields ...string) ([]*json.Video2GifFileJson, error) {
	var (
		sortField string
	)
	//check
	if userId <= 0 {
		return nil, errors.New("invalid user id")
	}
	if len(sortFields) > 0 {
		sortField = sortFields[0]
	}
	if page <= 0 {
		page = define.DefaultPage
	}
	if pageSize <= 0 {
		pageSize = define.RecPerPage
	}

	//setup key value
	offset := (page - 1) * pageSize
	if sortField == "" {
		sortField = define.TabFieldOfScore
	}

	//format sql
	sql := fmt.Sprintf("SELECT * FROM %s WHERE userId= %v ORDER BY %s desc LIMIT ?, ?",
		define.TabNameOfVideo2GifFiles, userId, sortField)
	values := make([]interface{}, 0)
	values = append(
		values,
		offset,
		pageSize,
	)

	//get from db
	records, err := f.db.Query(sql, values)
	if err != nil || len(records) <= 0 {
		return nil, err
	}

	//format result
	result := make([]*json.Video2GifFileJson, 0)
	for _, record := range records {
		if record == nil {
			continue
		}
		fileObj := f.analyzeOneFileInfo(record)
		if fileObj == nil {
			continue
		}
		result = append(result, fileObj)
	}
	return result, nil
}

//get file by short url
func (f *Video2Gif) GetFile(
	userId int64,
	shortUrl string) (*json.Video2GifFileJson, error) {
	//check
	if userId <= 0 || shortUrl == "" {
		return nil, errors.New("invalid parameter")
	}

	//format sql
	sql := fmt.Sprintf("SELECT * FROM %s WHERE userId = %v AND shortUrl = ? LIMIT 1",
		define.TabNameOfVideo2GifFiles, userId)
	values := []interface{}{
		shortUrl,
	}

	//get from db
	records, err := f.db.Query(sql, values)
	if err != nil || len(records) <= 0 {
		return nil, err
	}

	//process single record
	fileObj := f.analyzeOneFileInfo(records[0])
	return fileObj, nil
}

//get file by md5
func (f *Video2Gif) GetFileByMd5(
	userId int64,
	md5 string) (*json.Video2GifFileJson, error) {
	//check
	if md5 == "" {
		return nil, errors.New("invalid parameter")
	}

	//format sql
	sql := fmt.Sprintf("SELECT * FROM %s WHERE userId = %v AND md5 = ? LIMIT 1",
		define.TabNameOfVideo2GifFiles, userId)
	values := []interface{}{
		md5,
	}

	//get from db
	records, err := f.db.Query(sql, values)
	if err != nil || len(records) <= 0 {
		return nil, err
	}

	//process single record
	fileObj := f.analyzeOneFileInfo(records[0])
	return fileObj, nil
}

//delete file by short url
func (f *Video2Gif) DelFile(
	userId int64,
	shortUrl string) error {
	//check
	if userId <= 0 || shortUrl == "" {
		return errors.New("invalid parameter")
	}

	//format sql
	sql := fmt.Sprintf("DELETE FROM %v WHERE userId = ? AND shortUrl = ?", define.TabNameOfVideo2GifFiles)
	values := []interface{}{
		userId,
		shortUrl,
	}

	//remove from db
	_, _, err := f.db.Execute(sql, values)
	return err
}

//add new file
func (f *Video2Gif) AddFile(
	obj *json.Video2GifFileJson) error {
	//check
	if obj == nil || obj.ShortUrl == "" {
		return errors.New("invalid parameter")
	}

	//format sql
	sql := fmt.Sprintf("INSERT INTO %v(shortUrl, userId, md5, snap, gif, tags, createAt) " +
		"VALUES(?, ?, ?, ?, ?, ?, ?)", define.TabNameOfVideo2GifFiles)
	now := time.Now().Unix()
	values := []interface{}{
		obj.ShortUrl,
		obj.UserId,
		obj.Md5,
		obj.Snap,
		obj.Gif,
		obj.Tags,
		now,
	}

	//save into db
	_, _, err := f.db.Execute(sql, values)
	return err
}

//analyze one file info
func (f *Video2Gif) analyzeOneFileInfo(
	record map[string]interface{}) *json.Video2GifFileJson {
	if record == nil {
		return nil
	}
	jsonBytes, err := genJson.Marshal(record)
	if err != nil || jsonBytes == nil {
		return nil
	}
	fileObj := json.NewVideo2GifFileJson()
	err = genJson.Unmarshal(jsonBytes, &fileObj)
	if err != nil {
		return nil
	}
	return fileObj
}

//inter init
func (f *Video2Gif) interInit() {
	//open sqlite db
	db, err := f.OpenDB(define.SqliteFileOfVideo2Gif)
	if err != nil {
		panic(any(err))
	}
	f.db = db
}
