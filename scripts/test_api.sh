#!/bin/bash

# API Test Script for Supabase-E API
# Tests: Register, Login, Get User Profile, Create Product, List Products

# Set the API base URL
API_URL="http://127.0.0.1:8080/api/v1"
TOKEN=""
USER_ID=""

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print section headers
print_header() {
  echo -e "\n${BLUE}=== $1 ===${NC}\n"
}

# Function to check if the API is running
check_api() {
  print_header "Checking if API is running"
  
  response=$(curl -s -o /dev/null -w "%{http_code}" http://127.0.0.1:8080/health)
  
  if [ "$response" == "200" ]; then
    echo -e "${GREEN}API is running!${NC}"
  else
    echo -e "${RED}API is not running. Please start the API server first.${NC}"
    exit 1
  fi
}

# Function to register a new user
register_user() {
  print_header "Registering a new user"
  
  # Generate a random email to avoid conflicts
  local timestamp=$(date +%s)
  local email="test${timestamp}@example.com"
  
  echo "Using email: $email"
  
  response=$(curl -s -X POST "${API_URL}/auth/register" \
    -H "Content-Type: application/json" \
    -d '{
      "email": "'"$email"'",
      "password": "Password123",
      "first_name": "Test",
      "last_name": "User"
    }')
  
  echo "$response" | jq .
  
  # Check if registration was successful
  if echo "$response" | jq -e '.success' > /dev/null; then
    echo -e "${GREEN}User registration successful!${NC}"
    USER_ID=$(echo "$response" | jq -r '.data.id')
    echo "User ID: $USER_ID"
    return 0
  else
    echo -e "${RED}User registration failed.${NC}"
    return 1
  fi
}

# Function to login
login() {
  print_header "Logging in"
  
  # Use the email from the registration step
  local email=$(echo "$1" | jq -r '.data.email')
  
  response=$(curl -s -X POST "${API_URL}/auth/login" \
    -H "Content-Type: application/json" \
    -d '{
      "email": "'"$email"'",
      "password": "Password123"
    }')
  
  echo "$response" | jq .
  
  # Check if login was successful and extract token
  if echo "$response" | jq -e '.success' > /dev/null; then
    TOKEN=$(echo "$response" | jq -r '.data.token')
    echo -e "${GREEN}Login successful!${NC}"
    echo "Token: ${TOKEN:0:20}... (truncated)"
    return 0
  else
    echo -e "${RED}Login failed.${NC}"
    return 1
  fi
}

# Function to get user profile
get_profile() {
  print_header "Getting user profile"
  
  response=$(curl -s -X GET "${API_URL}/users/me" \
    -H "Authorization: Bearer $TOKEN")
  
  echo "$response" | jq .
  
  # Check if profile retrieval was successful
  if echo "$response" | jq -e '.success' > /dev/null; then
    echo -e "${GREEN}Profile retrieval successful!${NC}"
    return 0
  else
    echo -e "${RED}Profile retrieval failed.${NC}"
    return 1
  fi
}

# Function to create a product
create_product() {
  print_header "Creating a product"
  
  response=$(curl -s -X POST "${API_URL}/products" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $TOKEN" \
    -d '{
      "name": "Test Product",
      "description": "This is a test product created by the API test script",
      "price": 99.99,
      "category": "test",
      "image_url": "https://example.com/image.jpg"
    }')
  
  echo "$response" | jq .
  
  # Check if product creation was successful
  if echo "$response" | jq -e '.success' > /dev/null; then
    PRODUCT_ID=$(echo "$response" | jq -r '.data.id')
    echo -e "${GREEN}Product creation successful!${NC}"
    echo "Product ID: $PRODUCT_ID"
    return 0
  else
    echo -e "${RED}Product creation failed.${NC}"
    return 1
  fi
}

# Function to list products
list_products() {
  print_header "Listing products"
  
  response=$(curl -s -X GET "${API_URL}/products" \
    -H "Authorization: Bearer $TOKEN")
  
  echo "$response" | jq .
  
  # Check if product listing was successful
  if echo "$response" | jq -e '.success' > /dev/null; then
    echo -e "${GREEN}Product listing successful!${NC}"
    return 0
  else
    echo -e "${RED}Product listing failed.${NC}"
    return 1
  fi
}

# Main execution flow
main() {
  echo "Starting API tests..."
  
  # Check if jq is installed
  if ! command -v jq &> /dev/null; then
    echo -e "${RED}Error: jq is not installed. Please install it to run this script.${NC}"
    echo "On macOS: brew install jq"
    echo "On Ubuntu/Debian: apt-get install jq"
    exit 1
  fi
  
  # Check if the API is running
  # check_api
  
  # Register a new user
  # Generate a timestamp for unique email
  TIMESTAMP=$(date +%s)
  EMAIL="test${TIMESTAMP}@gmail.com"
  
  echo "Using email: $EMAIL"
  
  register_response=$(curl -s -X POST "${API_URL}/auth/register" \
    -H "Content-Type: application/json" \
    -d '{
      "email": "'"$EMAIL"'",
      "password": "Password123",
      "first_name": "Test",
      "last_name": "User"
    }')
  
  echo "$register_response" | jq .
  
  # Wait for email confirmation
  print_header "Waiting for email confirmation"
  echo "Waiting for 60 seconds to allow for email confirmation..."
  echo "(You can press Ctrl+C to skip waiting if you've already confirmed the email)"
  
  # Start a countdown timer
  for i in {60..1}; do
    echo -ne "$i seconds remaining...\r"
    sleep 1
  done
  echo -e "\nContinuing with login..."
  
  # Login with the registered user
  login "$register_response"
  
  # If login successful, continue with other tests
  if [ -n "$TOKEN" ]; then
    get_profile
    create_product
    list_products
    
    echo -e "\n${GREEN}All tests completed!${NC}"
  else
    echo -e "\n${RED}Tests stopped due to login failure.${NC}"
    exit 1
  fi
}

# Run the main function
main
