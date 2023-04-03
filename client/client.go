package main

import (
	"context"
	"grpc-tennis/location"
	"grpc-tennis/user"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateRequest(c location.LocationServiceClient) {
	request := location.CreateLocationRequest{
		CityId:    3,
		Latitude:  22.3,
		Longitude: 4.4,
		Address:   "Kikiceva",
	}

	response, err := c.Create(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling Create: %s", err)
	}

	log.Printf("Response from Server: %s", response.GetMessage())
}

func UpdateRequest(c location.LocationServiceClient) {
	request := location.UpdateLocationRequest{
		Id:        101,
		Latitude:  22.5,
		Longitude: 4.7,
		Address:   "Pantovčak",
	}

	response, err := c.Update(context.Background(), &request)
	if err != nil {
		log.Fatalf("Error when calling Update: %s", err)
	}

	log.Printf("Response from Server: %s", response.GetMessage())
}

func GetAll(c location.LocationServiceClient) {
	requestAll := location.GetLocationsRequest{}
	responseAll, err := c.GetLocations(context.Background(), &requestAll)
	if err != nil {
		log.Fatalf("Error when calling GerLocations: %s", err)
	}

	log.Printf("Response from Server: %s", responseAll.GetLocations())
}

func GetIdRequest(c location.LocationServiceClient) {
	requestID := location.GetLocationRequest{Id: 101}
	responseID, err := c.Get(context.Background(), &requestID)
	if err != nil {
		log.Fatalf("Error when calling Get: %s", err)
	}

	log.Printf("Response from Server: %s", responseID.GetLocation())
}

func DeleteRequest(c location.LocationServiceClient) {
	requestID := location.DeleteLocationRequest{Id: 101}
	responseID, err := c.Delete(context.Background(), &requestID)
	if err != nil {
		log.Fatalf("Error when calling Delete: %s", err)
	}

	log.Printf("Response from Server: %s", responseID.GetMessage())
}

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := location.NewLocationServiceClient(conn)

	// C R U D
	CreateRequest(c)
	GetAll(c)

	GetIdRequest(c)
	UpdateRequest(c)
	GetIdRequest(c)
	DeleteRequest(c)
	GetIdRequest(c)

	GetAll(c)

	c2 := user.NewUserServiceClient(conn)

	requestUser := user.CreateUserRequest{
		FirstName: "Zoki",
		Email:     "plenki@zoki.hr",
		RoleId:    1,
	}

	responseUser, err := c2.Create(context.Background(), &requestUser)
	if err != nil {
		log.Fatalf("Error when calling Create: %s", err)
	}

	log.Printf("Response from Server: %s", responseUser.GetMessage())
}
