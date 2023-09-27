package usecase

import (
	"os"
	"os/exec"

	"ssh-proxy-app/internal/domain"
	"ssh-proxy-app/internal/service"
)

type ProxySSHUseCase struct {
	service service.SSHProxyService
	process *exec.Cmd
}

func NewProxySSHUseCase(serv service.SSHProxyService) *ProxySSHUseCase {
	return &ProxySSHUseCase{
		service: serv,
	}
}

func (uc *ProxySSHUseCase) SetParams(username, host string) *domain.ProxySSH {
	return &domain.ProxySSH{
		Username: username,
		Host:     host,
	}
}

func (uc *ProxySSHUseCase) StartProxy(params domain.ProxySSH) error {
	cmd, err := uc.service.StartProxy(params)
	uc.process = cmd
	if err != nil {
		return err
	}
	return nil
}

func (uc *ProxySSHUseCase) StopProxy() error {
	proc := uc.process.Process
	if err := proc.Signal(os.Interrupt); err != nil {
		return err
	}
	return nil
}
