package main

import (
	"context"
	"fmt"
	"io"
	"log"

	greet "github.com/Dinesh789kumar12/ServerSideStreamingwithGRPC/greet"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect:%v", err)
	}
	defer con.Close()
	s := greet.NewGreetServiceClient(con)
	req := greet.GreetManyTimesRequest{
		Greeting: &greet.Greeting{
			FirstName: "Dinesh",
			LastName:  "Kumar",
		},
	}
	res, err := s.GreetManyTimes(context.Background(), &req)
	if err != nil {
		log.Fatalf("error occured while calling GreetManyTimes")
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream %v", err)
		}
		fmt.Printf("Response: %v\n", msg.GetResult())
	}

}
