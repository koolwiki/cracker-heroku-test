package main

import (
	"fmt"
	"os"

	"github.com/koolwiki/cracker"
)

func main() {

	var port, secret string
	port = os.Getenv("PORT")
	secret = os.Getenv("SECRET")
	rpAddr = os.Getenv("RPADDR")
	if port == "" {
		port = "80"
	}
	if secret == "" {
		secret = "123456"
	}
	if rpAddr == "" {
		rpAddr = "http://mirrors.xmission.com/ubuntu/"
	}
	addr := fmt.Sprintf(":%s", port)
	p := cracker.NewHttpProxy(addr, secret, false, rpAddr)
	p.Listen()

}
