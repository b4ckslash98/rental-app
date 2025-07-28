package carclient

import "fmt"

type CarGRPCClient interface {
	CheckAvailability(carID int) bool
}

type carGRPCClient struct{}

func New() CarGRPCClient {
	return &carGRPCClient{}
}

func (c *carGRPCClient) CheckAvailability(carID int) bool {
	// Simulasi panggilan gRPC
	fmt.Println("[gRPC] Checking availability for car", carID)
	return true // anggap tersedia untuk stub awal
}
