package define


//inter code, global, DO NOT change these sequence!!!
const (
	CodeSuccess = iota + 1000
	CodeInterError //1001
	CodeLostParameter
	CodeInvalidParam
	CodeParseFormFailed
	CodeGetConfFailed
	CodeGetInterServiceFailed
	CodeInvalidData
	CodeInvalidCaptcha
	CodeInvalidOpt
	CodeNeedLogin //1010
	CodeInterFailed
	CodeHasBadWords
	CodeReadFileFailed
	CodeDataQueryFailed
	CodeDataSaveFailed
	CodeDataDelFailed
	CodeNoSuchData
	CodeDataHasExists
	CodeNotAllowSelf
	CodeBeBlocked //1020
	CodeBeBanned
	CodeNoAccess
	CodeOptFailed
	CodeIpLimited
	CodePostLimited
	CodeSameValue
	CodeInvalidApi
	CodeInvalidFunc
	CodeUploadFileFailed
	CodeInvalidAppAndToken
	CodeDocHasQuoted
)
