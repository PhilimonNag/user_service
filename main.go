package main

import "fmt"

func main() {
	fmt.Println("Jesus not bring me this far to leave")
	grpcServer := NewGrpcServer(":8000")
	grpcServer.Run()
}
