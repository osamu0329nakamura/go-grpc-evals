package main

import (
    "fmt"
    pb "rpc"
    "context"
    "google.golang.org/grpc"
    "time"
    "io"
)

func enter(conn *grpc.ClientConn) {
    c := pb.NewOrderServiceClient(conn)

    ctx := context.TODO()
    for i := int32(10); i >= 0; i-- {
        time.Sleep(200 * time.Millisecond)
        order := &pb.Order{Price:i}
        fmt.Printf("New Order %d\n", i)
        _, err := c.EnterOrder(ctx, order)
        if err != nil {
            panic(err)
        }
    }
}

func show(conn *grpc.ClientConn) {
    c := pb.NewOrderServiceClient(conn)

    stream, err := c.ShowOrders(context.TODO(), &pb.ShowOrder{})
    if err != nil {
        panic(err)
    }
    for {
        o1, err := stream.Recv()
        if err != nil {
            if err == io.EOF {
                break
            }
            panic(err)
        }
        fmt.Println(o1)
    }
    conn.Close()
}

func main() {
    conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
    if err != nil {
        panic(err)
    }

    go show(conn)
    enter(conn)
}
