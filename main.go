package main

import "db-go-gin/internal/app"

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn
//go:generate go mod tidy
//go:generate go mod download

// @title 文档
// @version 1.0
// @description 文档
// @termsOfService http://swagger.io/terms/
// @contact.name 
// @contact.url http://www.swagger.io/support
// @contact.email 624796905@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8888
// @BasePath /
func main() {
	// 初始化服务
	app.Start()

	// 优雅关闭连接
	defer app.Close()
}
