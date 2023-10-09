package consumer_status_mc_detection

// json di struct ini disesuaikan dengan key payload rmq
type StatusMcDetection struct {
	Tid           string `json:"Tid"`
	DateTime      string `json:"DateTime"`
	StatusSignal  string `json:"StatusSignal"`
	StatusStorage string `json:"StatusStorage"`
	StatusRam     string `json:"StatusRam"`
	StatusCpu     string `json:"StatusCpu"`
}

// table name
func (m *StatusMcDetection) TableName() string {
	return "tb_status_mc"
}
