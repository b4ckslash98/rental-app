package grpc

import (
	"context"
	"log"
	"net"
	pb "github.com/b4ckslash/rental-app/proto"
	"github.com/b4ckslash/rental-app/services/car/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CarServer struct {
	pb.UnimplementedCarServiceServer
	Repo repository.CarRepository
}

func (s *CarServer) CheckAvailability(ctx context.Context, req *pb.CarRequest) (*pb.CarResponse, error) {
	car, err := s.Repo.FindByID(int(req.CarId))
	return &pb.CarResponse{IsAvailable: err == nil && car != nil}, nil
}

func StartCarGRPC(repo repository.CarRepository) {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterCarServiceServer(server, &CarServer{Repo: repo})
	reflection.Register(server)
	log.Println("Car gRPC server listening on port 50052")
	server.Serve(lis)
}
