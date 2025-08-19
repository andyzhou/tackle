package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/andyzhou/tackle/app"
	"github.com/andyzhou/tackle/cmd"
	"github.com/andyzhou/tackle/conf"
	"github.com/andyzhou/tackle/define"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

/*
 * service web service
 */

//init web service
func InitWebService(c *cli.Context) error {
	var (
		m any = nil
	)
	//catch panic
	defer func() {
		if err := recover(); err != m {
			log.Printf("InitWebService panic, err:%v\n", err)
		}
	}()

	//get service port
	cmdCfg := cmd.GetRunCfg(c)
	webPort := cmdCfg.Web
	if webPort <= 0 {
		return errors.New("web port not setup")
	}

	//init gin engine and web entry
	gin.SetMode(gin.ReleaseMode)
	WebGin = gin.Default()
	WebGin.Use(gin.Recovery())

	//get root path from conf
	mainConf := conf.RunAppConfig.GetMainConf().GetConfInfo()
	appRoot := mainConf.AppRoot

	//setup tpl and web path
	webTplPath := fmt.Sprintf("%v%v%v/%v", appRoot, define.WebSubPath, mainConf.TplPath, define.WebTplPattern)
	webPath := fmt.Sprintf("%v%v/%v", appRoot, define.WebSubPath, mainConf.HtmlPath)

	//init templates
	WebGin.LoadHTMLGlob(webTplPath)

	//init static path
	WebGin.Static(define.UriOfHtml, webPath)

	//register request entry
	//like:  /[app]/[module]/[dataId]
	webSubGroupUrl := fmt.Sprintf("/:%v",
		define.ParaOfSubApp,
	)
	webSubModuleUrl := fmt.Sprintf("/:%v/:%v",
		define.ParaOfSubApp,
		define.ParaOfSubModule,
	)
	webSubDataIdUrl := fmt.Sprintf("/:%v/:%v/:%v",
		define.ParaOfSubApp,
		define.ParaOfSubModule,
		define.ParaOfSubDataId,
	)

	//init web entry
	//for web dynamic page
	webPage := app.NewWebPageEntry(WebGin)
	WebGin.Any(define.UriOfRoot, webPage.Entry)
	WebGin.Any(webSubGroupUrl, webPage.Entry)
	WebGin.Any(webSubModuleUrl, webPage.Entry)
	WebGin.Any(webSubDataIdUrl, webPage.Entry)

	//start web service
	fmt.Printf("start web service http://localhost:%v\n", webPort)
	go WebGin.Run(fmt.Sprintf(":%v", webPort))
	return nil
}