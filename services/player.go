package services

import (
	context "context"
	"fmt"
	"grpc-tennis/database"
	"grpc-tennis/gen"
	"grpc-tennis/models"
	"log"

	"github.com/PeteProgrammer/go-automapper"
)

type PlayerServer struct {
	gen.UnimplementedPlayerServiceServer
}

var players = []*gen.Player{}

func (s *PlayerServer) Create(ctx context.Context, request *gen.CreatePlayerRequest) (*gen.Player, error) {

	var l models.Player
	automapper.MapLoose(request, &l)

	database.DB.Create(&l)

	var response = &gen.Player{}
	automapper.MapLoose(&l, response)

	return response, nil
}

func (s *PlayerServer) GetPlayers(ctx context.Context, request *gen.GetPlayersRequest) (*gen.GetPlayersResponse, error) {
	log.Printf("Received reequst to list all Players: ")

	database.DB.Find(&players)

	return &gen.GetPlayersResponse{Players: players}, nil
}

func (s *PlayerServer) Get(ctx context.Context, request *gen.GetPlayerRequest) (*gen.GetPlayerResponse, error) {
	log.Printf("Received reequst to GET a Player by ID: ")

	var l *gen.Player
	id := request.GetId()
	fmt.Println("Get Player with ID: ", id)

	database.DB.First(&l, id)
	fmt.Println(l)
	return &gen.GetPlayerResponse{Player: l}, nil
}

func (s *PlayerServer) Update(ctx context.Context, request *gen.UpdatePlayerRequest) (*gen.Player, error) {
	log.Printf("Received reequst to Update a Player: ")

	var l *gen.Player
	id := request.GetId()

	database.DB.First(&l, id)

	if id == 0 {
		fmt.Println("Player not found")
	}

	automapper.MapLoose(request, &l)

	database.DB.Save(&l)

	var response = &gen.Player{}
	automapper.MapLoose(&l, response)

	return &gen.Player{}, nil
}

func (s *PlayerServer) Delete(ctx context.Context, request *gen.DeletePlayerRequest) (*gen.Player, error) {
	log.Printf("Received reequst to Delete a Player: ")

	var l models.Player
	id := request.GetId()

	database.DB.First(&l, id)
	fmt.Println(l)

	if l.ID == 0 {
		fmt.Println("Player not found")
	}

	database.DB.Delete(&l, id)
	return &gen.Player{}, nil
}
