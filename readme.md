
# Dating App (User Service)

This is the API documentation for the User Service. This service includes signup, login, and creating a user bio. Below are the step-by-step instructions for using this service.

Unfortunately I didn't have time to make this development using Docker, so with a heavy heart you need to install the software on your machine yourself 

## 1. Whats need to be Installed Before
    1. Golang (latest)
    2. Apache Kakfa (latest)
    3. PostgreSQL (latest)

## 2. Clone This Repo
    npm git clone https://github.com/anandaadam/user-service-dealls
    cd user-service-dealls
    go mod tidy
    cp .env.example .env
    
## 3. Setup
    1. Create database on your PostgreSQL, include create "user" schema
    2. Create "user" kafka topic
    3. Customize your env

## 4. Run the Service
    cd cmd
    air server --port 8080
    cd ..
    cd database
    goose -dir migrations postgres "postgres://adamanandasantoso:@localhost:5432/dating?sslmode=disable" up

The service is already to use, the service is running on Port :3000


