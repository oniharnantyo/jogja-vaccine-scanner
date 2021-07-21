package faskes

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/adapter/service/slemankab"

	"github.com/go-resty/resty/v2"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/domain/faskes"
)

type FaskesService struct {
	client *resty.Client
}

func (s *FaskesService) List(ctx context.Context) (faskes.Faskeses, error) {

	res := new(slemankab.ListResponse)
	req := s.client.R().SetContext(ctx).
		ForceContentType("application/json").
		SetResult(res)

	resp, err := req.Get("/list-faskes")
	if err != nil {
		return nil, fmt.Errorf("%s: %q", "faskes", err)
	}

	if statusCode := resp.StatusCode(); statusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %q", "faskes", errors.New(http.StatusText(statusCode)))
	}

	return res.Result, nil
}

func NewFaskesService(client *resty.Client) slemankab.Faskes {
	return &FaskesService{client: client}
}
