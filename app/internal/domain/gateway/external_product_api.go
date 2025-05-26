package gateway

import "context"

type ScrapedProduct struct {
	Id       string        `json:"id"`
	Image    *ScrapedImage `json:"image"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
	Price    float64       `json:"price"`
}

type ScrapedImage struct {
	Thumbnail string `json:"thumbnail"`
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
}

type ExternalProductAPI interface {
	FetchProducts(ctx context.Context) ([]ScrapedProduct, error)
}
