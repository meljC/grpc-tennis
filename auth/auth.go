package auth

import (
	"context"
	"grpc-tennis/user"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	UnimplementedAuthServiceServer
}

func (s *Server) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	user := user.GetUserByEmail(req.Email)

	if user == nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid email or password")
	}

	if user.Password != req.Password {
		return nil, status.Errorf(codes.Unauthenticated, "invalid email or password")
	}

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

	return &LoginResponse{Token: tokenString}, nil
}
