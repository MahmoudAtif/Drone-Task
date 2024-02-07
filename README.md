# Drone Task
This is a API project developed using Golang that allows users to list & register drones

## Requirements
- Go version go1.21.0 

## Installation

### Using local environment
1.Clone the repository:
    `git clone https://github.com/MahmoudAtif/Drone-Task.git`

2.move to project directory:
    `cd Drone-Task`

3.Install the project dependencies:
    `go get`

## Configration
1. Edit the emaples.env file with your database connection details. Add your database connection string as shown below:
    `export DB_CONNECTION_STRING="host=test user=test port=test password=test dbname=test sslmode=disable"`

2. Run this command after modifying exmaple.env file
    `source example.env`
    
## Usage
1.Start the development server
-   `go run main.go`

## Running Tests
To run tests for the entire project, use the following command:
-   ```go test ./...```
### Additional Options
1.If you want more detailed output, you can use the -v (verbose) flag:
-   ```go test -v ./...```

2.you can run specific tests by providing the names of the test functions as arguments:
-   ```go test -v -run TestYourFunctionName ./...```

## Api Documentation
To run swagger documentation use the following command
-   ```swagger serve swagger.yaml```

### Get All Drones
- Request Example:
    ```
    curl -X GET http://127.0.0.1:9999/api/drone/
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
            "weight": 200,
            "battery_capacity": 60,
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
            "battery_capacity": 50,
            "state": "IDLE"
        },
        {
            "serial_number": "SN456",
            "model": "Heavyweight",
            "weight": 200,
            "battery_capacity": 60,
            "state": "LOADING"
        }
    ]' http://localhost:9999/api/drone/
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
                "battery_capacity": 50,
                "state": "IDLE",
                "created_at": "",
                "updated_at": ""
            },
            {
                "id":2,
                "serial_number": "SN456",
                "model": "Heavyweight",
                "weight": 200,
                "battery_capacity": 60,
                "state": "LOADING",
                "created_at": "",
                "updated_at": ""
            }
        ]
    }
    ```

### Register Medication
- Request Example:
    ```
    curl -X POST -H "Content-Type: application/json" -d '[
       {
            "name": "TestName",
            "code": "XXX123",
            "weight": 100,
            "image": "example.png"
        },
        {
            "name": "TestName2",
            "code": "XXX1234",
            "weight": 100,
            "image": "example2.jpeg"
        }
    ]' http://localhost:9999/api/medication/
    ```
- Response Example:
    ```
    {
        "created_medications": [
            {   
                "id":1,
                "name": "TestName",
                "code": "XXX123",
                "weight": 100,
                "image": "example.png"
                "created_at": "",
                "updated_at": ""
            },
            {
                "id":2,
                "name": "TestName2",
                "code": "XXX1234",
                "weight": 100,
                "image": "example2.jpeg"
                "created_at": "",
                "updated_at": ""
            }
        ]
    }
    ```
