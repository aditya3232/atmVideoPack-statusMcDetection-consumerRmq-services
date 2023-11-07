package consumer_status_mc_detection

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/log"
	"github.com/aditya3232/atmVideoPack-statusMcDetection-consumerRmq-services.git/model/add_status_mc_detection_to_elastic"
	esv7 "github.com/elastic/go-elasticsearch/v7"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Repository interface {
	ConsumerQueueStatusMcDetection() (RmqConsumerStatusMcDetection, error)
}

type repository struct {
	db            *gorm.DB
	rabbitmq      *amqp.Connection
	elasticsearch *esv7.Client
}

func NewRepository(db *gorm.DB, rabbitmq *amqp.Connection, elasticsearch *esv7.Client) *repository {
	return &repository{db, rabbitmq, elasticsearch}
}

func (r *repository) ConsumerQueueStatusMcDetection() (RmqConsumerStatusMcDetection, error) {
	var rmqConsumerStatusMcDetection RmqConsumerStatusMcDetection

	// create channel
	channel, err := r.rabbitmq.Channel()
	if err != nil {
		return rmqConsumerStatusMcDetection, err
	}
	defer channel.Close()

	// consume queue
	msgs, err := channel.Consume(
		"StatusMcDetectionQueue", // name queue
		"",                       // Consumer name (empty for random name)
		true,                     // Auto-acknowledgment (set to true for auto-ack)
		false,                    // Exclusive
		false,                    // No-local
		false,                    // No-wait
		nil,                      // Arguments
	)

	if err != nil {
		return rmqConsumerStatusMcDetection, err
	}

	// get message
	for d := range msgs {
		newStatusMcDetection := rmqConsumerStatusMcDetection
		err := json.Unmarshal(d.Body, &newStatusMcDetection)
		if err != nil {
			return rmqConsumerStatusMcDetection, err
		}

		// add data newHumanDtection to elasticsearch with CreateElasticHumanDetection
		repoElastic := add_status_mc_detection_to_elastic.NewRepository(r.elasticsearch)
		resultElastic, err := repoElastic.CreateElasticStatusMcDetection(
			add_status_mc_detection_to_elastic.ElasticStatusMcDetection{
				ID:            helper.DateTimeToStringWithStrip(time.Now()),
				Tid:           newStatusMcDetection.Tid,
				DateTime:      newStatusMcDetection.DateTime,
				StatusSignal:  newStatusMcDetection.StatusSignal,
				StatusStorage: newStatusMcDetection.StatusStorage,
				StatusRam:     newStatusMcDetection.StatusRam,
				StatusCpu:     newStatusMcDetection.StatusCpu,
			},
		)
		if err != nil {
			return rmqConsumerStatusMcDetection, err
		}
		// log result elastic
		log_function.Info(fmt.Sprintf("Result elastic: %v\n", resultElastic))

		// create data tb_human_detection
		// repo := tb_status_mc_detection.NewRepository(r.db)
		// _, err = repo.Create(
		// 	tb_status_mc_detection.TbStatusMc{
		// 		TidID:         newStatusMcDetection.TidID,
		// 		DateTime:      newStatusMcDetection.DateTime,
		// 		StatusSignal:  newStatusMcDetection.StatusSignal,
		// 		StatusStorage: newStatusMcDetection.StatusStorage,
		// 		StatusRam:     newStatusMcDetection.StatusRam,
		// 		StatusCpu:     newStatusMcDetection.StatusCpu,
		// 	},
		// )
		// if err != nil {
		// 	return rmqConsumerStatusMcDetection, err
		// }
	}

	return rmqConsumerStatusMcDetection, nil

}
