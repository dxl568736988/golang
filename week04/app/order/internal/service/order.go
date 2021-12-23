package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"homework/week04/app/order/internal/biz"

	pb "homework/week04/api/order/service/v1"
)

type OrderService struct {
	pb.UnimplementedOrderServer
	oc *biz.OrderUseCase
	log *log.Helper
}

func NewOrderService(oc *biz.OrderUseCase, logger log.Logger) *OrderService {
	return &OrderService{
		oc: oc,
		log: log.NewHelper(log.With(logger, "module", "service/order")),
	}
}

func (s *OrderService) ListOrder(ctx context.Context, req *pb.ListOrderReq) (*pb.ListOrderReply, error) {
	rv, err := s.oc.List(ctx, req.PageNum, req.PageSize)
	rs := make([]*pb.ListOrderReply_Order, 0)
	for _, x := range rv {
		rs = append(rs, &pb.ListOrderReply_Order{
			Id: x.Id,
		})
	}
	return &pb.ListOrderReply{
		Orders: rs,
	}, err
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderReq) (*pb.CreateOrderReply, error) {
	x, err := s.oc.Create(ctx, req)
	return &pb.CreateOrderReply{
		Id: x.Id,
	}, err
}
func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderReq) (*pb.GetOrderReply, error) {
	x, err := s.oc.Get(ctx, req.Id)
	return &pb.GetOrderReply{
		Id: x.Id,
	}, err
}

func (s *OrderService) UpdateOrder(ctx context.Context, req *pb.UpdateOrderReq) (*pb.UpdateOrderReply, error) {
	x, err := s.oc.Update(ctx, req)
	return &pb.UpdateOrderReply{
		Id: x.Id,
	}, err
}
