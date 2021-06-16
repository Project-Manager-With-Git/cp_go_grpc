package echo_serv

import (
	"context"
	"cp_go_grpc/echo_pb"
	"io"
	"math"

	log "github.com/Golang-Tools/loggerhelper"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//Square req-res模式方法实现模板
func (s *Server) Square(ctx context.Context, in *echo_pb.Message) (*echo_pb.Message, error) {
	log.Info("Square get message", log.Dict{"in": in})
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Info("get header from client", log.Dict{"header": md})
	}
	header := metadata.Pairs("header-key", "val")
	grpc.SendHeader(ctx, header) //发送header
	// create and set trailer
	defer func() {
		trailer := metadata.Pairs("trailer-key", "val")
		grpc.SetTrailer(ctx, trailer) //设置trailer在函数执行完成后发送
	}()
	m := &echo_pb.Message{Message: math.Pow(in.Message, 2)}
	log.Info("Square send message", log.Dict{"result": m})
	return m, nil
}

//RangeSquare  req-stream模式方法实现模板
func (s *Server) RangeSquare(in *echo_pb.Message, stream echo_pb.ECHO_RangeSquareServer) error {
	log.Info("RangeSquare get message ", log.Dict{"in": in})
	max := int(in.Message)
	i := 0
	for {
		if i >= max {
			break
		} else {
			m := &echo_pb.Message{Message: math.Pow(float64(i), 2)}
			stream.Send(m)
			log.Info("RangeSquare send message ", log.Dict{"result": m})
			i++
		}
	}
	log.Info("RangeSquare send message done")
	return nil
}

//SumSquare stream-res模式方法实现模板
func (s *Server) SumSquare(stream echo_pb.ECHO_SumSquareServer) error {
	sumResult := 0.0
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		log.Info("SumSquare get message ", log.Dict{"msg": msg})
		sumResult += math.Pow(msg.Message, 2)
	}
	m := &echo_pb.Message{Message: sumResult}
	log.Info("SumSquare send message ", log.Dict{"result": m})
	return stream.SendAndClose(m)
}

//StreamrangeSquare stream-stream模式方法实现模板
func (s *Server) StreamrangeSquare(stream echo_pb.ECHO_StreamrangeSquareServer) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	header := metadata.Pairs("header-key", "val")
	grpc.SendHeader(ctx, header) //发送header
	// create and set trailer
	defer func() {
		trailer := metadata.Pairs("trailer-key", "val")
		grpc.SetTrailer(ctx, trailer) //设置trailer在函数执行完成后发送
	}()
	if ok {
		log.Info("get header from client", log.Dict{"header": md})
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		log.Info("StreamrangeSquare get message ", log.Dict{"msg": msg})
		m := &echo_pb.Message{Message: math.Pow(msg.Message, 2)}
		log.Info("StreamrangeSquare send message ", log.Dict{"result": m})
		stream.Send(m)
	}
}
