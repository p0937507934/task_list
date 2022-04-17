package main

import (
	"fmt"
	"tasks_list/config"
	"tasks_list/route"
)

// gin-swagger middleware

func main() {
	r := route.Serve()
	err := r.Run(fmt.Sprintf(":%d", config.GlobalConfig.ServerPort))
	if err != nil {
		panic(err)
	}
}
