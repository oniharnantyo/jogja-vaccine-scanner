package slemankab

import (
	"context"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/domain/schedule"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/domain/faskes"
)

type Faskes interface {
	List(ctx context.Context) (faskes.Faskeses, error)
}

type Schedule interface {
	Check(ctx context.Context, request *ScheduleRequest) (schedule.Schedules, error)
}
