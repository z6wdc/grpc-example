package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	greeter "grpc-example/proto"
	"log"
	"os"
)

var (
	scanner *bufio.Scanner
	client  greeter.GreetingServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	scanner = bufio.NewScanner(os.Stdin)

	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	client = greeter.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send Request")
		fmt.Println("2: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Hello()

		case "2":
			fmt.Println("bye.")
			goto M
		}
	}
M:
}

func Hello() {
	fmt.Println("Please enter your name.")
	scanner.Scan()
	name := scanner.Text()

	req := &greeter.HelloRequest{
		Name: name,
	}
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetMessage())
	}
}
