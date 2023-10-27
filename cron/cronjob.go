package cron

import (
	"sync"

	"github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/connection"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/log"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/model/del_old_log_from_elastic"
	"github.com/robfig/cron/v3"
)

func init() {
	log_function.Info("Cron Job Started")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer helper.RecoverPanic()
		defer wg.Done()

		cron := cron.New(cron.WithChain(
			cron.SkipIfStillRunning(cron.DefaultLogger),
		))

		cron.AddFunc("0 0 * * *", func() {
			delOldLogFromElasticRepository := del_old_log_from_elastic.NewRepository(connection.ElasticSearch())
			delOldLogFromElasticService := del_old_log_from_elastic.NewService(delOldLogFromElasticRepository)

			err := delOldLogFromElasticService.DelOneMonthOldLogs()
			if err != nil {
				log_function.Error("Error delete log:", err)
			}

			log_function.Info("delete log in elastic berhasil dilakukan")
		})

		cron.Start()
	}()

	wg.Wait()

}
