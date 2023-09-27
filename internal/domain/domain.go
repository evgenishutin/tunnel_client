package domain

type ProxySSH struct {
	Username string
	Host     string
}

type ProxySSHService interface {
	StartProxy(params ProxySSH) error
	StopProxy() error
}
