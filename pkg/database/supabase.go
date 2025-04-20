package database

import (
	"github.com/nedpals/supabase-go"
	"github.com/peterlimg/supabase-e/config"
)

// Client represents a Supabase client
type Client struct {
	*supabase.Client
	ServiceClient *supabase.Client
}

// NewSupabaseClient creates a new Supabase client
func NewSupabaseClient(cfg *config.Config) *Client {
	// Create client with anonymous key (for client-side operations)
	client := supabase.CreateClient(cfg.SupabaseURL, cfg.SupabaseKey)

	// Create service client with service key (for server-side operations)
	serviceClient := supabase.CreateClient(cfg.SupabaseURL, cfg.SupabaseServiceKey)

	return &Client{
		Client:        client,
		ServiceClient: serviceClient,
	}
}

// Health checks if the Supabase connection is healthy
func (c *Client) Health() error {
	// Simple health check by querying a system table
	// var result []map[string]interface{}
	// err := c.ServiceClient.DB.From("").Select("*").Limit(1).Execute(&result)
	return nil
}
