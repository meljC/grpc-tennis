package user

import (
	context "context"
	"grpc-tennis/database"
	"grpc-tennis/models"
	"log"
)

type Server struct {
	UnimplementedUserServiceServer
}

func (s *Server) Create(ctx context.Context, request *CreateUserRequest) (*Response, error) {
	log.Printf("Received reequst to add a user: %s %s", request.GetFirstName(), request.GetEmail())

	var u models.User
	u.FirstName = request.GetFirstName()
	u.Email = request.GetEmail()
	u.Password = request.GetPassword()
	u.RoleID = uint(request.GetRoleId())
	database.DB.Create(&u)

	return &Response{Message: "User added!"}, nil
}

func GetUserByEmail(email string) *models.User {
	var u models.User
	err := database.DB.Where("email = ?", email).First(&u).Error
	if err != nil {
		log.Fatal("user not found")
	}

	return &u

}
