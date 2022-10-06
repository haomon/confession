package env

var (
	Version string // Version

	Commit string // Commit

	BuildTime string // BuildTime

	SysToken string // SysToken

	Port string // Port

	Debug bool // Debug

	Pprof bool // Pprof

	Build string // Build

	ChatPrefix string

	LogchatID int64
	ChatID    int64
	TGToken   string

	WSReConnection int
	// 桌號前綴
	TidCode string
	// 訊息警報Title
	AlertTitle string

	Nuniskill map[int64]string
)
