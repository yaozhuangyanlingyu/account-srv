package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	proto "github.com/yaozhuangyanlingyu/account-srv/proto"
)

func main() {
	// 实例化
	service := micro.NewService(
		micro.Name("greeter"),
	)

	// 初始化
	service.Init()

	client := proto.NewGreeterService("greeter", service.Client())
	res, err := client.SayHello(context.TODO(), &proto.SayRequest{Message: "张三"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
