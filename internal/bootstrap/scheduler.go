package bootstrap

import (
	"context"
	"log"

	"github.com/jasonlvhit/gocron"

	"github.com/oniharnantyo/jogja-vaccine-scanner/config"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/command"
)

type Scheduler interface {
	Init(ctx context.Context)
}

type SchedulerInitiator struct {
	checkQuota command.CheckQuota
	config     config.Scheduler
}

func (i *SchedulerInitiator) Init(ctx context.Context) {
	s := gocron.NewScheduler()
	err := s.Every(uint64(i.config.Every.Minutes())).Minutes().Do(func() {
		if err := i.checkQuota.Check(ctx); err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		log.Println(err)
	}

	<-s.Start()
}

func NewSchedulerInitiator(quotaCommand command.CheckQuota, scheduler config.Scheduler) Scheduler {
	return &SchedulerInitiator{
		checkQuota: quotaCommand,
		config:     scheduler,
	}
}
