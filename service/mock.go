package service

import (
	"context"
	"errors"

	pb "shift-api/grpc/shift/shift-api/grpc/shift"
)

// ShiftMockService is the service
type ShiftMockService struct {
}

// GetShift can get user shift
func (s *ShiftMockService) GetShift(ctx context.Context, message *pb.GetShiftMessage) (*pb.ShiftResponse, error) {
	switch message.TargetUser {
	case 1:
		//まずはテスト
		return &pb.ShiftResponse{
			Name:  "工藤新一",
			Shift: 1,
		}, nil
	case 2:
		//まずはテスト
		return &pb.ShiftResponse{
			Name:  "毛利蘭",
			Shift: 2,
		}, nil
	}
	return nil, errors.New("Not Found Your Shift")
}
