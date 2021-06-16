package echo_sdk

import (
	"context"
	"cp_go_grpc/echo_pb"
	"fmt"
	"io"
	"os"

	log "github.com/Golang-Tools/loggerhelper"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//Main 执行测试
func (c *SDKConfig) Main() {
	fmt.Println(c)
	sdk := c.NewSDK()
	Conn, err := sdk.NewConn()
	if err != nil {
		log.Error("new comm get err", log.Dict{"err": err.Error()})
		os.Exit(1)
	}
	defer Conn.Close()
	//req-res
	ctx, cancel := sdk.NewCtx()
	defer cancel()
	md := metadata.Pairs("a", "1", "b", "2")
	ctx_meta := metadata.NewOutgoingContext(ctx, md)
	var header, trailer metadata.MD
	req, err := Conn.Square(ctx_meta, &echo_pb.Message{Message: 2.0},
		grpc.Header(&header),   // will retrieve header
		grpc.Trailer(&trailer), // will retrieve trailer
	)
	if err != nil {
		log.Error("Square get error", log.Dict{"err": err.Error()})
		os.Exit(1)
	}
	log.Info("Square get result", log.Dict{"header": header, "req": req, "trailer": trailer})

	//req-stream
	ResStream, err := Conn.RangeSquare(context.Background(), &echo_pb.Message{Message: 4.0})
	if err != nil {
		log.Error("RangeSquare get error", log.Dict{"err": err.Error()})
		os.Exit(1)
	}
	for {
		feature, err := ResStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Error("RangeSquare(_) = _", log.Dict{"err": err.Error()})
				os.Exit(1)
			}
		}
		log.Info("RangeSquare get res", log.Dict{"res": feature})
	}

	//stream-res
	ReqStream, err := Conn.SumSquare(context.Background())
	if err != nil {
		log.Error("SumSquare get error", log.Dict{"err": err.Error()})
		os.Exit(1)
	}
	for i := 0; i < 5; i++ {
		m := &echo_pb.Message{Message: float64(i)}
		err := ReqStream.Send(m)
		if err != nil {
			log.Error("SumSquare Send() get err", log.Dict{"m": m, "err": err.Error()})
			os.Exit(1)
		}
	}
	reply, err := ReqStream.CloseAndRecv()
	if err != nil {
		log.Error("ReqStream CloseAndRecv() got err", log.Dict{"err": err.Error()})
		os.Exit(1)
	}
	log.Info("ReqStream get summary", log.Dict{"reply": reply})

	//stream-stream
	ctx_stream_meta := metadata.NewOutgoingContext(context.Background(), md)
	ReqResStream, err := Conn.StreamrangeSquare(ctx_stream_meta)
	if err != nil {
		log.Error("StreamrangeSquare get err", log.Dict{"err": err.Error()})
		os.Exit(1)
	}

	// retrieve header
	streamheader, err := ReqResStream.Header()
	if err != nil {
		log.Error("StreamrangeSquare get header get error", log.Dict{"err": err.Error()})
		os.Exit(1)
	}
	log.Info("StreamrangeSquare get header", log.Dict{"header": streamheader})
	go func() {
		for i := 0; i < 5; i++ {
			m := &echo_pb.Message{Message: float64(i)}
			err := ReqResStream.Send(m)
			if err != nil {
				log.Error("StreamrangeSquare Send get error", log.Dict{"m": m, "err": err.Error()})
				os.Exit(1)
			}
		}
		err := ReqResStream.CloseSend()
		if err != nil {
			log.Error("StreamrangeSquare CloseSend get error", log.Dict{"err": err.Error()})
		}
	}()
	for {
		feature, err := ReqResStream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Error("StreamrangeSquare get err", log.Dict{"err": err.Error()})
				os.Exit(1)
			}
		}
		log.Info("StreamrangeSquare get result", log.Dict{"res": feature})

	}
	// retrieve trailer
	streamtrailer := ReqResStream.Trailer()
	log.Info("StreamrangeSquare get trailer", log.Dict{"trailer": streamtrailer})
}

var TestNode = SDKConfig{
	Address: []string{"localhost:5000"},
}
