package ssh

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	config "ssh-proxy-app/config"

	"golang.org/x/crypto/ssh"
)

type Proxy struct {
	privateKeyPath string
	user           string
	host           string
	port           string
	protocol       string
	sshPort        string
}

type IProxy interface {
	Connection() error
}

func NewProxy(conf config.Config) *Proxy {
	return &Proxy{
		privateKeyPath: conf.PrivateKeyPath,
		user:           conf.User,
		host:           conf.Host,
		port:           conf.Port,
		protocol:       conf.Protocol,
		sshPort:        conf.SSHPort,
	}
}

func (p *Proxy) socks5Proxy(conn net.Conn) {
	defer conn.Close()

	var b [1024]byte

	n, err := conn.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	conn.Write([]byte{0x05, 0x00})

	n, err = conn.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	var addr string
	switch b[3] {
	case 0x01:
		sip := sockIP{}
		if err := binary.Read(bytes.NewReader(b[4:n]), binary.BigEndian, &sip); err != nil {
			return
		}
		addr = sip.toAddr()
	case 0x03:
		host := string(b[5 : n-2])
		var port uint16
		err = binary.Read(bytes.NewReader(b[n-2:n]), binary.BigEndian, &port)
		if err != nil {
			log.Println(err)
			return
		}
		addr = fmt.Sprintf("%s:%d", host, port)
	}

	server, err := client.Dial(p.protocol, addr)
	if err != nil {
		log.Println(err)
		return
	}
	conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	go io.Copy(server, conn)
	io.Copy(conn, server)
}

type sockIP struct {
	A, B, C, D byte
	PORT       uint16
}

func (ip sockIP) toAddr() string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", ip.A, ip.B, ip.C, ip.D, ip.PORT)
}

func (p *Proxy) socks5ProxyStart() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	server, err := net.Listen(p.protocol, ":"+p.port)
	if err != nil {
		log.Println(err)
		return
	}
	defer server.Close()
	log.Println("Tcp server listens on port 1080")
	for {
		client, err := server.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go p.socks5Proxy(client)
	}
}

var client *ssh.Client

func (p *Proxy) Connection() error {
	b, err := os.ReadFile(p.privateKeyPath)
	if err != nil {
		return err
	}
	pKey, err := ssh.ParsePrivateKey(b)
	if err != nil {
		return err
	}
	config := ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            p.user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(pKey),
		},
	}
	client, err = ssh.Dial(p.protocol, p.host+":"+p.sshPort, &config)
	if err != nil {
		return err
	}
	log.Println("")
	defer client.Close()
	client.NewSession()
	p.socks5ProxyStart()
	return nil
}
