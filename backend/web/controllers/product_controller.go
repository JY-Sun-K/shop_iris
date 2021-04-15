package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"shop/common"
	"shop/datamodels"
	"shop/service"
	"strconv"
)

type ProductController struct {
	Ctx iris.Context
	ProductService service.IProductService
}

func (p *ProductController) GetAll()mvc.View {
	productArray,_:=p.ProductService.GetAllProduct()
	return mvc.View{
		Name: "product/view.html",
		Data: iris.Map{
			"productArray":productArray,
		},
	}
}

func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec:= common.NewDecoder(&common.DecoderOptions{
		TagName: "mTag",
	})
	if err := dec.Decode(p.Ctx.Request().Form,product);err!=nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	err:=p.ProductService.UpdateProduct(product)
	if err!=nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

// 添加商品Get
func (p *ProductController)GetAdd()mvc.View  {
	return mvc.View{
		Name:"product/add.html",
	}
}


// 添加商品Post
func (p *ProductController)PostAdd()  {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{"mTag", false, false})
	if err := dec.Decode(p.Ctx.Request().Form, product);err!= nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	_, err := p.ProductService.InsertProduct(product)
	if err!= nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

// 根据ID查询商品
func (p *ProductController)GetManager() mvc.View  {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 16)
	if err != nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	product, err := p.ProductService.GetProductByID(id)
	if err != nil{
		p.Ctx.Application().Logger().Debug(err)
	}
	return mvc.View{
		Name:"product/manager.html",
		Data:iris.Map{
			"product": product,
		},
	}
}

// 删除商品
func (p *ProductController)GetDelete()  {
	idString := p.Ctx.URLParam("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	isOk := p.ProductService.DeleteProductByID(id)
	if isOk {
		p.Ctx.Application().Logger().Debug("删除商品成功，ID为"+idString)
	}else {
		p.Ctx.Application().Logger().Debug("删除商品失败，ID为"+idString)
	}
	p.Ctx.Redirect("/product/all")
}