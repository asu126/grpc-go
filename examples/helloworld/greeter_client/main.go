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

package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-go/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

// server stream
func printFeatures(client pb.GreeterClient, rec *pb.HelloBytes) {
	log.Printf("Looking for features within %v", rec)
	stream, err := client.ServerStream(context.Background(), rec)
	if err != nil {
		log.Fatalf("%v.ServerStream(_) = _, %v", client, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ServerStream(_) = _, %v", client, err)
		}
		log.Println(feature)
	}
}

// client stream
func runRecordRoute(client pb.GreeterClient) {
	// Create a random number of random points
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2 // Traverse at least two points
	stream, err := client.ClientStream(context.Background())
	if err != nil {
		log.Fatalf("%v.RecordRoute(_) = _, %v", client, err)
	}
	for i := 0; i < pointCount; i++ {
		s := fmt.Sprintf("toutal count: %d", i)
		point := pb.HelloBytes{Message: []byte(s)}
		if err := stream.Send(&point); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, point, err)
		}
	}

	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Route summary: %v", reply)
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	r1, err := c.SayHelloAgain(context.Background(), &pb.HelloBytes{Message: []byte("request...")})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r1.Message)

	// 1
	printFeatures(c, &pb.HelloBytes{Message: []byte("001")})

	// ClientStream
	runRecordRoute(c)
}
