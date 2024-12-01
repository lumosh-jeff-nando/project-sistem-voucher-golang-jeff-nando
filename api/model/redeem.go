package model

import "time"

type Redeem struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `json:"user_id"`
	KodeVoucher   string    `json:"kode_voucher"`
	TanggalRedeem time.Time `json:"tanggal_redeem"`
}
