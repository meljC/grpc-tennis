package main

import (
	"context"
	"grpc-tennis/gen"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateRequest(c gen.LocationServiceClient) {
	request := gen.CreateLocationRequest{
		CityId:    3,
		Latitude:  22.3,
		Longitude: 4.4,
		Address:   "Kikiceva",
	}

	response, err := c.Create(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling Create: %s", err)
	}

	log.Printf("Response from Server: Created Location %d", response.GetId())
}

func UpdateRequest(c gen.LocationServiceClient) {
	request := gen.UpdateLocationRequest{
		Id:        101,
		Latitude:  22.5,
		Longitude: 4.7,
		Address:   "Pantovčak",
	}

	response, err := c.Update(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling Update: %s", err)
	}

	log.Printf("Response from Server: Updated %d", response.GetId())
}

func GetAll(c gen.LocationServiceClient) {
	requestAll := gen.GetLocationsRequest{}
	responseAll, err := c.GetLocations(context.Background(), &requestAll)
	if err != nil {
		log.Fatalf("Error when calling GerLocations: %s", err)
	}

	log.Printf("Response from Server: %s", responseAll.GetLocations())
}

func GetIdRequest(c gen.LocationServiceClient) {
	requestID := gen.GetLocationRequest{Id: 101}
	responseID, err := c.Get(context.Background(), &requestID)
	if err != nil {
		log.Fatalf("Error when calling Get: %s", err)
	}

	log.Printf("Response from Server: %s", responseID.GetLocation())
}

func DeleteRequest(c gen.LocationServiceClient) {
	requestID := gen.DeleteLocationRequest{Id: 101}
	responseID, err := c.Delete(context.Background(), &requestID)
	if err != nil {
		log.Fatalf("Error when calling Delete: %s", err)
	}

	log.Printf("Response from Server: Deleted %d", responseID.GetId())
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := gen.NewLocationServiceClient(conn)

	// C R U D
	CreateRequest(c)
	GetAll(c)

	GetIdRequest(c)
	UpdateRequest(c)
	GetIdRequest(c)
	DeleteRequest(c)
	GetIdRequest(c)

	GetAll(c)

	c2 := gen.NewUserServiceClient(conn)

	requestUser := gen.CreateUserRequest{
		FirstName: "Zoki",
		Email:     "plenki@zoki.hr",
		Password:  "pass123",
		RoleId:    1,
	}

	responseUser, err := c2.Create(context.Background(), &requestUser)
	if err != nil {
		log.Fatalf("Error when calling Create: %s", err)
	}

	log.Printf("Response from Server: %s", responseUser.GetMessage())
}
