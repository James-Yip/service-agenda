package main

import (
	"github.com/James-Yip/service-agenda/service/service"
)

const (
	port string = "8080"
)

func main() {
	server := service.NewServer()
	server.Run(":" + port)

}
