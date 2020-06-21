package main

import (
	"context"
	"net/http"

	pb "shift-api/rpc/shift"
)

// ShiftTwirpServer twirpでのサーバー
type ShiftTwirpServer struct{}

// GetShift シフトを取得
func (s *ShiftTwirpServer) GetShift(ctx context.Context, req *pb.GetShiftMessage) (resp *pb.ShiftResponse, err error) {
	return &pb.ShiftResponse{Name: "twirp テスト", Shift: 1}, nil
}

func main() {
	twirpHandler := pb.NewShiftServer(&ShiftTwirpServer{}, nil)
	mux := http.NewServeMux()
	mux.Handle(pb.ShiftPathPrefix, twirpHandler)
	http.ListenAndServe(":8080", mux)
}
