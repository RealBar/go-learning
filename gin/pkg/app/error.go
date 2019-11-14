package app

var codeMsgMap = map[int]string{
	RespCodeOK:            "OK",
	RespCodeInternalError: "internal error",
}

func GetCodeMsg(code int) string {
	return codeMsgMap[code]
}
