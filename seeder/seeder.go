package seeder

import (
	"fmt"
	"grpc-tennis/config"
	"grpc-tennis/models"

	"log"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	count := config.App.Config.SeedUserCount
	fmt.Println("\n Count: \n", count)
	batch := 10

	var roles = []models.Role{{Name: "Admin"}, {Name: "User"}}
	db.Create(&roles)

	if count < batch {
		log.Fatal("Seed count must be > ", batch)
	}
	for j := 0; j < count/batch; j++ {
		cities := make([]models.City, batch)
		users := make([]models.User, batch)
		players := make([]models.Player, batch)
		locations := make([]models.Location, batch)
		matches := make([]models.Match, batch)
		results := make([]models.Result, batch)
		tournaments := make([]models.Tournament, batch)

		for i := 0; i < batch; i++ {
			gofakeit.Struct(&cities[i])
			gofakeit.Struct(&users[i])
			gofakeit.Struct(&players[i])
			gofakeit.Struct(&locations[i])
			gofakeit.Struct(&matches[i])
			gofakeit.Struct(&tournaments[i])
			gofakeit.Struct(&results[i])

			cities[i].ID = 0
			users[i].ID = 0
			players[i].ID = 0
			locations[i].ID = 0
			matches[i].ID = 0
			tournaments[i].ID = 0
			results[i].ID = 0

			users[i].Role = models.Role{}
			locations[i].City = models.City{}
			matches[i].Location = models.Location{}
			results[i].Player = models.Player{}
			results[i].Match = models.Match{}
			results[i].User = models.User{}
		}

		db.Create(&cities)
		db.Create(&users)
		db.Create(&players)
		db.Create(&locations)
		db.Create(&matches)
		db.Create(&results)
		db.Create(&tournaments)

	}

}
