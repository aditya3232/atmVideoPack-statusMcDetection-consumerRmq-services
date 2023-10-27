package tb_status_mc_detection

import (
	"strconv"
	"time"
)

// entity TbStatusMc
type TbStatusMc struct {
	ID            int       `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
	TidID         *int      `json:"tid_id"`
	DateTime      string    `json:"date_time"`
	StatusSignal  string    `json:"status_signal"`
	StatusStorage string    `json:"status_storage"`
	StatusRam     string    `json:"status_ram"`
	StatusCpu     string    `json:"status_cpu"`
}

func (m *TbStatusMc) TableName() string {
	return "tb_status_mc"
}

func (e *TbStatusMc) RedisKey() string {
	if e.ID == 0 {
		return "tb_status_mc"
	}

	return "tb_status_mc:" + strconv.Itoa(e.ID)
}
