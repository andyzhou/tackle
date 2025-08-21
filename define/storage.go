package define

//sub storage path
const (
	StorageOfDB   = "db"
	StorageOfFile = "file"
)

//sqlite db file
const (
	SqliteFileOfVideo2Gif = "video2gif.db"
)

//sqlite table names
//for `video2gif`
const (
	TabNameOfVideo2GifIds     = `ids`
	TabNameOfVideo2GifUsers   = `users`
	TabNameOfVideo2GifFiles   = `files`
	TabNameOfVideo2GifOptLogs = `fileOptLogs`
)

//key table field
const (
	TabFieldOfScore    = "score"
	TabFieldOfCreateAt = "createAt"
)


//local file sub path
const (
	FilePathOfVideo2Gif = "file/video2gif"
)


//default ids
const (
	DefaultTableRecId = 0
)