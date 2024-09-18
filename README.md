# F1 Results API

The F1 Results API is a lightweight and efficient service designed to provide access to Formula 1 race result data.  
This API is built with Go, utilizing the Chi router for efficient routing and includes CORS support for cross-origin requests.

## Features

- Fetch a list of all F1 tracks for a specific year
- Fetch information about a specific track
- Retrieve race results for a particular year
- Get race results for a specific track in a given year
- Comprehensive test coverage
- Swagger documentation

## API Endpoints

The API provides the following endpoints:

- `GET /getTracks/{year}`: Retrieve all tracks for a specific year
- `GET /getTracks/{year}/{trackName}`: Get information about a specific track in a given year
- `GET /getResults/{year}`: Retrieve all race results for a specific year
- `GET /getResult/{year}/{trackId}`: Get race results for a specific track in a given year
- `GET /swagger/`: Open documentation made with Swagger

For detailed information about request parameters, responses, and possible error codes, please refer to the Swagger documentation.

## Prerequisites

- PostgreSQL database
- Make (for using Makefile)

## Installation

### Local

1. Clone the repository:
   ```
   git clone https://github.com/kamildemocko/F1ResultsAPI.git
   ```

2. Navigate to the project directory:
   ```
   cd F1ResultsAPI
   ```

3. Install dependencies:
   ```
   go mod tidy
   ```

4. Set up your environment variables in a `.env` file:
   ```
   DSN=your_postgresql_connection_string
   ```

5. Build the application:
   ```
   make build
   ```

6. Run the application:
   ```
   make run
   ```

### Docker

1. Build the Docker image:
   ```
   make dbuild
   ```

2. Run the Docker container:
   ```
   make dstart
   ```

3. Stop the Docker container:
   ```
   make dstop
   ```


The API will be available at `http://localhost/f1/api`.

## Documentation

This API is documented using Swagger. To generate or update the Swagger documentation, run:

```
make swag
```