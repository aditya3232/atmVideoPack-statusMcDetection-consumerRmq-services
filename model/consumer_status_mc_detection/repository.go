package consumer_status_mc_detection

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Repository interface {
	ConsumerQueueStatusMcDetection() (StatusMcDetection, error)
	Create(statusMcDetection StatusMcDetection) (StatusMcDetection, error)
}

type repository struct {
	db       *gorm.DB
	rabbitmq *amqp.Connection
}

func NewRepository(db *gorm.DB, rabbitmq *amqp.Connection) *repository {
	return &repository{db, rabbitmq}
}

func (r *repository) ConsumerQueueStatusMcDetection() (StatusMcDetection, error) {

	// create channel
	channel, err := r.rabbitmq.Channel()
	if err != nil {
		return StatusMcDetection{}, err
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
		return StatusMcDetection{}, err
	}

	// get message
	for d := range msgs {
		statusMcDetection := StatusMcDetection{}
		err := json.Unmarshal(d.Body, &statusMcDetection)
		if err != nil {
			return StatusMcDetection{}, err
		}

		// insert
		_, err = r.Create(
			StatusMcDetection{
				Tid:           statusMcDetection.Tid,
				DateTime:      statusMcDetection.DateTime,
				StatusSignal:  statusMcDetection.StatusSignal,
				StatusStorage: statusMcDetection.StatusStorage,
				StatusRam:     statusMcDetection.StatusRam,
				StatusCpu:     statusMcDetection.StatusCpu,
			},
		)
		if err != nil {
			return StatusMcDetection{}, err
		}
	}

	return StatusMcDetection{}, nil

}

func (r *repository) Create(statusMcDetection StatusMcDetection) (StatusMcDetection, error) {
	err := r.db.Create(&statusMcDetection).Error
	if err != nil {
		return StatusMcDetection{}, err
	}

	return statusMcDetection, nil
}
