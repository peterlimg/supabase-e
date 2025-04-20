package repository

import (
	"context"
	"fmt"

	"github.com/nedpals/supabase-go"
	"github.com/peterlimg/supabase-e/internal/models"
	"github.com/peterlimg/supabase-e/pkg/database"
)

// UserRepository handles user data operations
type UserRepository struct {
	db *database.Client
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *database.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create creates a new user in Supabase Auth and database
func (r *UserRepository) Create(user models.CreateUserRequest) (*models.User, error) {
	// First, create the user in Supabase Auth
	creds := supabase.UserCredentials{
		Email:    user.Email,
		Password: user.Password,
	}
	
	authResp, err := r.db.ServiceClient.Auth.SignUp(context.Background(), creds)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in auth: %w", err)
	}

	// Create a new user model
	newUser := models.NewUser(user.Email, user.FirstName, user.LastName)
	newUser.ID = authResp.ID // Use ID from the auth response

	// Insert the user into the users table
	var result []models.User
	err = r.db.ServiceClient.DB.From("users").Insert(newUser).Execute(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in database: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no user returned after insert")
	}

	return &result[0], nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(id string) (*models.User, error) {
	var users []models.User
	err := r.db.ServiceClient.DB.From("users").Select("*").Eq("id", id).Execute(&users)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &users[0], nil
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var users []models.User
	err := r.db.ServiceClient.DB.From("users").Select("*").Eq("email", email).Execute(&users)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &users[0], nil
}

// Update updates a user
func (r *UserRepository) Update(id string, user models.UpdateUserRequest) (*models.User, error) {
	var result []models.User
	err := r.db.ServiceClient.DB.From("users").Update(user).Eq("id", id).Execute(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &result[0], nil
}

// Delete deletes a user
func (r *UserRepository) Delete(id string) error {
	// Delete from the database
	err := r.db.ServiceClient.DB.From("users").Delete().Eq("id", id).Execute(nil)
	if err != nil {
		return fmt.Errorf("failed to delete user from database: %w", err)
	}

	// Note: Deleting from Auth would typically be handled by a trigger in Supabase

	return nil
}

// List lists all users with pagination
func (r *UserRepository) List(page, pageSize int) ([]models.User, error) {
	var users []models.User
	
	// Calculate offset
	offset := (page - 1) * pageSize
	limit := pageSize
	
	// Use the pagination parameters
	query := r.db.ServiceClient.DB.From("users").Select("*")
	
	// Add limit
	query = query.Limit(limit)
	
	// Execute query
	err := query.Execute(&users)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	
	// Apply offset manually if needed (for older versions of the client)
	if len(users) > offset {
		users = users[offset:]
	} else {
		users = []models.User{}
	}
	
	return users, nil
}
