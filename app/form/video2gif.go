package form

type Video2GifUploadForm struct {
	FileId    string `form:"fileId"`
	StartTime int    `form:"startTime"`
	Tag       string `form:"tag"`
}

type Video2GifDeleteForm struct {
	Uri string `form:"uri"`
}
