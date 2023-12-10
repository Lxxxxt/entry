package internal

import (
	"entry/kitex_gen/user_service/userservice"
	"log"
	"os"

	"github.com/cloudwego/kitex/client"
)

var rpcClient userservice.Client

func Startup() {
	var err error
	rpcClient, err = userservice.NewClient("entry-rpc", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

}
func GetRpcClient() userservice.Client {
	return rpcClient
}
