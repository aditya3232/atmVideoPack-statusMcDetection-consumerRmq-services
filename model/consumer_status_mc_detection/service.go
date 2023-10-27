package consumer_status_mc_detection

type Service interface {
	ConsumerQueueStatusMcDetection() (RmqConsumerStatusMcDetection, error)
}

type service struct {
	statusMcDetectionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// consume and save to db
func (s *service) ConsumerQueueStatusMcDetection() (RmqConsumerStatusMcDetection, error) {

	// consume queue
	newRmqConsumerStatusMcDetection, err := s.statusMcDetectionRepository.ConsumerQueueStatusMcDetection()
	if err != nil {
		return newRmqConsumerStatusMcDetection, err
	}

	return newRmqConsumerStatusMcDetection, nil

}
