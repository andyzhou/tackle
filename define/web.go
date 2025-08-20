package define

const (
	WebSubPath     = "/web"
	WebTplPattern  = "**/*"
)

const (
	UriOfRoot = "/"
	UriOfApi  = "/api"
	UriOfHtml = "/html" //for css/js/image files
)

//web request groups
const (
	WebReqAppOfHome      = "home"
	WebReqAppOfVideo2gif = "video2gif" //video2gif
	WebReqAppOfPdf       = "pdf"       //pdf2html
)

const (
	//path para
	ParaOfSubApp    = "app"
	ParaOfSubModule = "module"
	ParaOfSubAct    = "act"
	ParaOfSubDataId = "dataId"
	ParaOfShortUri  = "uri"

	ParaOfMode = "mode"
	ParaOfPage = "page"

	SubPathParaOfAct = "act"
)

// inter code, global, DO NOT change these sequence!!!
const (
	CodeOfSuccess = iota + 1000
	CodeOfInvalidRequest
	CodeOfInvalidAppAndToken
	CodeOfInterError
	CodeOfNoSuchData
	CodeOfNoSuchPlayer
	CodeOfPlayerNickExists
	CodeOfGameHasExists
	CodeOfDataSaveFailed
	CodeOfDataOptFailed
	CodeOfNeedLogin	//1010
	CodeOfHasBadwords
)