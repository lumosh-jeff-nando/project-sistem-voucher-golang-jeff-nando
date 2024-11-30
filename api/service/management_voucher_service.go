package service

import (
	"errors"

	"github.com/project-sistem-voucher/api/model"
	"github.com/project-sistem-voucher/api/repository"
)

type VoucherService interface {
	CreateVoucher(input model.Voucher) (*model.Voucher, error)
	DeleteVoucherByID(voucherID uint) error
}

type voucherService struct {
	repo repository.VoucherRepository
}

func NewVoucherService(repo repository.VoucherRepository) VoucherService {
	return &voucherService{repo: repo}
}

func (s *voucherService) CreateVoucher(input model.Voucher) (*model.Voucher, error) {
	// Validasi masa aktif voucher
	if input.BerakhirBerlaku.Before(input.MulaiBerlaku) {
		return nil, errors.New("tanggal kadaluarsa tidak boleh sebelum tanggal mulai")
	}

	// Validasi kode voucher unik (opsional, jika perlu)
	existingVoucher, _ := s.repo.FindByKodeVoucher(input.KodeVoucher)
	if existingVoucher != nil {
		return nil, errors.New("kode voucher sudah digunakan")
	}

	// Simpan voucher ke repository
	err := s.repo.CreateVoucher(&input)
	if err != nil {
		return nil, err
	}

	return &input, nil
}

func (s *voucherService) DeleteVoucherByID(voucherID uint) error {
	voucher, err := s.repo.FindByID(voucherID)
	if err != nil {
		return err
	}
	if voucher == nil {
		return errors.New("voucher tidak ditemukan")
	}

	// Lakukan penghapusan
	return s.repo.DeleteVoucherByID(voucherID)
}
