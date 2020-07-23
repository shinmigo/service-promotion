package main

import (
	grpcserver "goshop/service-promotion/pkg/grpc/server"
)

func initService() {
	go grpcserver.Run()
	//go user.Hello()
}
