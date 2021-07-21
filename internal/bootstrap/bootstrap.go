package bootstrap

import (
	"context"
	"log"

	"github.com/oniharnantyo/jogja-vaccine-scanner/config"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/adapter/service/email"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/adapter/service/slemankab/faskes"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/adapter/service/slemankab/schedule"
	"github.com/oniharnantyo/jogja-vaccine-scanner/internal/command"
)

func Run(conf *config.Config) {
	ctx := context.Background()

	restyClient := NewRestyClient(conf)

	//Slemankab Service
	faskesService := faskes.NewFaskesService(restyClient)
	scheduleService := schedule.NewScheduleService(restyClient)

	//Email Service
	emailService := email.NewEmailService(conf)

	//command
	checkQuotaCommand := command.NewCheckQuotaCommand(faskesService, scheduleService, emailService, conf)

	scheduler := NewSchedulerInitiator(checkQuotaCommand, conf.Scheduler)

	log.Println("Scheduler is running")
	scheduler.Init(ctx)
}
