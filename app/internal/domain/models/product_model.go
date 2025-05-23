package models

type Product struct {
	Id       string
	Image    *Image
	Name     string
	Category string
	Price    float64
}

func (p *Product) TableName() string {
	return "products"
}
