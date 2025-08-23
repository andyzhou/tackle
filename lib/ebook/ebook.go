package ebook

import "sync"

/*
 * inter ebook face
 */

//global variable
var (
	_inter *InterEBook
	_interOnce sync.Once
)

//face info
type InterEBook struct {
}

//get single instance
func GetInterEBook() *InterEBook {
	_interOnce.Do(func() {
		_inter = NewInterEBook()
	})
	return _inter
}

//construct
func NewInterEBook() *InterEBook {
	this := &InterEBook{
	}
	return this
}

//get sub face