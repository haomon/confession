package caclient

type ConnectInfo struct {
	WSSUrl    string
	AuthToken string
}
type MsgSt struct {
	Action    string      `json:"action"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
	Mid       string      `json:"mid"`
}

type CAMaintenance struct {
	TId           int    `json:"tId"`
	HId           int    `json:"hId"`
	Memo          string `json:"memo"`
	IsMaintenance bool   `json:"isMaintenance"`
}
type HallInfo struct {
	IsMaintenance bool   `json:"isMaintenance"`
	Memo          string `json:"memo"`
	HId           int    `json:"hId"`
	ProductId     string `json:"ProductId"`
}
