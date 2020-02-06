package main

import (
	"github.com/kataras/iris"
)

func main() {
	//1.创建iris 实例
	app := iris.New()

	//2.设置模板
	app.HandleDir("/public", "./fronted/web/public")
	//3.访问生成好的html静态文件
	app.HandleDir("/html", "./fronted/web/htmlProductShow")

	app.Run(
		iris.Addr(":80"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}