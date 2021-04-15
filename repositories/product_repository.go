package repositories

import (
	"database/sql"
	"shop/common"
	"shop/datamodels"
	"strconv"
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
	SubProductNum( int64) error
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

func (p *ProductManager)Insert(product *datamodels.Product)(int64, error){
	if err:= p.Conn();err!=nil{
		return 0,err
	}
	sql:= "INSERT product SET productName =? ,productNum=?,productImage=?,productUrl=?"
	stmt,err := p.mysqlConn.Prepare(sql)
	if err!=nil {
		return 0,err
	}
	result,err := stmt.Exec(product.ProductName,product.ProductNum,product.ProductImage,product.ProductUrl)
	if err!=nil {
		return 0,err
	}

	productId,err := result.LastInsertId()
	return productId,err


}
func (p *ProductManager)Delete(productId int64) bool{
	if err:= p.Conn();err!=nil{
		return false
	}

	sql:="delete from product where ID =?"
	stmt,err := p.mysqlConn.Prepare(sql)
	if err!=nil {
		return false
	}
	_,err = stmt.Exec(productId)
	if err!=nil {
		return false
	}
	return true
}

func (p *ProductManager)Update(product *datamodels.Product) error{
	if err:= p.Conn();err!=nil{
		return err
	}

	sql:="Update product set productName=?,productNum=?,productImage=?,productUrl=? where ID ="+strconv.FormatInt(product.ID,10)
	stmt,err := p.mysqlConn.Prepare(sql)
	if err!=nil {
		return err
	}
	_,err = stmt.Exec(product.ProductName,product.ProductNum,product.ProductImage,product.ProductUrl)
	if err!=nil {
		return err
	}
	return nil


}
func (p *ProductManager) SelectByKey(productID int64) (*datamodels.Product,  error) {
	if err := p.Conn(); err != nil {
		return &datamodels.Product{}, err
	}
	sql := "select * from " + p.table + " where ID=" + strconv.FormatInt(productID, 10)
	row, errRow := p.mysqlConn.Query(sql)
	if errRow != nil {
		return &datamodels.Product{}, errRow
	}
	result := common.GetResultRow(row)
	if len(result) == 0 {
		return &datamodels.Product{}, nil
	}
	productResult:= &datamodels.Product{}
	common.DataToStructByTagSql(result, productResult)
	return productResult,nil
}

// 获取所有商品
func (p *ProductManager) SelectAll() (ProductArray []*datamodels.Product, errProduct error) {
	// 1. 判断连接是否存在
	if err := p.Conn(); err != nil {
		return nil, err
	}
	sql := "select * from " + p.table
	rows, err := p.mysqlConn.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	result := common.GetResultRows(rows)
	if len(result) == 0 {
		return nil, err
	}
	for _, v := range result {
		product := &datamodels.Product{}
		common.DataToStructByTagSql(v, product)
		ProductArray = append(ProductArray, product)
	}
	return
}

func (p *ProductManager) SubProductNum(productID int64) error {
	if err := p.Conn(); err != nil {
		return err
	}
	sql := "update " + p.table + " set " + "productNum=productNum-1 where ID=" + strconv.FormatInt(productID, 10)
	stmt, err := p.mysqlConn.Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}
