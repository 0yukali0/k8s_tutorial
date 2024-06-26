package main

import (
	"fmt"
	"heroroad/pkg/webservice"
)

func main() {
	server := webservice.GetGin()
	if err := server.Run(); err != nil {
		fmt.Printf("server encouter in %s\n", err.Error())
	}
}
