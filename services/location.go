package services

import (
	context "context"
	"fmt"
	"grpc-tennis/database"
	"grpc-tennis/gen"
	"grpc-tennis/models"
	"log"
)

type LocationServer struct {
	gen.UnimplementedLocationServiceServer
}

var locations = []*gen.Location{}

func (s *LocationServer) Create(ctx context.Context, request *gen.CreateLocationRequest) (*gen.Location, error) {
	log.Printf("Received reequst to add a Location: %d %s", request.GetCityId(), request.GetAddress())

	var l models.Location
	l.CityID = uint(request.GetCityId())
	l.Latitude = request.GetLatitude()
	l.Longitude = request.GetLongitude()
	l.Address = request.GetAddress()
	database.DB.Create(&l)

	return &gen.Location{}, nil
}

func (s *LocationServer) GetLocations(ctx context.Context, request *gen.GetLocationsRequest) (*gen.GetLocationsResponse, error) {
	log.Printf("Received reequst to list all Locations: ")

	database.DB.Find(&locations)

	return &gen.GetLocationsResponse{Locations: locations}, nil
}

func (s *LocationServer) Get(ctx context.Context, request *gen.GetLocationRequest) (*gen.GetLocationResponse, error) {
	log.Printf("Received reequst to GET a Location by ID: ")

	var l *gen.Location
	id := request.GetId()
	fmt.Println("Get Location with ID: ", id)

	database.DB.First(&l, id)
	fmt.Println(l)
	return &gen.GetLocationResponse{Location: l}, nil
}

func (s *LocationServer) Update(ctx context.Context, request *gen.UpdateLocationRequest) (*gen.Location, error) {
	log.Printf("Received reequst to Update a Location: ")

	var l *gen.Location
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

	return &gen.Location{}, nil
}

func (s *LocationServer) Delete(ctx context.Context, request *gen.DeleteLocationRequest) (*gen.Location, error) {
	log.Printf("Received reequst to Delete a Location: ")

	var l models.Location
	id := request.GetId()

	database.DB.First(&l, id)
	fmt.Println(l)

	if l.ID == 0 {
		fmt.Println("Location not found")
	}

	database.DB.Delete(&l, id)
	return &gen.Location{}, nil
}
