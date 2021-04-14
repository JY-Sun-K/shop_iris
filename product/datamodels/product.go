package datamodels

type Product struct {
	ID int64 `json:"id" sql:"ID" mTag:"id"`
	ProductName string `json:"ProductName" sql:"productName" mTag:"ProductName"`
	ProductNum string `json:"ProductNum" sql:"productNum" mTag:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" mTag:"ProductImage"`
	ProductUrl string `json:"ProductUrl" sql:"productUrl" mTag:"ProductUrl"`

}
