# hospital-app
API Postman Documentation: https://documenter.getpostman.com/view/21439640/2sAXqzWdca

## File Structure

    ├── README.md
    ├── app
    │   ├── config
    │   │   └── config.go
    │   ├── database
    │   │   ├── database.go
    │   │   ├── model.go
    │   │   └── postgres.go
    │   ├── handler
    │   │   ├── handler.go
    │   │   ├── patientHttp.go
    │   │   └── staffHttp.go
    │   ├── middleware
    │   │   ├── authMiddleware.go
    │   │   └── middleware.go
    │   ├── repository
    │   │   ├── patient.go
    │   │   ├── patientPostgres.go
    │   │   ├── staff.go
    │   │   ├── staffAuthPostgres.go
    │   │   └── staffPostgres.go
    │   ├── schema
    │   │   ├── patient.go
    │   │   └── staff.go
    │   ├── server
    │   │   ├── ginServer.go
    │   │   └── server.go
    │   ├── service
    │   │   ├── helper.go
    │   │   ├── patientService.go
    │   │   ├── service.go
    │   │   ├── staffAuthService.go
    │   │   └── staffService.go
    │   ├── main.go
    │   └── test
    │       ├── init_test.go
    │       ├── patient_test.go
    │       └── staff_test.go
    └── nginx

# Test
    <!-- Run Docker for Test Database -->
    docker-compose -f docker-compose.test.yml up --build
    <!-- Run Test Commant -->
    cd ./app
    go test ./test

# Deploy
    docker-compose up --build
