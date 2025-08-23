package conf

import (
	"fmt"
	"github.com/andyzhou/tackle/define"
	"github.com/andyzhou/tinylib/config"
	"log"
	"sync"
)

//internal macro variables
const (
	//others
	CheckConfRate = 120 //xxx seconds
)

//declare as global variable
var (
	RunAppConfig *AppConfig
)

//sub config sections info
type SubConfSections struct {
	mainConf      *MainConf
	video2gifConf *Video2gifConf
}

//app config info
type AppConfig struct {
	//base
	subConfOfMain  *config.SubConfig
	subConfOfVideo2gif *config.SubConfig

	confDir string
	subConf *SubConfSections
	//cb for updated func map
	cbForUpdatedMap map[string]func() bool //tag -> func
	sync.RWMutex
}

//construct
func NewAppConfig(confDir string) *AppConfig {
	this := &AppConfig{
		confDir: confDir,
	}

	//init sub config section instance
	this.subConf = &SubConfSections{
		mainConf:NewMainConf(),
		video2gifConf: NewVideo2gifConf(),
	}

	//get relate config file path
	mainConfigFile := fmt.Sprintf("%s/%s", confDir, define.ConfOfMain)
	video2gifConfigFile := fmt.Sprintf("%s/%s", confDir, define.ConfOfVideo2Gif)

	//init sub config files
	this.subConfOfMain = config.NewSubConfig(mainConfigFile, this.CBForMain, CheckConfRate)
	this.subConfOfVideo2gif = config.NewSubConfig(video2gifConfigFile, this.CBForVideo2gif, CheckConfRate)

	return this
}

//add cb for config updated
func (c *AppConfig) AddCBForUpdated(tag string, cb func() bool) bool {
	if tag == "" || cb == nil {
		return false
	}
	c.Lock()
	defer c.Unlock()
	c.cbForUpdatedMap[tag] = cb
	return true
}

//quit
func (c *AppConfig) Quit() {
	c.subConfOfMain.Quit()
}

//get config root path
func (c *AppConfig) GetConfPath() string {
	return c.confDir
}

//get main config
func (c *AppConfig) GetMainConf() *MainConf {
	return c.subConf.mainConf
}

//get video2gif config
func (c *AppConfig) GetVideo2gifConf() *Video2gifConf {
	return c.subConf.video2gifConf
}

//call back for main
func (c *AppConfig) CBForMain(allConfMap map[string]interface{}) bool {
	if allConfMap == nil || len(allConfMap) <= 0 {
		log.Println("AppConfig::CBForMain, no any config info")
		return false
	}
	c.subConf.mainConf.AnalyzeConf(allConfMap)
	return true
}

//call back for video2gif
func (c *AppConfig) CBForVideo2gif(allConfMap map[string]interface{}) bool {
	if allConfMap == nil || len(allConfMap) <= 0 {
		log.Println("AppConfig::CBForVideo2gif, no any config info")
		return false
	}
	c.subConf.video2gifConf.AnalyzeConf(allConfMap)
	return true
}

////////////////
//private func
////////////////

//get cb func for config updated
func (c *AppConfig) getCBFunc(tag string) func() bool {
	//check
	if tag == "" || c.cbForUpdatedMap == nil {
		return nil
	}
	//get from map
	c.Lock()
	defer c.Unlock()
	v, ok := c.cbForUpdatedMap[tag]
	if ok && v != nil {
		return v
	}
	return nil
}