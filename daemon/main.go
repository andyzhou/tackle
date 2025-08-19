package main

import (
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"sync"

	"github.com/andyzhou/tackle/cmd"
	"github.com/andyzhou/tackle/conf"
	"github.com/andyzhou/tackle/define"
	"github.com/urfave/cli/v2"
)

/*
 * main daemon
 */

//start app daemon
func startApp(c *cli.Context) error {
	//init cmd conf
	cmd.InitRunCmdCfg(c)

	//init logger
	InitLogger(c)

	//get core cmd para
	cmdConf := cmd.GetRunCfg(c)
	confPath := cmdConf.ConfPath

	//init config
	conf.RunAppConfig = conf.NewAppConfig(confPath)

	//init web service
	err := InitWebService(c)

	return err
}

//main entry
func main() {
	var (
		wg sync.WaitGroup
		m any = nil
	)

	//setup run time
	runtime.GOMAXPROCS(runtime.NumCPU())

	//try catch panic
	defer func() {
		if err := recover(); err != m {
			log.Println("daemon.main, panic happened, err:", err)
			log.Println("daemon.main, stack:", string(debug.Stack()))
			os.Exit(1)
		}
	}()

	//get command flags
	flags := cmd.Flags()

	//init app
	app := &cli.App{
		Name: define.AppName,
		Action: func(c *cli.Context) error {
			ClientCTX = c
			return startApp(c)
		},
		Flags: flags,
	}
	wg.Add(1)

	//start app
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("%v run failed, err:%v\n", define.AppName, err.Error())
		return
	}

	//wait
	wg.Wait()

	//all done, clear up
}