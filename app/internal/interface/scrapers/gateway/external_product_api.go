package gateway

import (
	"context"
	"encoding/json"
	domain "github.com/khanghld27/kelvin-kart-challenge-api/app/internal/domain/gateway"
	"github.com/khanghld27/kelvin-kart-challenge-api/pkg/logger"
	"io"
	"net/http"
)

type HTTPExternalProductAPI struct {
	BaseURL string
	Client  *http.Client
}

func NewHTTPExternalProductAPI(baseURL string) *HTTPExternalProductAPI {
	return &HTTPExternalProductAPI{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (api *HTTPExternalProductAPI) FetchProducts(ctx context.Context) ([]domain.ScrapedProduct, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", api.BaseURL+"/product", nil)
	if err != nil {
		return nil, err
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Debugf("Error closing response body: %v", err)
		}
	}(resp.Body)

	var products []domain.ScrapedProduct
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}

	return products, nil
}
