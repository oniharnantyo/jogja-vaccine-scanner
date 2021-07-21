package bootstrap

import (
	"time"

	"github.com/oniharnantyo/jogja-vaccine-scanner/config"

	"github.com/go-resty/resty/v2"
)

func NewRestyClient(config *config.Config) *resty.Client {
	client := resty.New()
	client.SetTimeout(10 * time.Second)
	client.SetHostURL(config.Service.Slemankab.BaseUrl)
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"user-agent":   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36",
	})

	return client
}
