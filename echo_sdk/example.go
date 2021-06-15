package echo_sdk

import (
	"context"
	"cp_go_grpc/echo_pb"
	"fmt"
	"io"
	"log"
)

//Main 执行测试
func (c *SDKConfig) Main() {
	fmt.Println(c)
	sdk := c.NewSDK()
	Conn, err := sdk.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer Conn.Close()
	//req-res
	ctx, cancel := sdk.NewCtx()
	defer cancel()
	req, err := Conn.Square(ctx, &echo_pb.Message{Message: 2.0})
	if err != nil {
		log.Fatalf("%v.Square get error %v", Conn, err)
		return
	}
	fmt.Println(req)
	//req-stream
	ResStream, err := Conn.RangeSquare(context.Background(), &echo_pb.Message{Message: 4.0})
	if err != nil {
		log.Fatalf("%v.RangeSquare get error %v", Conn, err)
		return
	}
	for {
		feature, err := ResStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("%v.RangeSquare(_) = _, %v", Conn, err)
			}
		}
		log.Println(feature)
	}

	//stream-res
	ReqStream, err := Conn.SumSquare(context.Background())
	if err != nil {
		log.Fatalf("%v.SumSquare get error, %v", Conn, err)
	}
	for i := 0; i < 5; i++ {
		m := &echo_pb.Message{Message: float64(i)}
		err := ReqStream.Send(m)
		if err != nil {
			log.Fatalf("%v.Send(%v) = %v", ReqStream, m, err)
		}
	}
	reply, err := ReqStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", ReqStream, err, nil)
	}
	log.Printf("summary: %v", reply)

	//stream-stream
	ReqResStream, err := Conn.StreamrangeSquare(context.Background())
	if err != nil {
		log.Fatalf("%v.StreamrangeSquare get error, %v", Conn, err)
	}
	go func() {
		for i := 0; i < 5; i++ {
			m := &echo_pb.Message{Message: float64(i)}
			err := ReqResStream.Send(m)
			if err != nil {
				log.Fatalf("%v.Send(%v) = %v", ReqResStream, m, err)
			}
		}
	}()
	for {
		feature, err := ReqResStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("%v.StreamrangeSquare(_) = _, %v", Conn, err)
			}
		}
		log.Println(feature)
	}
}
