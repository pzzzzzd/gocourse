package main

import (
	"context"
	"log"
	mainapipb "simplegrpcclient/proto/gen"
	farewellpb "simplegrpcclient/proto/gen/farewell"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	cert := "cert.pem"

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalln("Did not connect:", err)
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("Did not connect:", err)
	}
	defer conn.Close()

	client := mainapipb.NewCalculateClient(conn)

	client2 := mainapipb.NewGreeterClient(conn)

	// fwClient := farewellpb.NewAufWiedersehenClient(conn)
	client3 := mainapipb.NewBigFarewellClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &mainapipb.AddRequest{
		A: 10,
		B: 20,
	}
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Fatalln("Could not add", err)
	}

	reqGreet := &mainapipb.HelloRequest{
		Name: "John",
	}
	res1, err := client2.Greet(ctx, reqGreet)
	if err != nil {
		log.Fatalln("Could not greet", err)
	}

	resAddFromGreeter, err := client2.Add(ctx, reqGreet)
	if err != nil {
		log.Fatalln("Could not add---", err)
	}

	reqGoodBye := &farewellpb.GoodByeRequest{
		Name: "Jane",
	}
	// resFw, err := fwClient.BidGoodBye(ctx, reqGoodBye)
	// if err != nil {
	// 	log.Fatalln("Could not bid Goodbye", err)
	// }

	res3, err := client3.BidGoodBye(ctx, reqGoodBye)
	if err != nil {
		log.Fatalln("Could not bid Goodbye", err)
	}

	log.Println("Sum:", res.Sum)
	log.Println("Greeting message:", res1.Message)
	log.Println("===============Greeting message from the second Add func in proto:", resAddFromGreeter.Message)
	// log.Println("Goodbye message:", resFw.Message)
	log.Println("Goodbye message:", res3.Message)
	state := conn.GetState()
	log.Println("Connection State:", state)

}
