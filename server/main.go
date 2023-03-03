package main

import (
	"context"
	"log"
	"net"

	pb "github.com/DavidLakeT/djcardonan-st0263/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedUserInformationServer
}

func (s *server) GetUserInformation(ctx context.Context, userInfo *pb.UserInfo) (*pb.UserResponse, error) {
	log.Printf("Recibido: %v", userInfo)
	letterAmount := len(userInfo.Username)
	return &pb.UserResponse{NameLength: int32(letterAmount)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserInformationServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
