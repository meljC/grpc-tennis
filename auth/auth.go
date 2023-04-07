package auth

import (
	"context"
	"grpc-tennis/user"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServiceServer struct{}

func (s *authServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	// Check if the email and password match with the user's records
	user := user.GetUserByEmail(req.Email)

	//if !checkPasswordHash(req.Password, user.Password) {
	//	return nil, status.Errorf(codes.Unauthenticated, "invalid email or password")
	//}

	// Create a new JWT token with the user's email as the subject
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &jwt.StandardClaims{
		Subject:   user.Email,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte("my-secret-key")
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate JWT token")
	}

	// Return the JWT token as a response
	return &pb.LoginResponse{
		Token: tokenString,
	}, nil
}
