package service

import (
	"context"
	"errors"
	"time"

	pb "shift-api/grpc/shift/shift-api/grpc/shift"

	"github.com/jinzhu/gorm"
)

// ShiftMySQLService is the service
type ShiftMySQLService struct {
}

// GetShift can get user shift
func (s *ShiftMySQLService) GetShift(ctx context.Context, message *pb.GetShiftMessage) (*pb.ShiftResponse, error) {
	db := gormConnect()
	defer db.Close()

	id := message.TargetUser
	// 構造体のインスタンス化
	shiftsEx := []shift{}
	// 指定した条件を元に複数のレコードを引っ張ってくる
	db.Find(&shiftsEx, "name_id=?", id)

	if len(shiftsEx) == 0 {
		return nil, errors.New("Not Found Your Shift")
	}

	// nameにコメント入れるという不毛なことをしているけど実験なので...
	// 一つしか返していないのもそういう理由
	return &pb.ShiftResponse{
		Name:  shiftsEx[0].Comment,
		Shift: shiftsEx[0].Shift,
	}, nil
}

func gormConnect() *gorm.DB {
	// テストなのでmampのデフォルト利用
	// TODO:ここを環境変数から読み込むようにする
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "otakeslsc_shift"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

type shift struct {
	ID        uint64
	ShiftDate *time.Time
	Shift     uint64
	Comment   string
	NameID    uint64
	Dinner    uint64
	Lunch     uint64
	Morning   uint64
}

func (s *shift) TableName() string {
	// テーブル名を指定しないと単数形でアクセスしにいくため
	return "shift_shift"
}
