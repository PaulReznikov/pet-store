package models

import "time"

type Product struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Category    ProductCategory   `json:"category"`
	ImgURL      string            `json:"img_url"`
	Description string            `json:"description"`
	Price       float64           `json:"price"`
	Sale        *ProductSale      `json:"sale,omitempty"`
	Animal      Animal            `json:"animal"`
	Brand       Brand             `json:"brand"`
	Packages    []ProductPackages `json:"packages"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type ProductCategory struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID *int   `json:"parent_id,omitempty"`
}

type ProductPackages struct {
	ID        int         `json:"id"`
	ProductID int         `json:"product_id"`
	Value     float64     `json:"value"`
	Unit      ProductUnit `json:"unit"`
	Count     int         `json:"count"`
	Price     float64     `json:"price"`
}

type ProductSale struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	SalePercent float64   `json:"sale_percent"`
	StartSale   time.Time `json:"start_sale"`
	EndSale     time.Time `json:"end_sale"`
}
