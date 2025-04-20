package services

import (
	"context"
	"fmt"

	"github.com/nedpals/supabase-go"
	"github.com/peterlimg/supabase-e/config"
	"github.com/peterlimg/supabase-e/internal/models"
	"github.com/peterlimg/supabase-e/internal/repository"
	"github.com/peterlimg/supabase-e/pkg/database"
	"github.com/peterlimg/supabase-e/pkg/utils"
)

// AuthService handles authentication operations
type AuthService struct {
	userRepo *repository.UserRepository
	db       *database.Client
	config   *config.Config
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo *repository.UserRepository, db *database.Client, config *config.Config) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		db:       db,
		config:   config,
	}
}

// Register registers a new user
func (s *AuthService) Register(req models.CreateUserRequest) (*models.User, error) {
	// Create the user
	user, err := s.userRepo.Create(req)
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %w", err)
	}

	return user, nil
}

// Login authenticates a user and returns a JWT token
func (s *AuthService) Login(req models.LoginRequest) (*models.LoginResponse, error) {
	// Authenticate with Supabase Auth
	creds := supabase.UserCredentials{
		Email:    req.Email,
		Password: req.Password,
	}

	authResp, err := s.db.Client.Auth.SignIn(context.Background(), creds)
	if err != nil {
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	// Get the user from the database
	user, err := s.userRepo.GetByID(authResp.User.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Generate a JWT token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role, s.config.JWTSecret, s.config.JWTExpiry)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	// Return the user and token
	return &models.LoginResponse{
		User:  *user,
		Token: token,
	}, nil
}

// GetUserByID gets a user by ID
func (s *AuthService) GetUserByID(id string) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

// UpdateUser updates a user
func (s *AuthService) UpdateUser(id string, req models.UpdateUserRequest) (*models.User, error) {
	return s.userRepo.Update(id, req)
}
