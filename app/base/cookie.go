package base

import (
	"fmt"
	"github.com/andyzhou/tackle/conf"
	"github.com/andyzhou/tinylib/crypt"
	"github.com/andyzhou/tinylib/web"
	"github.com/gin-gonic/gin"
)

//cookie face
type Cookie struct {
	crypt *crypt.Crypt
	web.Cookie
}

//construct
func NewCookie() *Cookie {
	this := &Cookie{
		crypt: crypt.NewCrypt(),
	}
	return this
}

//del cookie
func (f *Cookie) DelCookie(
	ctx *gin.Context) error {
	//get cookie name
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	cookieNameVal := mainConf.CookieName

	//del cookie
	cookieFace := web.GetWeb().GetCookie()
	err := cookieFace.DelCookie(cookieNameVal, "", ctx)
	return err
}

//get cookie origin data
//return origin, decrypted
func (f *Cookie) GetCookieOrg(
	ctx *gin.Context) (string, string) {
	//get cookie name
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	cookieNameVal := mainConf.CookieName

	//get cookie
	cookieVal, _ := f.GetCookie(cookieNameVal, ctx)
	if cookieVal == "" {
		return cookieVal, ""
	}

	//try decode crypt cookie
	decryptCookie, _ := f.getCrypt().GetSimple().Decrypt(cookieVal)
	return cookieVal, decryptCookie
}

//set cookie
func (f *Cookie) SetCookie(
	cookieVal interface{},
	ctx *gin.Context) (string, error) {
	cookieExpire := 0 //forever

	//encode cookie
	cookieValStr := fmt.Sprintf("%v", cookieVal)
	cryptCookie, err := f.getCrypt().GetSimple().Encrypt(cookieValStr)
	if err != nil {
		return "", err
	}

	//get cookie name
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	cookieNameVal := mainConf.CookieName

	//set cookie
	cookieFace := web.GetWeb().GetCookie()
	err = cookieFace.SetCookie(cookieNameVal, cryptCookie, cookieExpire, "", ctx)
	return cryptCookie, err
}

//check or init cookie crypt
func (f *Cookie) getCrypt() *crypt.Crypt {
	if f.crypt != nil {
		return f.crypt
	}

	//get cookie security key
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	securityKey := mainConf.CookieSecurity

	//init crypt
	f.crypt = crypt.NewCrypt()
	if securityKey != "" {
		f.crypt.GetSimple().SetKey(securityKey)
	}
	return f.crypt
}