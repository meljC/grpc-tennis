package services

import (
	context "context"
	"fmt"
	"grpc-tennis/database"
	"grpc-tennis/gen"
	"grpc-tennis/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserServer struct {
	gen.UnimplementedUserServiceServer
}

func (s *UserServer) Create(ctx context.Context, request *gen.CreateUserRequest) (*gen.Response, error) {
	log.Printf("Received reequst to add a user: %s %s", request.GetFirstName(), request.GetEmail())

	var u gen.User
	u.FirstName = request.GetFirstName()
	u.LastName = request.GetLastName()
	u.Email = request.GetEmail()
	u.Password = hashPassword(request.GetPassword())
	u.RoleId = request.GetRoleId()

	database.DB.Create(&u)

	return &gen.Response{Message: "User added!"}, nil
}

func (s *UserServer) GetUsers(ctx context.Context, request *gen.GetUsersRequest) (*gen.GetUsersResponse, error) {
	log.Printf("Received reequst to list all users: ")

	var u = []*gen.User{}
	database.DB.Find(&u)

	return &gen.GetUsersResponse{Users: u}, nil
}

func (s *UserServer) Get(ctx context.Context, request *gen.GetUserRequest) (*gen.GetUserResponse, error) {
	log.Printf("Received reequst to GET a User by ID: ")

	var u *gen.User
	id := request.GetId()
	fmt.Println("Get User with ID: ", id)

	database.DB.First(&u, id)
	fmt.Println(u)
	return &gen.GetUserResponse{User: u}, nil
}

func (s *UserServer) Update(ctx context.Context, request *gen.UpdateUserRequest) (*gen.Response, error) {
	log.Printf("Received reequst to Update a User: ")

	var u *gen.User
	id := request.GetId()

	database.DB.First(&u, id)

	if id == 0 {
		fmt.Println("User not found")
	}

	u.FirstName = request.GetFirstName()
	u.LastName = request.GetLastName()
	u.Email = request.GetEmail()
	u.Password = hashPassword(request.GetPassword())
	u.RoleId = request.GetRoleId()

	database.DB.Save(&u)

	return &gen.Response{Message: "Upddated Users!"}, nil
}

func (s *UserServer) Delete(ctx context.Context, request *gen.DeleteUserRequest) (*gen.Response, error) {
	log.Printf("Received reequst to Delete a User: ")

	var u models.User
	id := request.GetId()

	database.DB.First(&u, id)
	fmt.Println(u)

	if u.ID == 0 {
		fmt.Println("User not found")
	}

	database.DB.Delete(&u, id)
	return &gen.Response{Message: "User deleted"}, nil
}

func GetUserByEmail(email string) *models.User {
	var u models.User
	err := database.DB.Where("email = ?", email).First(&u).Error
	if err != nil {
		log.Fatal("user not found")
	}

	return &u

}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hashedPassword)
}
