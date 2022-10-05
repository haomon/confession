package sendmsg

type RequireString struct {
	StringMap map[string]string
}

var TGAlertText = RequireString{
	StringMap: map[string]string{
		"測試": "",
	},
}
