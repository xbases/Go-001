package service

import (
	"Week04/internal/biz"
	"context"
	"errors"

	"gorm.io/gorm"
)

// ZService service
type ZService struct {
	biz *biz.ZBiz
}

// NewZService new
func NewZService(biz *biz.ZBiz) *ZService {
	return &ZService{biz: biz}
}

func (s *ZService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	hello, err := s.biz.GetZ(in.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errcode.ErrNotFound
		}
	}
	return &pb.HelloReply{Message: hello.Name}, nil
}
