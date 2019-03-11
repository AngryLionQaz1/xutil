package main

import (
	"xutil/cmd/filter/account"
	"xutil/cmd/filter/proxy"
)

func main() {

	id := "100011"
	i := account.New(id, "萧毅", 100)
	i.Query(id)
	i.Update(id, 500)

}

func init() {

	account.New = func(id, name string, value int) account.Account {
		a := &account.AccountImpl{id, name, value}
		p := &proxy.Proxy{a}
		return p
	}

}
