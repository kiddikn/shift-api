package service

import (
	"context"
	"errors"

	pb "shift-api/grpc/shift/shift-api/grpc/shift"
)

// ShiftService is the service
type ShiftService struct {
}

// GetShift can get user shift
func (s *ShiftService) GetShift(ctx context.Context, message *pb.GetShiftMessage) (*pb.ShiftResponse, error) {
	switch message.TargetUser {
	case "1":
		//まずはテスト
		return &pb.ShiftResponse{
			Name:  "工藤新一",
			Shift: "am",
		}, nil
	case "2":
		//まずはテスト
		return &pb.ShiftResponse{
			Name:  "毛利蘭",
			Shift: "pm",
		}, nil
	}
	return nil, errors.New("Not Found Your Shift")
}
