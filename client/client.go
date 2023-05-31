package main

import (
	"context"
	"log"

	pb "example.com/go-product-service-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	// Ürünü almak için bir istek oluştur
	req := &pb.Product{
		Id: 1,
	}

	// Sunucuya GetProduct RPC çağrısı yap
	res, err := client.GetProduct(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get product: %v", err)
	}

	// Sunucudan gelen ürünü yazdır
	log.Printf("Product: ID=%d, Name=%s, Price=%.2f", res.Id, res.Name, res.Price)
}
