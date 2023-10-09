package main

import (
	"fmt"

	"github.com/aditya3232/gatewatchApp-services.git/config"
	"github.com/aditya3232/gatewatchApp-services.git/connection"
	"github.com/aditya3232/gatewatchApp-services.git/helper"
	"github.com/aditya3232/gatewatchApp-services.git/model/consumer_status_mc_detection"
	"github.com/gin-gonic/gin"
)

func main() {
	defer helper.RecoverPanic()

	forever := make(chan bool)
	go func() {
		consumerStatusMcDetectionRepository := consumer_status_mc_detection.NewRepository(connection.DatabaseMysql(), connection.RabbitMQ())
		consumerStatusMcDetectionService := consumer_status_mc_detection.NewService(consumerStatusMcDetectionRepository)

		_, err := consumerStatusMcDetectionService.ConsumerQueueStatusMcDetection()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(" [*] - waiting for messages")
	<-forever

	router := gin.Default()
	if config.CONFIG.DEBUG == 0 {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Run(fmt.Sprintf("%s:%s", config.CONFIG.APP_HOST, config.CONFIG.APP_PORT))

}
