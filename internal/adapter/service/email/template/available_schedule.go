package template

import (
	"strconv"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/adapter/service/email"

	"github.com/matcornic/hermes/v2"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/domain/schedule"
)

type AvailableScheduleTemplate struct {
	Schedules schedule.Schedules
}

func (t *AvailableScheduleTemplate) Generate() (string, error) {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "Available Vaccine",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: "Vaccine searcher",
			Intros: []string{
				"Here are available schedules:.",
			},
			Table: t.generateTable(),
			Outros: []string{
				"Grab it fast",
			},
		},
	}

	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return "", err
	}

	return emailBody, nil
}

func (t *AvailableScheduleTemplate) generateTable() hermes.Table {
	table := hermes.Table{
		Data: t.generateTableEntry(),
		Columns: hermes.Columns{
			CustomWidth: map[string]string{
				"Faskes":     "55%",
				"Date":       "25%",
				"Quota":      "10%",
				"Quota Left": "10%",
			},
		},
	}

	return table
}

func (t *AvailableScheduleTemplate) generateTableEntry() [][]hermes.Entry {
	var entries [][]hermes.Entry

	for _, s := range t.Schedules {
		entry := []hermes.Entry{
			{
				"Faskes",
				s.Faskes,
			},
			{
				"Date",
				s.Date,
			},
			{
				"Quota",
				strconv.Itoa(s.Quota),
			},
			{
				"Quota Left",
				strconv.Itoa(s.QuotaLeft),
			},
		}
		entries = append(entries, entry)
	}

	return entries
}

func NewAvailableScheduleTemplate(schedules schedule.Schedules) email.Generator {
	return &AvailableScheduleTemplate{Schedules: schedules}
}
