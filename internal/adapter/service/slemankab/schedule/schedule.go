package schedule

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/adapter/service/slemankab"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/domain/schedule"

	"github.com/go-resty/resty/v2"
)

type ScheduleService struct {
	client *resty.Client
}

func (s *ScheduleService) Check(ctx context.Context, request *slemankab.ScheduleRequest) (schedule.Schedules, error) {
	res := new(slemankab.ScheduleResponse)
	req := s.client.R().SetContext(ctx).
		ForceContentType("application/json").
		SetResult(res).
		SetQueryString(request.ToQueryParams())

	resp, err := req.Get("/list-faskes")
	if err != nil {
		return nil, fmt.Errorf("%s: %q", "faskes", err)
	}

	if statusCode := resp.StatusCode(); statusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %q", "faskes", errors.New(http.StatusText(statusCode)))
	}

	return res.Result, nil
}

func NewScheduleService(client *resty.Client) slemankab.Schedule {
	return &ScheduleService{client: client}
}
