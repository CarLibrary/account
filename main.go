package main

import (
	"fmt"
	"github.com/CarLibrary/account/account"
	"github.com/CarLibrary/account/cache"
	"github.com/CarLibrary/account/model"
	pb "github.com/CarLibrary/proto/account"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {

	//加载配置
	model.InitMYSQL()
	cache.InitREDIS()
	fmt.Println("ok")

	//
	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &account.AccountServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
