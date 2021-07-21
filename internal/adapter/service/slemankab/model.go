package slemankab

import (
	"github.com/google/go-querystring/query"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/domain/faskes"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/domain/schedule"
)

type ListResponse struct {
	Success bool            `json:"success"`
	Result  faskes.Faskeses `json:"result"`
}

type ScheduleRequest struct {
	FaskesID int    `url:"id_faskes"`
	NIK      string `url:"nik"`
	Age      int    `url:"umur"`
}

type ScheduleResponse struct {
	Success    bool               `json:"success"`
	Result     schedule.Schedules `json:"result"`
	FaskesData FaskesData         `json:"datafaskes"`
}

type FaskesData struct {
	Faskes string `json:"faskes"`
	Region string `json:"wilayah"`
}

func (r *ScheduleRequest) ToQueryParams() string {
	v, _ := query.Values(r)
	return v.Encode()
}
