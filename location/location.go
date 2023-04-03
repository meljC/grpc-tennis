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

var locations = []*Location{}

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

func (s *Server) GetLocations(ctx context.Context, request *GetLocationsRequest) (*GetLocationsResponse, error) {
	log.Printf("Received reequst to list all Locations: ")

	database.DB.Find(&locations)

	return &GetLocationsResponse{Locations: locations}, nil
}

func (s *Server) Get(ctx context.Context, request *GetLocationRequest) (*GetLocationResponse, error) {
	log.Printf("Received reequst to GET a Location by ID: ")

	var l *Location
	id := request.GetId()
	fmt.Println("Get Location with ID: ", id)

	database.DB.First(&l, id)
	fmt.Println(l)
	return &GetLocationResponse{Location: l}, nil
}

func (s *Server) Update(ctx context.Context, request *UpdateLocationRequest) (*Response, error) {
	log.Printf("Received reequst to Update a Location: ")

	var l *Location
	id := request.GetId()

	database.DB.First(&l, id)

	if id == 0 {
		fmt.Println("Location not found")
	}

	//l.CityId = request.GetCityId()
	l.Latitude = request.GetLatitude()
	l.Longitude = request.GetLongitude()
	l.Address = request.GetAddress()

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
