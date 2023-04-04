package main

import (
	"context"
	"grpc-tennis/config"
	"grpc-tennis/database"
	"grpc-tennis/location"
	"grpc-tennis/models"
	"grpc-tennis/seeder"
	"grpc-tennis/user"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func run() error {
	// Create a new ServeMux for the HTTP server
	mux := runtime.NewServeMux()

	// Regđister your gRPC service with the ServeMux
	ctx := context.Background()
	endpoint := "localhost:9000"
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := location.RegisterLocationServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8080", mux)
}

func main() {
	config.Run()

	//connect to database
	db := database.ConnectDB()
	models.MigrateDB(db)
	seeder.Seed(db)

	//start gRPC server
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Failed to listen to port 9000 ", err)
	}

	s := location.Server{}
	s2 := user.Server{}

	grpcServer := grpc.NewServer()

	location.RegisterLocationServiceServer(grpcServer, &s)
	user.RegisterUserServiceServer(grpcServer, &s2)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("Failed to serve gRPC setver on port 9000", err)
		}
	}()

	if err := run(); err != nil {
		log.Fatal(err)
	}

}
