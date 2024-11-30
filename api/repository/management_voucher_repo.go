package repository

import (
	"github.com/project-sistem-voucher/api/model"
	"gorm.io/gorm"
)

type VoucherRepository interface {
	CreateVoucher(voucher *model.Voucher) error
	FindByKodeVoucher(kode string) (*model.Voucher, error)
}

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepository{db: db}
}

func (r *voucherRepository) CreateVoucher(voucher *model.Voucher) error {
	return r.db.Create(voucher).Error
}

func (r *voucherRepository) FindByKodeVoucher(kode string) (*model.Voucher, error) {
	var voucher model.Voucher
	err := r.db.Where("kode_voucher = ?", kode).First(&voucher).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &voucher, nil
}
