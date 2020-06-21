package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	pb "shift-api/rpc/shift"
)

func main() {
	client := pb.NewShiftProtobufClient("http://localhost:8080", &http.Client{})

	resp, err := client.GetShift(context.Background(), &pb.GetShiftMessage{TargetUser: 3})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.GetName())
	fmt.Println(resp.GetShift())
}
