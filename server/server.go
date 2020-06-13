package main

import (
	"log"
	"net"

	pb "shift-api/grpc/shift/shift-api/grpc/shift"

	"shift-api/service"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	shiftService := &service.ShiftService{}
	// 実行したい実処理をseverに登録する
	pb.RegisterShiftServer(server, shiftService)
	server.Serve(listenPort)
}
