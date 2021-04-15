package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"shop/service"

	"shop/backend/web/controllers"

	"github.com/kataras/iris/v12/mvc"
	"log"
	"shop/common"
	"shop/repositories"
)

func main() {
	//创建iris实例
	app:= iris.New()
	//2.设置错误模式,在mvc模式下提示错误
	app.Logger().SetLevel("debug")
	//注册模板
	tmplate := iris.HTML(
		"./backend/web/views",
		".html",
		).Layout(
		"shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	//设置模板目标
	app.HandleDir("/assets","./backend/web/assets")
	//出现异常跳转指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message",ctx.Values().GetStringDefault("message","访问页面出错了"))
		ctx.ViewLayout("")
		ctx.View("shared/error_404.html")
	})

	db,err := common.NewMysqlConn()
	if err != nil {
		log.Fatalln(err)
	}

	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	//注册服务
	productRepository := repositories.NewProductManager("product", db)
	productService := service.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	//启动服务

	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		)



}
