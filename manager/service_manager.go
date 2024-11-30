package manager

import "github.com/project-sistem-voucher/api/service"

type ServiceManager interface {
	VoucherService() service.VoucherService
}

type serviceManager struct {
	repoManager RepoManager
}

func NewServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{
		repoManager: repo,
	}
}

func (m *serviceManager) VoucherService() service.VoucherService {
	return service.NewVoucherService(m.repoManager.VoucherRepo())
}
