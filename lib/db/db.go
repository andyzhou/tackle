package db

import "sync"

/*
 * inter db face
 */

//global instance
var (
	_inter *InterDB
	_interOnce sync.Once
)

//face info
type InterDB struct {
	video2Gif *Video2Gif
}

//get single instance
func GetInterDB() *InterDB {
	_interOnce.Do(func() {
		_inter = NewInterDB()
	})
	return _inter
}

//construct
func NewInterDB() *InterDB {
	this := &InterDB{
		video2Gif: NewVideo2Gif(),
	}
	return this
}

//get sub db face
func (f *InterDB) GetVideo2GifDB() *Video2Gif {
	return f.video2Gif
}