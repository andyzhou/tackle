package file

import "sync"

/*
 * inter file face
 * - base on `pond` storage
 */

//global variable
var (
	_inter *InterFile
	_interOnce sync.Once
)

//face info
type InterFile struct {
	video2Gif *Video2Gif
}

//get single instance
func GetInterFile() *InterFile {
	_interOnce.Do(func() {
		_inter = NewInterFile()
	})
	return _inter
}

//construct
func NewInterFile() *InterFile {
	this := &InterFile{
		video2Gif: NewVideo2Gif(),
	}
	return this
}

//get sub face
func (f *InterFile) GetVideo2GifFile() *Video2Gif {
	return f.video2Gif
}