package services

import (
	"context"
	"grpc-tennis/gen"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	gen.UnimplementedAuthServiceServer
}

func (s *AuthServer) Login(ctx context.Context, req *gen.LoginRequest) (*gen.LoginResponse, error) {
	user := GetUserByEmail(req.Email)

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

	return &gen.LoginResponse{Token: tokenString}, nil
}

func (s *AuthServer) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
