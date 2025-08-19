package cmd

// flag field name
const (
	NameOfWeb		= "web"	 //web port
	NameOfConf      = "conf" //conf path
	NameOfLogPath   = "logPath"
	NameOfLogPrefix = "logPrefix"
)

//command config define
type (
	RunCfg struct {
		Web       int    //web port
		ConfPath  string //config path
		LogPath   string //log file path
		LogPrefix string //log file prefix
	}
)

