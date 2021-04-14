package repositories

import (
	"database/sql"
	"shop/product/common"
	"shop/product/datamodels"
)

//先开发接口
//实现定义接口

type IProduct interface {
	//连接数据库
	Conn() error
	Insert(*datamodels.Product)(int64,error)
	Delete(int64) bool
	Update(*datamodels.Product) error
	SelectByKey(int64)(*datamodels.Product,error)
	SelectAll()([]*datamodels.Product,error)
}

type ProductManager struct {
	table string
	mysqlConn *sql.DB
}

func NewProductManager(table string,db *sql.DB) IProduct {
	return &ProductManager{
		table:     table,
		mysqlConn: db,
	}
}

func (p *ProductManager)Conn()error  {
	if p.mysqlConn==nil{
		mysql,err:= common.NewMysqlConn()
		if err != nil {
			return err
		}
		p.mysqlConn=mysql
	}
	if p.table != "" {
		p.table="product"
	}
	return nil
}

func (p *ProductManager)Insert(*datamodels.Product)(int64,error){
	return 0,nil
}
func (p *ProductManager)Delete(int64) bool{
	return true
}

func (p *ProductManager)Update(*datamodels.Product) error{
	return nil
}
func (p *ProductManager)SelectByKey(int64)(*datamodels.Product,error){
	return nil,nil
}
func (p *ProductManager)SelectAll()([]*datamodels.Product,error){
	return nil,nil
}

