package main

import (
	"github.com/andyzhou/tackle/cmd"
	"github.com/andyzhou/tackle/define"
	tcmd "github.com/andyzhou/tinylib/cmd"
	"github.com/andyzhou/tinylib/util"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

//global variable
var (
	ClientCTX *cli.Context
	TCmd      *tcmd.Cmd
	Logger    *util.LogService
	WebGin    *gin.Engine
	CloseChan chan bool
)

//init
func init() {
	//watch signal
	WatchSignal()

	//init run cmd
	TCmd = tcmd.NewCmd()
}

//cb for daemon quit
func CBForDaemonQuit() {
	if Logger != nil {
		Logger.Close()
		Logger = nil
	}
}

//watch signal
func WatchSignal() {
	//init chan
	CloseChan = make(chan bool, 1)
	//register signal
	s := util.NewSignal()
	s.RegisterShutDownChan(CloseChan, CBForDaemonQuit)
	s.MonSignal()
}

//init logger
func InitLogger(c *cli.Context) {
	//get and set para
	runCfg := cmd.GetRunCfg(c)
	logPath := runCfg.LogPath
	logPrefix := runCfg.LogPrefix
	if logPath == "" {
		logPath = define.DefaultLoggerPath
	}
	if logPrefix == "" {
		logPrefix = define.DefaultLoggerPrefix
	}

	//init logger service
	Logger = util.NewLogService(logPath, logPrefix)
}