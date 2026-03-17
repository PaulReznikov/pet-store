package core

import "time"

type Product struct {
  ID int `json:"id"`
  Name string `json:"name"`
  ImgURL string `json:"img_url"`
  Description string `json:"description"`
  Price float64 `json:"price"`
  Sale ProductSale `json:"sale"`
  Animal Animal `json:"animal"`
  Brand Brand `json:"brand"`
  Packages []ProductPackages `json:"packages"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}

type ProductUnit string

const (
  UnitKg  ProductUnit = "кг"
	UnitG   ProductUnit = "г"
	UnitL   ProductUnit = "л"
	UnitMl  ProductUnit = "мл"
  UnitPcs ProductUnit = "шт"
)

type ProductPackages struct {
  ID int `json:"id"`
  ProductID int `json:"product_id"`
  Value int `json:"quantity"`
  Unit ProductUnit `json:"unit"`
  Count int `json:"count"`
}


type ProductSale struct {
  ID int `json:"id"`
  ProductID int `json:"product_id"`
  Quantity int `json:"quantity"`
  TotalPrice float64 `json:"total_price"`
  StartSale time.Time `json:"start_sale"`
  EndSale time.Time `json:"end_sale"`
}