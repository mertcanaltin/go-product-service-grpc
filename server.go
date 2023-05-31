package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "example.com/go-product-service-grpc/pb" // Protobuf dosyasının paketini import ediyorum
)

type productServiceServer struct{}

func (s *productServiceServer) GetProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	logRequest(req)

	// Bu örnekte, hardcoded olarak bir ürün dönüyorum
	product := &pb.Product{
		Id:    1,
		Name:  "Example Product",
		Price: 9.99,
	}

	logResponse(product, nil)

	return product, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &productServiceServer{})

	log.Println("gRPC server started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func logRequest(req *pb.Product) {
	log.Printf("Received request: ID=%d, Name=%s, Price=%.2f", req.Id, req.Name, req.Price)
}

func logResponse(res *pb.Product, err error) {
	if err != nil {
		log.Printf("Error occurred: %v", err)
		return
	}
	log.Printf("Sent response: ID=%d, Name=%s, Price=%.2f", res.Id, res.Name, res.Price)
}
