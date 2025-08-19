package cmd

import (
	"github.com/urfave/cli/v2"
)

/*
 * command config face
 */

//global variable
var (
	RunCmdConf *RunCfg
)

//init run command cfg
func InitRunCmdCfg(c *cli.Context) {
	RunCmdConf = GetRunCfg(c)
}

//command flags
func Flags() []cli.Flag  {
	return []cli.Flag{
		&cli.IntFlag{Name: NameOfWeb, Usage: "web port"},
		&cli.StringFlag{Name: NameOfConf, Usage: "conf path"},
		&cli.StringFlag{Name: NameOfLogPath, Usage: "log path"},
		&cli.StringFlag{Name: NameOfLogPrefix, Usage: "log prefix"},
	}
}

//get run config
func GetRunCfg(c *cli.Context) *RunCfg {
	cfg := &RunCfg{
		Web: c.Int(NameOfWeb),
		ConfPath: c.String(NameOfConf),
		LogPath:  c.String(NameOfLogPath),
		LogPrefix: c.String(NameOfLogPrefix),
	}
	return cfg
}
