package echo_serv

import (
	"context"
	"cp_go_grpc/echo_pb"
	"io"
	"log"
	"math"
)

//Square req-res模式方法实现模板
func (s *Server) Square(ctx context.Context, in *echo_pb.Message) (*echo_pb.Message, error) {
	log.Println("Square get message ", in)
	m := &echo_pb.Message{Message: math.Pow(in.Message, 2)}
	log.Println("Square send message ", m)
	return m, nil
}

//RangeSquare  req-stream模式方法实现模板
func (s *Server) RangeSquare(in *echo_pb.Message, stream echo_pb.ECHO_RangeSquareServer) error {
	log.Println("RangeSquare get message ", in)
	max := int(in.Message)
	i := 0
	for {
		if i >= max {
			break
		} else {
			m := &echo_pb.Message{Message: math.Pow(float64(i), 2)}
			stream.Send(m)
			log.Println("RangeSquare send message ", m)
			i++
		}
	}
	log.Println("RangeSquare send message done")
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
		log.Println("SumSquare get message ", msg)
		sumResult += math.Pow(msg.Message, 2)
	}
	m := &echo_pb.Message{Message: sumResult}
	log.Println("SumSquare send message ", m)
	return stream.SendAndClose(m)
}

//StreamrangeSquare stream-stream模式方法实现模板
func (s *Server) StreamrangeSquare(stream echo_pb.ECHO_StreamrangeSquareServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		log.Println("StreamrangeSquare get message ", msg)
		m := &echo_pb.Message{Message: math.Pow(msg.Message, 2)}
		log.Println("StreamrangeSquare send message ", m)
		stream.Send(m)
	}
}
