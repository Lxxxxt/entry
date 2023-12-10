package main

import (
	"entry/internal"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default(server.WithHostPorts("localhost:8080"))
	internal.Startup()
	internal.RegisterRouter(h)
	h.Spin()
}
