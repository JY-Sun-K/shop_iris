package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"shop/common"
	"shop/fronted/web/controllers"
	"shop/repositories"
	"shop/service"
)

func main() {
	// 1. 创建iris实例
	app := iris.New()
	// 2. 设置debug模式
	app.Logger().SetLevel("debug")
	// 3. 设置注册模板
	template := iris.HTML(
		"./fronted/web/views", ".html").Layout(
		"shared/layout.html").Reload(true)
	app.RegisterView(template)
	// 4. 设置模板目标
	app.HandleDir("/public", "./fronted/web/public")
	//app.HandleDir("/html", "./fronted/web/htmlProductShow")
	// 出现异常跳转到指定页面
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Fatalln(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 5. 注册控制器
	// 注册user控制器
	userRepository := repositories.NewUserRepository("user", db)
	userService := service.NewUserService(userRepository)
	userPro := mvc.New(app.Party("/user"))
	//userPro.Register(userService, ctx, session.Start)
	userPro.Register(userService, ctx)
	userPro.Handle(new(controllers.UserController))

	productRepository := repositories.NewProductManager("product", db)
	//orderRepository := repositories.NewOrderManagerRepository("order", db)
	productService := service.NewProductService(productRepository)
	//orderService := service.NewOrderService(orderRepository)
	proProduct := app.Party("/product")
	//proProduct.Use(middleware.AuthConProduct) // 添加登录验证
	product := mvc.New(proProduct)
	//product.Register(productService,  ctx, session.Start)
	product.Register(productService,  ctx)
	product.Handle(new(controllers.ProductController))
	// 6. 启动服务
	app.Run(
		iris.Addr("127.0.0.1:8082"),
		//iris.WithoutVersionChecker,
		// 忽略服务器错误
		iris.WithoutServerError(iris.ErrServerClosed),
		// 尽可能优化
		iris.WithOptimizations,
	)
}