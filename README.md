# Drone Task
This is a API project developed using Golang that allows users to list & register drones

## Requirements
- Go version go1.21.0 

## Installation

### Using local environment
1.Clone the repository:
    `git clone https://github.com/MahmoudAtif/Drone-Task.git`

2.Install the project dependencies:
    `go get`

## Configration
- Edit the .env file with your database connection details. Add your database connection string as shown below:
    `export DB_CONNECTION_STRING="host=test user=test port=test password=test dbname=test sslmode=disable"`

## Usage
1.Start the development server
-   `go run main.go`

## Api Documentation

### Get All Drones
- Request Example:
    ```
    curl -X GET http://127.0.0.1:9999/api/drones/
    ```
- Response Example:
    ```
    [
        {
            "serial_number": "SN123",
            "model": "Lightweight",
            "weight": 100,
            "battery_capacity": 5000,
            "state": "IDLE"
        },
        {
            "serial_number": "SN456",
            "model": "Heavyweight",
            "weight_limit": 200,
            "battery_capacity": 6000,
            "state": "LOADING"
        }
    ]
    ```
### Register Drones
- Request Example:
    ```
    curl -X POST -H "Content-Type: application/json" -d '[
       {
            "serial_number": "SN123",
            "model": "Lightweight",
            "weight": 100,
            "battery_capacity": 5000,
            "state": "IDLE"
        },
        {
            "serial_number": "SN456",
            "model": "Heavyweight",
            "weight_limit": 200,
            "battery_capacity": 6000,
            "state": "LOADING"
        }
    ]' http://localhost:9999/api/drones/
    ```
- Response Example:
    ```
    {
        "created_drones": [
            {   
                "id":1,
                "serial_number": "SN123",
                "model": "Lightweight",
                "weight": 100,
                "battery_capacity": 5000,
                "state": "IDLE",
                "created_at": "",
                "updated_at": ""
            },
            {
                "id":2,
                "serial_number": "SN456",
                "model": "Heavyweight",
                "weight_limit": 200,
                "battery_capacity": 6000,
                "state": "LOADING",
                "created_at": "",
                "updated_at": ""
            }
        ]
    }
    ```