package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ajaykakodia/go-grpc/basic"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05" + " " + string(bytes)))
}

func main() {
	log.SetFlags(0)

	log.SetOutput(&logWriter{})

	basic.BasicHello()
	//basic.BasicUser()
	//basic.ProtoToJsonUser()
	//basic.JsonToProtoUser()
	basic.UserGroupCall()
}
