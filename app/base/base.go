package base

import (
	"bytes"
	genJson "encoding/json"
	"errors"
	"fmt"
	wDefine "github.com/andyzhou/tackle/app/define"
	"github.com/andyzhou/tackle/conf"
	"github.com/andyzhou/tackle/define"
	"github.com/andyzhou/tackle/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
)

//face info
type BaseEntry struct {
}

//init tpl obj
func (f *BaseEntry) InitTplObj(
	tpl *TplFace,
	subTplDirs ...string) {
	var (
		subTplDir string
	)
	if len(subTplDirs) > 0 {
		subTplDir = subTplDirs[0]
	}

	//get tpl path
	tplPath := f.GetTplFullPath()
	subTplPath := tplPath
	if subTplDir != "" {
		subTplPath = fmt.Sprintf("%v/%v", subTplPath, subTplDir)
	}

	//setup tpl
	tpl.SetTplPath(subTplPath)
}

//get tpl full path
func (f *BaseEntry) GetTplFullPath() string {
	//get root path from conf
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	webRoot := mainConf.AppRoot
	return fmt.Sprintf("%v/%v/%v", webRoot, define.WebSubPath, mainConf.TplPath)
}

//read request json body
func (f *BaseEntry) ReadReqBody(
	outObj interface{},
	ctx *gin.Context) error {
	//read request body
	jsonByte, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	if jsonByte == nil || len(jsonByte) <= 0 {
		return errors.New("can't read body data")
	}
	//decode json data
	err = genJson.Unmarshal(jsonByte, outObj)
	return err
}

//get para
func (f *BaseEntry) GetPara(
	name string,
	ctx *gin.Context) (string, error) {
	//get para from query
	val := ctx.Query(name)
	if val != "" {
		return val, nil
	}
	//get para from post
	val = ctx.PostForm(name)
	return val, nil
}

//ajax json response
func (f *BaseEntry) AjaxResp(
		value interface{},
		errCode int,
		errMsg ...string,
	) interface{} {
	respJson := json.NewResponseJson()
	respJson.Val = value
	respJson.ErrCode = errCode
	if errMsg != nil && len(errMsg) > 0 {
		respJson.ErrMsg = errMsg[0]
	}
	return respJson
}

//check or init player
//return playerId, cookieOrg, error
func (f *BaseEntry) CheckOrInitPlayer(
	cookie *Cookie,
	ctx *gin.Context) (int64, string, error) {
	var (
		playerId int64
		err error
	)
	//get player cookie
	cookieOrg, cookieInfo := cookie.GetCookieOrg(ctx)
	if cookieInfo != "" {
		playerId, _ = strconv.ParseInt(cookieInfo, 10, 64)
		return playerId, cookieOrg, nil
	}

	////get relate data face
	//idData := data.GetData().GetRedisData().GetId()
	//playerData := data.GetData().GetRedisData().GetPlayer()
	//
	////init new player
	//playerId, _ := idData.GenPlayerId()
	//
	////init player obj
	//playerObj := json.NewPlayerJson()
	//playerObj.Id = playerId
	//playerObj.CreateAt = time.Now().Unix()
	//
	////save into redis
	//err := playerData.SetPlayer(playerObj)
	//if err != nil {
	//	return 0, cookieOrg, err
	//}

	//set player id into cookie
	cookieOrg, err = cookie.SetCookie(fmt.Sprintf("%v", playerId), ctx)
	return playerId, cookieOrg, err
}

//convert json to obj
//force convert big integer to json.Number
func (f *BaseEntry) EncodeJsonObj2Map(
	jsonObj interface{}) (map[string]interface{}, error) {
	//convert to hash map obj
	jsonBytes, err := genJson.Marshal(jsonObj)
	if err != nil {
		return nil, err
	}
	resultMap := make(map[string]interface{})

	//decode map obj
	decoder := genJson.NewDecoder(bytes.NewReader(jsonBytes))
	decoder.UseNumber()
	err = decoder.Decode(&resultMap)
	return resultMap, err
}

//analyze response page values
func (f *BaseEntry) AnalyzeCallResp(
		values ...reflect.Value,
	) (string, error) {
	var (
		respPage string
		err error
	)
	//check
	if values == nil || len(values) < wDefine.WebResponseValLen {
		return respPage, errors.New("invalid parameter")
	}

	//get response json obj value
	//first element of values
	respVal := values[0].Interface()
	respPage, _ = respVal.(string)

	//check and convert error interface
	//third element of values
	if errVal, ok := values[1].Interface().(error); ok {
		err = errVal
	}
	return respPage, err
}

//get not found ajax page
func (f *BaseEntry) GetNotFoundAjaxPage(
		ctx *gin.Context,
	) (string, error) {
	//global tpl
	tpl := NewTplFace()

	////setup tpl data map
	tplDataMap := make(map[string]interface{})

	//load and parse dynamic tpl file
	mainTpl, err := tpl.ParseTpl(wDefine.TplOfNotFound)
	if err != nil {
		log.Printf("web.base.notFoundPage, err:%v\n", err.Error())
		return "", err
	}

	//fill and gen tpl content
	return tpl.GetTplContent(mainTpl, tplDataMap)
}