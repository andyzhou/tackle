package json

//file json
type FileJson struct {
	ShortUrl    string `json:"shortUrl"` //origin file url
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Size        int64  `json:"size"`
	Width       int    `json:"width"`
	Height		int    `json:"height"`
	ContentType string `json:"contentType"`
	FileKind    string `json:"fileKind"`
	Duration    int    `json:"duration"` //for video kind
	Data        []byte `json:"data"`
	BaseJson
}

//construct
func NewFileJson() *FileJson {
	this := &FileJson{
		Data: []byte{},
	}
	return this
}