package main

import (
	"fmt"
	mcNet "github.com/Tnze/go-mc/net"
	"go.uber.org/zap"
)

func main() {
	connection, err := mcNet.DialRCON("127.0.0.1:25575", "T4Yk7QF2wV9kjdwcQfZc2CmwxLkInxfL")
	if err != nil {
		zap.Error(err)
		return
	}
	commandErr := connection.Cmd("kick boredOfEnnui")
	if commandErr != nil {
		zap.Error(commandErr)
		return
	}
	response, responseErr := connection.Resp()
	if responseErr != nil {
		zap.Error(responseErr)
		return
	}

	fmt.Println(response)
}
