package app

import (
	"fmt"
	"github.com/andyzhou/tackle/app/api"
	"github.com/andyzhou/tackle/define"
	"github.com/gin-gonic/gin"
)

/*
 * web api interface entry
 * - used for ajax, client, etc.
 * - start like /api/xxx format
 */

//web entry info
type ApiEntry struct {
	g *gin.Engine
	apiEntry *api.InterApiEntry
}

//construct
func NewApiEntry(g *gin.Engine) *ApiEntry {
	//self init
	this := &ApiEntry{
		g: g,
		apiEntry: api.NewInterApiEntry(),
	}
	//init entry request
	this.interInit()
	return this
}

//main entry
func (f *ApiEntry) Entry(ctx *gin.Context) {
	//call sub entry
	f.apiEntry.Entry(ctx)
}

//entry
func (f *ApiEntry) interInit() {
	//like:  /api/[subApp]/[module]/[dataId]
	apiRootReqUrl := fmt.Sprintf("/%v", define.UriOfApi)
	apiSubAppUrl := fmt.Sprintf("%v/:%v",
		apiRootReqUrl,
		define.ParaOfSubApp,
	)
	webSubActUrl := fmt.Sprintf("%v/:%v/:%v",
		apiRootReqUrl,
		define.ParaOfSubApp,
		define.ParaOfSubAct,
	)

	//register request entry
	f.g.Any(apiRootReqUrl, f.apiEntry.Entry)
	f.g.Any(apiSubAppUrl, f.apiEntry.Entry)
	f.g.Any(webSubActUrl, f.apiEntry.Entry)
}