/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	// "time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc-go/examples/helloworld/helloworld"
	"grpc-go/examples/helloworld/helper"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloBytes) (*pb.HelloBytes, error) {
	b := []byte("123")
	return &pb.HelloBytes{Message: b}, nil
}

// ServerStream(*HelloBytes, Greeter_ServerStreamServer) error
func (s *server) ServerStream(in *pb.HelloBytes, stream pb.Greeter_ServerStreamServer) error {
	fmt.Println("ServerStream inputs %v", in)
	// for i := 0; i < 10; i++ {
	// 	s := fmt.Sprintf("%d", i)
	// 	stream.Send(&pb.HelloBytes{Message: []byte(s)})
	// }

	blobReader := strings.NewReader("some io.Reader stream to be read")

	sw := helper.NewWriter(func(p []byte) error {
		msg := &pb.HelloBytes{}
		msg.Message = p
		return stream.Send(msg)
	})

	_, err := io.CopyN(sw, blobReader, 10)
	if err != nil {
		// return status.Errorf(codes.Unavailable, "GetBlob: send: %v", err)
		return err
	}

	return nil
}

// ClientStream(Greeter_ClientStreamServer) error
func (s *server) ClientStream(stream pb.Greeter_ClientStreamServer) error {
	var count int
	// startTime := time.Now()
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			// endTime := time.Now()
			s := fmt.Sprintf("toutal count: %d", count)
			return stream.SendAndClose(&pb.HelloBytes{
				Message: []byte(s),
			})
		}
		if err != nil {
			return err
		}
		fmt.Println("ClientStream get: %v", point)
		count++
	}
}

// ServerAndClientStream(Greeter_ServerAndClientStreamServer) error
func (s *server) ServerAndClientStream(stream pb.Greeter_ServerAndClientStreamServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println("ClientStream get: %v", in)
		for i := 0; i < 3; i++ {
			s := fmt.Sprintf("ServerAndClientStream return: %d", i)
			if err := stream.Send(&pb.HelloBytes{Message: []byte(s)}); err != nil {
				return err
			}
		}
	}
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
