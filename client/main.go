package main

import (
	"context"
	"log"

	pb "github.com/DavidLakeT/djcardonan-st0263/proto"

	"google.golang.org/grpc"
)

type user struct {
	username string
	email    string
}

func main() {

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse: %v", err)
	}

	defer conn.Close()
	c := pb.NewFilesInformationClient(conn)
	userInfo := &pb.ListFilesRequest{}
	response, err := c.ListFiles(context.Background(), userInfo)
	if err != nil {
		log.Fatalf("Error al obtener información de usuario: %v", err)
	}
	log.Printf("Respuesta (cantidad): %v", response.FileAmount)
	log.Printf("Respuesta (lista): %v", response.FileList)
}
