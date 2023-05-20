package res

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total int64       `json:"total"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
