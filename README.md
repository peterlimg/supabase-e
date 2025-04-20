# Supabase-E: Go API Service with Supabase Integration

A modern Go API service that integrates with Supabase for authentication and database operations.

## Features

- RESTful API using Gin framework
- JWT-based authentication
- Supabase integration for database and auth
- Structured project layout following Go best practices
- Middleware for logging, authentication, and authorization
- Graceful server shutdown
- Configuration management
- Error handling

## Project Structure

```
├── cmd/
│   └── api/              # Application entrypoints
│       └── main.go       # Main application
├── config/               # Configuration handling
│   └── config.go
├── internal/             # Private application code
│   ├── handlers/         # HTTP handlers
│   ├── middleware/       # HTTP middleware
│   ├── models/           # Data models
│   ├── repository/       # Database repositories
│   └── services/         # Business logic
├── pkg/                  # Public libraries
│   ├── database/         # Database client
│   ├── logger/           # Logging utilities
│   └── utils/            # Utility functions
├── .env                  # Environment variables
├── go.mod                # Go module file
├── go.sum                # Go module checksum
└── README.md             # Project documentation
```

## Prerequisites

- Go 1.22 or higher
- Supabase account and project

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/peterlimg/supabase-e.git
   cd supabase-e
   ```

2. Set up your Supabase project:
   - Create a new Supabase project
   - Create the necessary tables (users, products)
   - Get your Supabase URL and API keys

3. Configure environment variables:
   - Copy `.env.example` to `.env`
   - Update the values with your Supabase credentials

4. Install dependencies:
   ```
   go mod download
   ```

5. Run the application:
   ```
   go run cmd/api/main.go
   ```

## API Endpoints

### Authentication

- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login and get JWT token

### User Management

- `GET /api/v1/users/me` - Get current user profile
- `PUT /api/v1/users/me` - Update current user profile

### Products

- `GET /api/v1/products` - List all products
- `POST /api/v1/products` - Create a new product
- `GET /api/v1/products/:id` - Get a product by ID
- `GET /api/v1/products/:id/with-user` - Get a product with creator info
- `PUT /api/v1/products/:id` - Update a product
- `DELETE /api/v1/products/:id` - Delete a product

### Health Check

- `GET /health` - Check API health

## Database Schema

### Users Table

```sql
CREATE TABLE users (
  id UUID PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  role TEXT NOT NULL DEFAULT 'user',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

### Products Table

```sql
CREATE TABLE products (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  price DECIMAL NOT NULL,
  category TEXT NOT NULL,
  image_url TEXT,
  created_by UUID REFERENCES users(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
```

## License

MIT
