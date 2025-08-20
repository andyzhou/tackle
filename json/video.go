package json

/*
 * video json info
 */

//video meta info
type VideoMetaJson struct {
	Duration int `json:"duration"`
	Width    int `json:"width"`
	Height   int `json:"height"`
	Rotate   int `json:"rotate"`
	BaseJson
}

//video original info
type VideoInfoJson struct {
	Kind        string `json:"kind"`
	Duration    string `json:"duration"`
	RatioWidth  int    `json:"ratioWidth"`
	RatioHeight int    `json:"ratioHeight"`
	CreateTime  int64  `json:"createTime"`
	BaseJson
}

//construct
func NewVideoMetaJson() *VideoMetaJson {
	this := &VideoMetaJson{}
	return this
}

//construct
func NewVideoInfoJson() *VideoInfoJson {
	this := &VideoInfoJson{}
	return this
}
