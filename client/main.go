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

	testUser := &user{"David", "david@example.com"}

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse: %v", err)
	}

	defer conn.Close()
	c := pb.NewUserInformationClient(conn)
	userInfo := &pb.UserInfo{Username: testUser.username, Email: testUser.email}
	response, err := c.GetUserInformation(context.Background(), userInfo)
	if err != nil {
		log.Fatalf("Error al obtener información de usuario: %v", err)
	}
	log.Printf("Cantidad de letras del nombre del usuario: %v", response.NameLength)
}
