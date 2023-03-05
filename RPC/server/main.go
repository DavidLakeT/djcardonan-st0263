package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"

	pb "github.com/DavidLakeT/djcardonan-st0263/RPC/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedFilesInformationServer
}

func (s *server) ListFiles(ctx context.Context, requestInfo *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {

	log.Printf("Recibida petición de listado.")
	return &pb.ListFilesResponse{FileList: "Attack On Titan", FileAmount: 12}, nil
}

func (s *server) SearchFiles(ctx context.Context, requestInfo *pb.SearchFilesRequest) (*pb.SearchFilesResponse, error) {

	log.Printf("Recibida petición de búsqueda: %v", requestInfo)
	return &pb.SearchFilesResponse{Exists: searchFile(requestInfo.Pattern)}, nil
}

func main() {

	fmt.Println(listFiles())

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

func searchFile(name string) bool {

	if _, err := os.Stat("Files/" + name); err == nil {
		return true
	} else {
		return false
	}
}

func listFiles() (string, int) {

	names := make([]string, 0)
	amount := 0

	files, err := ioutil.ReadDir("Files/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		names = append(names, file.Name())
		amount += 1
	}

	return strings.Join(names[:], " , "), amount
}
