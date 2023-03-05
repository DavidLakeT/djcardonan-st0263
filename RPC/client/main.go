package rpc

import (
	"context"
	"log"

	pb "github.com/DavidLakeT/djcardonan-st0263/RPC/proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error al conectarse: %v", err)
	}

	defer conn.Close()
	//c := pb.NewFilesInformationClient(conn)

	//Aquí se llaman a los métodos de abajo
}

func RequestSearchFile(client pb.FilesInformationClient, filename string) bool {

	filesInfo := &pb.SearchFilesRequest{Pattern: filename}
	response, err := client.SearchFiles(context.Background(), filesInfo)
	if err != nil {
		log.Fatalf("Error al obtener información de usuario: %v", err)
	}

	result := response.Exists

	log.Printf("Existencia del archivo: %v", response.Exists)
	return result
}

func requestListFiles(client pb.FilesInformationClient) string {

	filesInfo := &pb.ListFilesRequest{}
	response, err := client.ListFiles(context.Background(), filesInfo)
	if err != nil {
		log.Fatalf("Error al obtener información de usuario: %v", err)
	}

	resultList := response.FileList

	log.Printf("Nombres de archivos: %v", response.FileList)
	return resultList

}
