package service

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"ssh-proxy-app/internal/domain"
)

type SSHProxyService struct{}

func NewSSHProxyService() *SSHProxyService {
	return &SSHProxyService{}
}

func (r *SSHProxyService) StartProxy(params domain.ProxySSH) (*exec.Cmd, error) {
	cmd := exec.Command("ssh", "-N", "-D", "1080", params.Username+"@"+params.Host) // "-f",
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		fmt.Println("ERROR:", err)
		log.Fatal(err)
		return &exec.Cmd{}, err
	}

	pid := cmd.Process.Pid
	fmt.Printf("PID процесса SSH: %d\n", pid)

	// cmd.Wait()
	return cmd, nil
}
