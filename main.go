package main

import (
	mongodb "cobasatu/repository/mongo_repo"
	"cobasatu/transport/rest"
	"context"
	"log"
	"net"

	pb "cobasatu/transport/grpc/gen/proto"

	"google.golang.org/grpc"
)

type ProductApiServer struct {
	pb.UnimplementedProductApiServer
}

func (s *ProductApiServer) GetProducts(ctx context.Context, req *pb.PaginationRequest) (*pb.ProductResponse, error) {

	return &pb.ProductResponse{}, nil
}

func (s *ProductApiServer) GetProductById(ctx context.Context, req *pb.ProductIdRequest) (*pb.Product, error) {

	dataProduct := &pb.Product{
		Id:   req.ProductId,
		Name: "Tolak angin",
	}

	return dataProduct, nil
}

func main() {

	// Mongo Database
	mongodb.ConnectMongo()

	// Rest

	go func() {

		rest := rest.RestServer()
		rest.Static("assets", "./assets")
		rest.Run("localhost:8000")

	}()

	listen, err := net.Listen("tcp", "localhost:8081")

	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductApiServer(grpcServer, &ProductApiServer{})

	err = grpcServer.Serve(listen)

	if err != nil {
		log.Println(err)
	}

}
