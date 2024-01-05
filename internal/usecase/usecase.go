package usecase

import (
	"log"
	"net"

	proxy "ssh-proxy-app/pkg/proxy"
)

type ProxyUseCase struct {
	service proxy.IProxy
}

func NewProxyUseCase(service proxy.IProxy) *ProxyUseCase {
	return &ProxyUseCase{
		service: service,
	}
}

func (uc *ProxyUseCase) StartProxy() {
	err := uc.service.Connection()
	if err != nil {
		log.Println(err)
		return
	}
}

func (uc *ProxyUseCase) StopProxy() {
	// TO DO
}

func (uc *ProxyUseCase) IsValidIPv4(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil && parsedIP.To4() != nil
}
