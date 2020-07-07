package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "shift-api/grpc/shift/shift-api/grpc/shift"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("logging first test",
		"url", "keyvalue",
		"testval", 3,
		"backoff", time.Now(),
	)

	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewShiftClient(conn)
	message := &pb.GetShiftMessage{TargetUser: 1}
	res, err := client.GetShift(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
