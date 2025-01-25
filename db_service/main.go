package main

func main() {
	grpcServer := NewGRPCServer(":8181")
	grpcServer.Run()

	
}