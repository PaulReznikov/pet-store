package core

import "time"

type Product struct {
  ID int `json:"id"`
  Name string `json:"name"`
  Category ProductCategory `json:"category"`
  ImgURL string `json:"img_url"`
  Description string `json:"description"`
  Price float64 `json:"price"`
  Sale *ProductSale `json:"sale,omitempty"`
  Animal Animal `json:"animal"`
  Brand Brand `json:"brand"`
  Packages []ProductPackages `json:"packages"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}


type ProductCategory struct {
  ID int `json:"id"`
  Name string `json:"name"`
  ParentID *int `json:"parent_id,omitempty"`
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
  // Value - фактическое значение фасовки
  Value float64 `json:"value"`
  // Unit - единица измерения
  Unit ProductUnit `json:"unit"`
  // Count - количество товара в наличии для данной фасовки
  Count int `json:"count"`
  // Price - цена товара в данной фасовке
  Price float64 `json:"price"`
}


type ProductSale struct {
  ID int `json:"id"`
  Name string `json:"name"`
  SalePercent float64 `json:"sale_percent"`
  StartSale time.Time `json:"start_sale"`
  EndSale time.Time `json:"end_sale"`
}