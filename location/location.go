package location

import (
	context "context"
	"fmt"
	"grpc-tennis/database"
	"grpc-tennis/models"
	"log"
)

type Server struct {
	UnimplementedLocationServiceServer
}

func (s *Server) Create(ctx context.Context, request *CreateLocationRequest) (*Response, error) {
	log.Printf("Received reequst to add a Location: %d %s", request.GetCityId(), request.GetAddress())

	var l models.Location
	l.CityID = uint(request.GetCityId())
	l.Latitude = request.GetLatitude()
	l.Longitude = request.GetLongitude()
	l.Address = request.GetAddress()
	database.DB.Create(&l)

	return &Response{Message: "Location added!"}, nil
}

func (s *Server) GetLocations(ctx context.Context, request *GetLocationsRequest) (*Response, error) {
	log.Printf("Received reequst to list all Locations: ")

	var l []models.Location
	database.DB.Find(&l)

	for _, loc := range l {
		fmt.Println(loc)
	}
	return &Response{Message: "List all Locations!"}, nil
}

func (s *Server) Get(ctx context.Context, request *GetLocationRequest) (*Response, error) {
	log.Printf("Received reequst to GET a Location by ID: ")

	var l models.Location
	id := request.GetId()

	database.DB.First(&l, id)
	fmt.Println(l)
	return &Response{Message: "Location By ID"}, nil
}

func (s *Server) Update(ctx context.Context, request *UpdateLocationRequest) (*Response, error) {
	log.Printf("Received reequst to Update a Location: ")

	var l models.Location
	id := request.GetId()

	database.DB.First(&l, id)
	fmt.Println(l)

	if l.ID == 0 {
		fmt.Println("Location not found")
	}

	database.DB.Save(&l)

	return &Response{Message: "Upddated Locations!"}, nil
}

func (s *Server) Delete(ctx context.Context, request *DeleteLocationRequest) (*Response, error) {
	log.Printf("Received reequst to Delete a Location: ")

	var l models.Location
	id := request.GetId()

	database.DB.First(&l, id)
	fmt.Println(l)

	if l.ID == 0 {
		fmt.Println("Location not found")
	}

	database.DB.Delete(&l, id)
	return &Response{Message: "Location deleted"}, nil
}
