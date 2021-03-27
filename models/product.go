package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	CategoryCode    string 	`json:"categoryCode" form:"category_code"`
	ProductsName    string 	`json:"productsName" form:"products_name"`
	BuyingPrice		int		`json:"buyingPrice" form:"buying_rice"`
	SellingPrice	int 	`json:"sellingPrice" from:"selling_price"`
	PriceOne		int 	`json:"priceOne" from:"price_one"`
	PriceTwo		int		`json:"priceTwo" from:"price_two"`
	PriceThree		int		`json:"priceThree" from:"price_three"`

}
