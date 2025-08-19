package json

/*
 * web api response json
 */

//json info
type ResponseJson struct {
	Val     interface{} `json:"val"`
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	BaseJson
}

type AjaxRespJson struct {
	Result interface{} `json:"result"`
}

//construct
func NewResponseJson() *ResponseJson {
	this := &ResponseJson{}
	return this
}

func NewAjaxRespJson() *AjaxRespJson {
	this := &AjaxRespJson{}
	return this
}