package service

import (
	"db-go-gin/internal/app/dto/request"

	"github.com/spf13/cast"
)

func Paginate(page int, pageSize int) (int, int) {
	if page == 0 {
		page = 1
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	page = (page - 1) * pageSize

	return page, pageSize
}

// SelectScope 筛选条件
func SelectScope(filter []request.Filter) map[string]interface{} {
	data := make(map[string]interface{})
	for k := range filter {
		if filter[k].Value == "" {
			continue
		}

		var key string
		var value interface{}
		if filter[k].Operator == "LIKE" || filter[k].Operator == "like" {
			key = filter[k].Column + " " + filter[k].Operator + " ?"
			value = "%" + cast.ToString(filter[k].Value) + "%"
		} else {
			key = filter[k].Column + " " + filter[k].Operator + " ?"
			value = filter[k].Value
		}
		data[key] = value
	}

	return data
}
