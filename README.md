# Simple Web API

This is a simple web API that performs basic CRUD
using Go and Mysql

### How to run the application

- Clone the repository
- Navigate to the project directory: `cd simpleAPI`
- Create an `.env` file in the root directory of the project.
- Populate the `.env` file with the required environment variables.<br>
  Here's an example: <br>
  `PORT=:8080`<br>
  `DB_HOST={hostname}`<br>
  `DB_PORT={port}`<br>
  `DB_USERNAME={username}`<br>
  `DB_PASSWORD={password}`<br>
  `DB_NAME={simpleapidb}`<br>
  `DB_NETWORK=tcp`

#### Running using Docker

- run `docker compose up`
- to stop the application run `docker compose down`

#### Running Locally

- run `go build -o {binaryName} .\cmd\app\main.go` to build the binary
- run `.\{binary name}` to run the binary (eg .\simpleAPI)

### Testing

- run `go test -v ./...` to run the unit tests
- run `go test -v ./... -coverprofile=c.out` to create the cover profile
- run `go tool cover -html c` to see the
  coverage of the unit tests

## Endpoints

### Create Student

This Endpoint Creates a Student

#### Request

`curl --location 'http://localhost:8001/student/' \
--header 'Content-Type: application/json' \
--data '{
"firstname":"Daniel",
"lastname":"Riccardo",
"year":3
}'`

#### Response

    {
      "status": "Success",
      "data": {
          "id": 8,
          "firstname": "Daniel",
          "lastname": "Riccardo",
          "year": 3
      },
      "message": "Student Created Successfully"
    }

### Get All Students

This Endpoint Returns all the Students

#### Request

`curl --location 'http://localhost:8001/student/'`

#### Response

    {
      "status": "Success",
      "data": [
          {
              "id": 1,
              "firstname": "Charles",
              "lastname": "Leclerc",
              "year": 3
          },
          {
              "id": 2,
              "firstname": "Carlos",
              "lastname": "Sainz",
              "year": 1
          },
          {
              "id": 3,
              "firstname": "Lewis",
              "lastname": "Hamilton",
              "year": 3
          }
      ],
      "message": "Student Queried Successfully"
    }

### Get Specific Student

This Endpoint Returns a specific Student

#### Request

`curl --location 'http://localhost:8001/student/getStudent/3'`

#### Response

    {
      "status": "Success",
      "data": {
        "id": 3,
        "firstname": "Lewis",
        "lastname": "Hamilton",
        "year": 3
      },
      "message": "Student Queried Successfully"
    }

### Update Student

This Endpoint Updates Data of a Student

#### Request

`curl --location --request PUT 'http://localhost:8001/student/' \
--header 'Content-Type: application/json' \
--data '{
"id":1,
"firstname":"Charles",
"lastname":"Leclerc",
"year":3
}'`

#### Response

    {
      "status": "Success",
      "data": {
          "id": 1,
          "firstname": "Charles",
          "lastname": "Leclerc",
          "year": 3
      },
      "message": "Student Updated Successfully"
    }

### Delete Student

This Endpoint Deletes a Student

#### Request

`curl --location --request DELETE
'http://localhost:8001/student/7'`

#### Response

    {
      "status": "Success",
      "data": {
        "id": 0,
        "firstname": "",
        "lastname": "",
        "year": 0
      },
      "message": "Student Deleted Successfully"
    }

### Search Student

This Endpoint can be used to search a student
based on their firstname or lastname and sort
the results. This endpoint give a paginated
response

#### Request

`curl --location --request GET 'http://localhost:8001/student/search' \
--header 'Content-Type: application/json' \
--data '{
"searchString":"charl",
"sortBy": {
"column":"firstname",
"direction":"ASC"
},
"pagination": {
"page":0,
"pageSize":2
}
}'`

#### Response

    {
      "status": "Success",
      "data": {
          "totalElements": 1,
          "data": [
              {
                  "id": 1,
                  "firstname": "Charles",
                  "lastname": "Leclerc",
                  "year": 3
              }
          ]
      },
      "message": "Student Queried Successfully"
    }

