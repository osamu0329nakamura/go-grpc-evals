package main

import (
    "net"
    "log"
    "google.golang.org/grpc"
    "context"
    pb "rpc"
)

type Order = pb.Order

type server struct {
    orders []*Order
    ch chan(*Order)
}

func (s *server) EnterOrder(ctx context.Context, incoming *pb.Order) (*pb.Response, error) {
    log.Printf("Enter Order %d", incoming.Price)
    s.orders = append(s.orders, incoming)
    s.ch <- incoming
    return &pb.Response{}, nil
}

func (s *server) ShowOrders(rqst *pb.ShowOrder, resp pb.OrderService_ShowOrdersServer) error {
    for o := range s.ch {
        resp.Send(o)
        if o.Price == 0 {
            break
        }
    }
    return nil
}

func main() {
    lis, err := net.Listen("tcp", ":50000")
    if err != nil {
        log.Fatal(err)
    }
    s := grpc.NewServer()
    pb.RegisterOrderServiceServer(s, &server{ch:make(chan *pb.Order)})

    if err := s.Serve(lis); err != nil {
        log.Panic(err)
    }
}
