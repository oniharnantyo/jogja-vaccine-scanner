package command

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/util/random"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/adapter/service/email"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/domain/schedule"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/domain/faskes"

	"github.com/oniharnantyo/jogja-vaccine-scanner/config"

	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/adapter/service/email/template"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/adapter/service/slemankab"
)

type CheckQuota interface {
	Check(ctx context.Context) error
}

type CheckQuotaCommand struct {
	Faskes   slemankab.Faskes
	Schedule slemankab.Schedule
	Email    email.Email
	Config   *config.Config
}

func (c *CheckQuotaCommand) Check(ctx context.Context) error {
	log.Println("Checking quota")
	faskeses, err := c.Faskes.List(ctx)
	if err != nil {
		return err
	}

	var schedules schedule.Schedules
	for _, faskes := range faskeses {
		availableSchedules, err := c.getAvailableSchedule(ctx, faskes)
		if err != nil {
			return err
		}

		schedules = append(schedules, availableSchedules...)
		time.Sleep(time.Duration(c.generateRandomSleep()) * time.Nanosecond)
	}

	c.logResult(schedules)

	if !c.checkSchedulesEmpty(schedules) {
		if err := c.sendEmail(schedules); err != nil {
			return err
		}
	}

	return nil
}

func (c *CheckQuotaCommand) getAvailableSchedule(ctx context.Context, faskes faskes.Faskes) (schedule.Schedules, error) {
	schedules, err := c.Schedule.Check(ctx, c.toScheduleRequest(faskes))
	if err != nil {
		return nil, err
	}

	var availableSchedules schedule.Schedules
	for _, schedule := range schedules {
		if schedule.QuotaLeft > 0 {
			availableSchedules = append(availableSchedules, schedule)
		}
	}

	return availableSchedules, nil
}

func (c *CheckQuotaCommand) toScheduleRequest(faskes faskes.Faskes) *slemankab.ScheduleRequest {
	return &slemankab.ScheduleRequest{
		FaskesID: faskes.ID,
		NIK:      c.Config.Participant.Nik,
		Age:      c.Config.Participant.Age,
	}
}

func (c *CheckQuotaCommand) checkSchedulesEmpty(schedules schedule.Schedules) bool {
	return len(schedules) == 0
}

func (c *CheckQuotaCommand) sendEmail(schedules schedule.Schedules) error {
	template, err := c.generateTemplate(schedules)
	if err != nil {
		return err
	}

	if err := c.Email.Send(template); err != nil {
		return err
	}

	return nil
}

func (c *CheckQuotaCommand) generateRandomSleep() int {
	randomizer := random.NewRandomizer()

	return randomizer.RandomByRange(100, 2000)
}

func (c *CheckQuotaCommand) generateTemplate(schedules schedule.Schedules) (string, error) {
	return template.NewAvailableScheduleTemplate(schedules).Generate()
}

func (c *CheckQuotaCommand) logResult(schedules schedule.Schedules) {
	if len(schedules) == 0 {
		log.Println("No available quota")
	} else {
		log.Println(fmt.Sprintf(`Wow... %d quota available`, len(schedules)))
	}
}

func NewCheckQuotaCommand(
	faskes slemankab.Faskes,
	schedule slemankab.Schedule,
	email email.Email,
	config *config.Config) CheckQuota {
	return &CheckQuotaCommand{
		Faskes:   faskes,
		Schedule: schedule,
		Email:    email,
		Config:   config,
	}
}
