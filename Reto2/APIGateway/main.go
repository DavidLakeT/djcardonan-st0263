package main

import (
	"context"
	"log"
	"strconv"
	"time"

	momclient "github.com/DavidLakeT/djcardonan-st0263/MOM/client"
	pb "github.com/DavidLakeT/djcardonan-st0263/RPC/proto"
	"github.com/streadway/amqp"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var rpc bool = true

func main() {

	config := GetConfig()

	r := gin.Default()

	connRPC, errRPC := grpc.Dial(config.IP+":"+config.PORT, grpc.WithInsecure())
	if errRPC != nil {
		log.Fatalf("Could not connect to gRPC server: %v", errRPC)
	}
	defer connRPC.Close()

	connMOM, errMOM := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if errMOM != nil {
		log.Fatalf("Could not connect to MOM server: %v", errRPC)
	}
	defer connMOM.Close()

	clienteRPC := pb.NewFilesInformationClient(connRPC)

	r.GET("/list", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if rpc {
			rpc = !rpc

			res, err := clienteRPC.ListFiles(ctx, &pb.ListFilesRequest{})
			if err != nil {
				log.Fatalf("Error when invoking ListFiles(): %v", err)
			}

			c.JSON(200, gin.H{
				"List of files (RPC)": res.FileList,
			})
		} else {
			rpc = !rpc

			res, err := momclient.RequestRPC("list")
			if err != nil {
				log.Fatalf("There was an error with the information received: %v", err)
			}

			c.JSON(200, gin.H{
				"List of files (MOM)": res,
			})
		}
	})

	r.GET("/search/:name", func(c *gin.Context) {

		name := c.Params.ByName("name")

		if rpc {
			rpc = !rpc

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			res, err := clienteRPC.SearchFiles(ctx, &pb.SearchFilesRequest{Pattern: name})
			if err != nil {
				log.Fatalf("Error when invoking SearchFile(): %v", err)
			}

			c.JSON(200, gin.H{
				"Found file (RPC)": strconv.FormatBool(res.Exists),
			})
		} else {
			rpc = !rpc

			res, err := momclient.RequestRPC(name)
			if err != nil {
				log.Fatalf("There was an error with the information received: %v", err)
			}

			c.JSON(200, gin.H{
				"Found file (MOM)": res,
			})
		}

	})

	r.Run(":8084") // Iniciar el servidor Gin
}
