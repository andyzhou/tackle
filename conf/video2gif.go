package conf

import (
	"encoding/json"
	"sync"
)

/*
 * video2gif config
 */

type video2gifConfInfo struct {
	SnapFps        int      `json:"snapFps"`
	SnapWidth      int      `json:"snapWidth"`
	AnimateSeconds int      `json:"animateSeconds"`
	AnimateScale   int      `json:"animateScale"`
}

//config info
type Video2gifConf struct {
	confInfo *video2gifConfInfo
	sync.RWMutex
}

//construct
func NewVideo2gifConf() *Video2gifConf {
	//self init
	this := &Video2gifConf{
		confInfo: newVideo2gifConfInfo(),
	}
	return this
}

func newVideo2gifConfInfo() *video2gifConfInfo {
	this := &video2gifConfInfo{}
	return this
}

//get config info
func (c *Video2gifConf) GetConfInfo() *video2gifConfInfo {
	return c.confInfo
}

//analyze config
func (c *Video2gifConf) AnalyzeConf(config interface{}) bool {
	//home check
	configMap, ok := config.(map[string]interface{})
	if !ok {
		return false
	}
	if len(configMap) <= 0 {
		return false
	}

	//json encode and decode
	c.Lock()
	defer c.Unlock()
	confBytes, _ := json.Marshal(configMap)
	subConfInfo := newVideo2gifConfInfo()
	json.Unmarshal(confBytes, &subConfInfo)
	c.confInfo = subConfInfo
	return true
}
