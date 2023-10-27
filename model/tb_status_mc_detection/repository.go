package tb_status_mc_detection

import "gorm.io/gorm"

type Repository interface {
	Create(tbStatusMc TbStatusMc) (TbStatusMc, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(tbStatusMc TbStatusMc) (TbStatusMc, error) {
	err := r.db.Create(&tbStatusMc).Error
	if err != nil {
		return tbStatusMc, err
	}

	return tbStatusMc, nil
}
