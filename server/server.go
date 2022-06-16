package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	greet "github.com/Dinesh789kumar12/ServerSideStreamingwithGRPC/greet"
	"google.golang.org/grpc"
)

type server struct {
	greet.UnimplementedGreetServiceServer
}

func (*server) GreetManyTimes(req *greet.GreetManyTimesRequest, stream greet.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes Function invoked with request:%v", req)
	fName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello" + fName + "number" + strconv.Itoa(i)
		res := &greet.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		log.Printf("Sent:%v", res)
		time.Sleep(1000 * time.Millisecond)

	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server Started listening on Port:0.0.0.0:50052")
	c := grpc.NewServer()
	greet.RegisterGreetServiceServer(c, &server{})
	if err := c.Serve(lis); err != nil {
		log.Fatalln(err)
	}

}
