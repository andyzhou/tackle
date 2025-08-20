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

	//init web page entry
	app.NewWebPageEntry(WebGin)

	//init web api entry
	app.NewApiEntry(WebGin)

	//start web service
	fmt.Printf("start web service http://localhost:%v\n", webPort)
	go WebGin.Run(fmt.Sprintf(":%v", webPort))
	return nil
}