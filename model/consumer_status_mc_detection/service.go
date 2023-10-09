package consumer_status_mc_detection

type Service interface {
	ConsumerQueueStatusMcDetection() (StatusMcDetection, error)
}

type service struct {
	statusMcDetectionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// consume and save to db
func (s *service) ConsumerQueueStatusMcDetection() (StatusMcDetection, error) {

	// consume queue
	statusMcDetection, err := s.statusMcDetectionRepository.ConsumerQueueStatusMcDetection()
	if err != nil {
		return StatusMcDetection{}, err
	}

	return statusMcDetection, nil

}
