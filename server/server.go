package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	proto "github.com/yaozhuangyanlingyu/account-srv/proto"
)

type Greeter struct{}

func (this *Greeter) SayHello(ctx context.Context, req *proto.SayRequest, rsp *proto.SayResponse) error {
	rsp.Greeting = "Hello " + req.Message
	return nil
}

func main() {
	// 创建新服务
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// 初始化方法
	service.Init()

	// 注册服务
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// 运行服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
