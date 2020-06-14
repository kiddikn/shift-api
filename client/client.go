package main

import (
	"context"
	"fmt"
	"log"

	pb "shift-api/grpc/shift/shift-api/grpc/shift"

	"google.golang.org/grpc"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewShiftClient(conn)
	message := &pb.GetShiftMessage{TargetUser: 9}
	res, err := client.GetShift(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
