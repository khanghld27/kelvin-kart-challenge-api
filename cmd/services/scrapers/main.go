package main

import (
	"encoding/json"
	"fmt"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/registry"
	"net/http"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/khanghld27/kelvin-kart-challenge-api/app/configs"
	"github.com/khanghld27/kelvin-kart-challenge-api/cmd/services/common"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/gormer"
	"github.com/sirupsen/logrus"
)

// Product represents the product structure from the API
type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Image    struct {
		Thumbnail string `json:"thumbnail"`
		Mobile    string `json:"mobile"`
		Tablet    string `json:"tablet"`
		Desktop   string `json:"desktop"`
	} `json:"image"`
}

// DBProduct represents the product structure in the database
type DBProduct struct {
	Id       string `gorm:"primaryKey"`
	Name     string
	Category string
	Price    float64
	Image    *DBImage `gorm:"foreignKey:ProductId"`
}

// DBImage represents the image structure in the database
type DBImage struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	ProductId string
	Thumbnail string
	Mobile    string
	Tablet    string
	Desktop   string
}

// TableName returns the table name for the product model
func (p *DBProduct) TableName() string {
	return "products"
}

// TableName returns the table name for the image model
func (i *DBImage) TableName() string {
	return "images"
}

func fetchProducts(url string) ([]Product, error) {
	// Create a client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Make the request
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decode the response
	var products []Product
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return products, nil
}

// createProduct creates a product in the database
func createProduct(db gormer.DBAdapter, apiProduct Product) error {
	// Create a new product with the data from the API
	product := &DBProduct{
		Id:       apiProduct.ID,
		Name:     apiProduct.Name,
		Category: apiProduct.Category,
		Price:    apiProduct.Price,
		Image: &DBImage{
			Thumbnail: apiProduct.Image.Thumbnail,
			Mobile:    apiProduct.Image.Mobile,
			Tablet:    apiProduct.Image.Tablet,
			Desktop:   apiProduct.Image.Desktop,
		},
	}

	// Create the product in the database
	if err := db.Gormer().Create(product).Error; err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

func main() {
	// Initialize configuration
	var cfg configs.Config
	if err := envconfig.Process("", &cfg); err != nil {
		logrus.Fatal(err)
	}
	configs.SetConfig(&cfg)

	// Initialize the database connection
	db := common.InitDBConnection(cfg.PostgreSQL)
	defer db.Close()

	// Crawl data from the API
	products, err := fetchProducts("https://orderfoodonline.deno.dev/api/product")
	if err != nil {
		logrus.Fatalf("Error fetching products: %v", err)
	}

	// Create the products in the database
	for _, product := range products {
		if err := createProduct(db, product); err != nil {
			logrus.Errorf("Error creating product %s: %v", product.ID, err)
		} else {
			logrus.Infof("Created product %s", product.ID)
		}
	}

	logrus.Info("Product scraping completed")

	productUseCase := registry.ProductHandler()
}
