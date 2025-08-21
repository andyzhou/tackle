package define

const (
	ApiResponseValLen = 3 //ajax, code, error
	WebResponseValLen = 2 //page, error
)

//sub face of entry
const (
	//for web
	SubFaceOfHome      = "home"
	SubFaceOfList	   = "list"
	SubFaceOfPost	   = "post"
	SubFaceOfUpload    = "upload"
)

//sub tpl file
const (
	//for global
	TplOfNotFound   = "404.html"
	TplOfeNoAccess  = "no_access.html"
	TplOfGlobalMain = "global_main.html"

	//for home page
	TplOfPageHome  = "home_home.html"
	TplOfPageSetup  = "page_setup.html"

	//for video2gif page
	TplOfVideo2GifHome = "video2gif_home.html"
	TplOfVideo2GifList = "video2gif_list.html"
	TplOfVideo2GifPost = "video2gif_post.html"
)