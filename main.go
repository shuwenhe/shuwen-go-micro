package main

import (
	"fmt"

	"github.com/micro/go-micro/web"
)

func main() {
	add := func(o *web.Options) {
		o.Address = ":8001"
	}
	server := web.NewService(add)
	fmt.Println("server = ", server)
}
