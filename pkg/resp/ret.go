package resp

// Ret Uniform result
type Ret struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	ServerTime int64       `json:"serverTime"`
	Data       interface{} `json:"data"`
}

type PaginationRet struct {
	Items interface{} `json:"items"`
	Count interface{} `json:"count"`
	Ext   interface{} `json:"ext,omitempty"`
}
