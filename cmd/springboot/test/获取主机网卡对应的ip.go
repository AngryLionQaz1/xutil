package main

import (
	"fmt"
	"net"
)

func main() {

	is, _ := ips()

	for k, v := range is {
		fmt.Println(k, v)
	}

}

func ips() (map[string]string, error) {

	ips := make(map[string]string)

	interfaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	for _, i := range interfaces {
		name, err := net.InterfaceByName(i.Name)
		if err != nil {
			return nil, err
		}
		add, err := name.Addrs()
		for _, v := range add {
			ips[name.Name] = v.String()
		}
	}

	return ips, nil

}
