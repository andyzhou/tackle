package conf

import (
	"encoding/json"
	"sync"
)

/*
 * main config
 */

//main config info
type mainConfInfo struct {
	AppRoot        string   `json:"appRoot"`
	TplPath        string   `json:"tplPath"`
	HtmlPath       string   `json:"htmlPath"`
	PrivatePath    string   `json:"privatePath"`
	TempPath       string   `json:"tempPath"`
	CommandPath    string   `json:"commandPath"`
	SnapFps        int      `json:"snapFps"`
	SnapWidth      int      `json:"snapWidth"`
	AnimateSeconds int      `json:"animateSeconds"`
	AnimateScale   int      `json:"animateScale"`
	CookieName     string   `json:"cookieName"`
	CookieSecurity string   `json:"cookieSecurity"`
	ReferDomain    []string `json:"referDomain"`
	DefaultApp     string   `json:"defaultApp"`
}
type MainConf struct {
	confInfo *mainConfInfo
	sync.RWMutex
}

//construct
func NewMainConf() *MainConf {
	//self init
	this := &MainConf{
		confInfo: newMainConfInfo(),
	}
	return this
}

func newMainConfInfo() *mainConfInfo {
	this := &mainConfInfo{
		ReferDomain: []string{},
	}
	return this
}

//get config info
func (c *MainConf) GetConfInfo() *mainConfInfo {
	return c.confInfo
}

//analyze config
func (c *MainConf) AnalyzeConf(config interface{}) bool {
	//home check
	configMap, ok := config.(map[string]interface{})
	if !ok {
		return false
	}
	if len(configMap) <= 0 {
		return false
	}

	c.Lock()
	defer c.Unlock()

	//json encode and decode
	confBytes, _ := json.Marshal(configMap)
	json.Unmarshal(confBytes, &c.confInfo)
	return true
}