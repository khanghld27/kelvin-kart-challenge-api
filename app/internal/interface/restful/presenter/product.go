package presenter

type ProductInformation struct {
	Id       string            `json:"id"`
	Image    *ImageInformation `json:"image"`
	Name     string            `json:"name"`
	Category string            `json:"category"`
	Price    float64           `json:"price"`
}

type ImageInformation struct {
	Thumbnail string `json:"thumbnail"`
	Mobile    string `json:"mobile"`
	Tablet    string `json:"tablet"`
	Desktop   string `json:"desktop"`
}
