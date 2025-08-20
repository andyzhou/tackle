package form

type Video2GifUploadForm struct {
	FileId    string `form:"fileId"`
	StartTime int    `form:"startTime"`
	Tag       string `form:"tag"`
}
