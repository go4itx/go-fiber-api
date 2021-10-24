package resp

// result Uniform results
type result struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	ServerTime int64       `json:"serverTime"`
	Data       interface{} `json:"data"`
}
