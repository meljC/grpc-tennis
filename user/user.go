package user

import (
	context "context"
	"fmt"
	"grpc-tennis/database"
	"grpc-tennis/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	UnimplementedUserServiceServer
}

func (s *Server) Create(ctx context.Context, request *CreateUserRequest) (*Response, error) {
	log.Printf("Received reequst to add a user: %s %s", request.GetFirstName(), request.GetEmail())

	var u User
	u.FirstName = request.GetFirstName()
	u.LastName = request.GetLastName()
	u.Email = request.GetEmail()
	u.Password = hashPassword(request.GetPassword())
	u.RoleId = request.GetRoleId()

	database.DB.Create(&u)

	return &Response{Message: "User added!"}, nil
}

func (s *Server) GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error) {
	log.Printf("Received reequst to list all users: ")

	var u = []*User{}
	database.DB.Find(&u)

	return &GetUsersResponse{Users: u}, nil
}

func (s *Server) Get(ctx context.Context, request *GetUserRequest) (*GetUserResponse, error) {
	log.Printf("Received reequst to GET a User by ID: ")

	var u *User
	id := request.GetId()
	fmt.Println("Get User with ID: ", id)

	database.DB.First(&u, id)
	fmt.Println(u)
	return &GetUserResponse{User: u}, nil
}

func (s *Server) Update(ctx context.Context, request *UpdateUserRequest) (*Response, error) {
	log.Printf("Received reequst to Update a User: ")

	var u *User
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

	return &Response{Message: "Upddated Users!"}, nil
}

func (s *Server) Delete(ctx context.Context, request *DeleteUserRequest) (*Response, error) {
	log.Printf("Received reequst to Delete a User: ")

	var u models.User
	id := request.GetId()

	database.DB.First(&u, id)
	fmt.Println(u)

	if u.ID == 0 {
		fmt.Println("User not found")
	}

	database.DB.Delete(&u, id)
	return &Response{Message: "User deleted"}, nil
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
