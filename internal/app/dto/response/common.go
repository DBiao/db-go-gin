package response

type PageResult struct {
	List     interface{} `json:"list"`      // 数据
	Total    int64       `json:"total"`     // 总数
	Page     int         `json:"page"`      // 页码
	PageSize int         `json:"page_size"` // 页数
}
