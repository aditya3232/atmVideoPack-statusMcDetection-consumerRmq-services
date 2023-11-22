package main

import (
	"fmt"

	"github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/config"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/connection"

	// _ "github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/cron"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/helper"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/model/consumer_status_mc_detection"
	"github.com/gin-gonic/gin"
)

func main() {
	defer helper.RecoverPanic()

	forever := make(chan bool)
	go func() {
		defer helper.RecoverPanic() // Menambahkan recover di dalam goroutine

		consumerStatusMcDetectionRepository := consumer_status_mc_detection.NewRepository(connection.DatabaseMysql(), connection.RabbitMQ(), connection.ElasticSearch())
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

	// Penutupan koneksi setelah aplikasi selesai berjalan
	defer connection.Close()

}
