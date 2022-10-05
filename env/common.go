package env

import "net/http"

//GenHeader 取得基本Header 含Token
func GenHeader() http.Header {
	header := http.Header{}
	header.Add("SysToken", SysToken)
	return header

}

//GenHeaderMap 取得基本Header 含Token
func GenHeaderMap() (header map[string]string) {
	header["SysToken"] = SysToken
	return header

}
