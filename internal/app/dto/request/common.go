package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`           // 页码
	PageSize int `json:"page_size" form:"page_size"` // 每页大小
}

type Filter struct {
	Column   string      `json:"column" form:"column"`     // 列
	Operator string      `json:"operator" form:"operator"` // 运算符
	Value    interface{} `json:"value" form:"value"`       // 值
}

type GetIdReq struct {
	Id uint32 `json:"id" form:"id"`
}
