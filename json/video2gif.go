package json

//json info
type Video2GifFileJson struct {
	ShortUrl  string `json:"shortUrl"`
	Md5       string `json:"md5"`
	Snap      string `json:"snap"`
	Gif       string `json:"gif"`
	Tags      string `json:"tags"`
	Likes     int    `json:"likes"`
	Downloads int    `json:"downloads"`
	CreateAt  int64  `json:"createAt"`
	BaseJson
}

//construct
func NewVideo2GifFileJson() *Video2GifFileJson {
	this := &Video2GifFileJson{}
	return this
}