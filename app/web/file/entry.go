package file

import (
	"github.com/andyzhou/tackle/define"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * file entry
 * - include doc file, etc.
 */
type InterEntry struct {
	//sub face
	video2Gif *Video2GifEntry
}

//construct
func NewInterEntry() *InterEntry {
	//self init
	this := &InterEntry{
		video2Gif: NewVideo2GifEntry(),
	}
	return this
}

//entry
func (f *InterEntry) Entry(
	cookieUserId int64,
	ctx *gin.Context) {
	//get core path parameter
	subModule := ctx.Param(define.ParaOfSubModule) //p1

	//call diff sub face
	switch subModule {
	case define.SubAppOfVideo2Gif:
		f.video2Gif.Entry(cookieUserId, ctx)
		break
	default:
		//un-support sub face
		ctx.String(http.StatusOK, "un-support sub face")
		break
	}
}