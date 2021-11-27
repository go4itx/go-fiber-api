package constant

const (
	// 状态：#1启用 #2禁用
	StatusEnable  uint8 = 1
	StatusDisable uint8 = 2

	// 是否删除：#1是 #2否
	DeleteYes = 1
	DeleteNo  = 2

	// 默认分页数
	DefaultPageSize int = 10

	// time format tpl
	DateTpl     = "2006-01-02"
	DateTimeTpl = "2006-01-02 15:04:05"
)
