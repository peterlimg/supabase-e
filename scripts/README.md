# API Test Scripts

This directory contains scripts for testing the Supabase-E API.

## Test API Script

The `test_api.sh` script tests the core functionality of the API, including:

1. User registration
2. User login
3. User profile retrieval
4. Product creation
5. Product listing

### Prerequisites

- The API server must be running on `localhost:8080`
- `jq` must be installed for JSON processing
  - On macOS: `brew install jq`
  - On Ubuntu/Debian: `apt-get install jq`

### Usage

To run the test script:

```bash
./test_api.sh
```

### What the Script Does

1. Checks if the API is running by making a request to the `/health` endpoint
2. Registers a new user with random email to avoid conflicts
3. Logs in with the newly created user credentials
4. Retrieves the user profile using the authentication token
5. Creates a test product
6. Lists all products

### Output

The script provides colored output to indicate success or failure of each operation:
- 🟢 Green: Success
- 🔴 Red: Failure
- 🔵 Blue: Section headers

Each API response is displayed in JSON format for inspection.

### Troubleshooting

If the script fails:

1. Ensure the API server is running on `localhost:8080`
2. Check that your database is properly configured
3. Verify that the API endpoints match those in the script
4. Make sure `jq` is installed

### Extending the Script

You can extend this script to test additional API endpoints by adding new functions following the same pattern as the existing ones.
