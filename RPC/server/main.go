package main

import (
	"context"
	"log"
	"net"

	pb "github.com/DavidLakeT/djcardonan-st0263/RPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedFilesInformationServer
}

func (s *server) ListFiles(ctx context.Context, requestInfo *pb.ListFilesRequest) (*pb.FilesResponse, error) {

	log.Printf("Recibido: %v", requestInfo)
	return &pb.FilesResponse{FileList: "Attack On Titan", FileAmount: 12}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFilesInformationServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
