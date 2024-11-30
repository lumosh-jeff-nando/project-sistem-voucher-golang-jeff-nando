package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/project-sistem-voucher/api/model"
	"github.com/project-sistem-voucher/api/service"
)

type VoucherHandler struct {
	Service service.VoucherService
}

func NewVoucherHandler(service service.VoucherService) *VoucherHandler {
	return &VoucherHandler{Service: service}
}

func (h *VoucherHandler) CreateVoucher(c *gin.Context) {
	var input model.Voucher
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil service untuk membuat voucher
	voucher, err := h.Service.CreateVoucher(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Voucher created successfully", "data": voucher})
}

func (h *VoucherHandler) DeleteVoucher(c *gin.Context) {
	idParam := c.Param("id")
	voucherID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "voucher_id tidak valid"})
		return
	}

	err = h.Service.DeleteVoucherByID(uint(voucherID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Voucher berhasil dihapus"})
}
