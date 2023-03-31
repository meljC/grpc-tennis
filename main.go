package main

import (
	"grpc-tennis/config"
	"grpc-tennis/database"
	"grpc-tennis/location"
	"grpc-tennis/models"
	"grpc-tennis/seeder"
	"grpc-tennis/user"
	"log"
	"net"

	"google.golang.org/grpc"
)

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
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve gRPC setver on port 9000", err)
	}
}
