package userclient

import "fmt"

type UserGRPCClient interface {
	ValidateUser(userID int) bool
}

type userGRPCClient struct{}

func New() UserGRPCClient {
	return &userGRPCClient{}
}

func (c *userGRPCClient) ValidateUser(userID int) bool {
	// Simulasi panggilan gRPC
	fmt.Println("[gRPC] Validating user", userID)
	return true // anggap valid untuk stub awal
}
