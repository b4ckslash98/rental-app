package grpc

import (
	"context"
	"log"
	"net"

	pb "github.com/b4ckslash98/rental-app/proto"
	"github.com/b4ckslash98/rental-app/services/user/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	Repo repository.UserRepository
}

func (s *UserServer) ValidateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := s.Repo.GetByID(int(req.UserId))
	return &pb.UserResponse{IsValid: err == nil && user != nil}, nil
}

func StartUserGRPC(repo repository.UserRepository) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserServer{Repo: repo})
	reflection.Register(server)
	log.Println("User gRPC server listening on port 50051")
	server.Serve(lis)
}
