package util

import (
	"bufio"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"net"
	"net/url"
)

type SSHClient struct {
	ssh.ClientConfig
}

func (p *SSHClient) Connection(user, host, pass string) (*ssh.Client, *ssh.Session, error) {

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
		//需要验证服务端，不做验证返回nil就可以，点击HostKeyCallback看源码就知道了
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	signer, err := ssh.ParsePrivateKey([]byte(pass))

	if err == nil {

		sshConfig.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	}

	//sshConfig.SetDefaults()

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()

	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}

func (p *SSHClient) Command(host url.URL, account, password, shell string, channel chan<- []byte) {
	defer close(channel)

	_, session, err := p.Connection(account, host.Host, password)

	if err != nil {
		channel <- []byte("Error: Connection remote server error => " + err.Error())
		return
	}

	defer func() {
		if session != nil {
			session.Close()
		}
	}()

	stdout, err := session.StdoutPipe()

	if err != nil {
		channel <- []byte("Error: StdoutPipe error => " + err.Error())
		return
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		channel <- []byte("Error: StderrPipe error => " + err.Error())
		return
	}

	if err := session.Start(shell); err != nil {
		channel <- []byte("Error: Start error => " + err.Error())
		return
	}

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadBytes('\n')

		if err2 != nil || io.EOF == err2 {
			break
		}
		channel <- line
	}

	bytesErr, err := ioutil.ReadAll(stderr)

	if err == nil {
		channel <- bytesErr

	} else {

		channel <- []byte("Error: Stderr error => " + err.Error())
	}

	if err := session.Wait(); err != nil {

		channel <- []byte("Error: " + err.Error())
		return
	}
}
